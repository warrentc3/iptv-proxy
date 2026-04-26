package validate

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ListeningAddress validates a listening address string given a user ID `uid`.
// If the address is empty, it is valid.
// If the port is 0, it is valid.
// If the port is below the start of unprivileged ports, it is valid if the user
// ID is 0 or -1 (windows); for any other user id, the running program Linux
// capabilities are checked to see if it can bind to privileged ports.
func ListeningAddress(address string, uid int) (err error) {
	if address == "" { // listen on all interfaces on any available port
		return nil
	}

	_, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("splitting host and port: %w", err)
	}

	return validatePort(portStr, uid)
}

func validatePort(value string, uid int) (err error) {
	port, err := parsePortString(value)
	if err != nil {
		return err
	}

	isWindows := uid == -1
	isRoot := uid == 0
	if port == 0 || isWindows || isRoot {
		return nil
	}

	unprivilegedPortStart, err := getUnprivilegedPortStart()
	if err != nil {
		return fmt.Errorf("getting unprivileged port start: %w", err)
	}

	if port >= unprivilegedPortStart {
		return nil
	}

	hasCapability, err := hasNetBindServiceCapability()
	if err != nil {
		return fmt.Errorf("getting net bind service capability: %w", err)
	} else if hasCapability {
		return nil
	}

	return makePrivilegedPortError(port, uid, unprivilegedPortStart)
}

var (
	ErrPortNotAnInteger = errors.New("port value is not an integer")
	ErrPortTooHigh      = errors.New("port cannot be higher than 65535")
)

func parsePortString(portString string) (port uint16, err error) {
	const maxPort = 65535
	const base, bitSize = 10, 64
	portUint64, err := strconv.ParseUint(portString, base, bitSize)
	switch {
	case err != nil:
		return 0, fmt.Errorf("%w: %s", ErrPortNotAnInteger, portString)
	case portUint64 > maxPort:
		return 0, fmt.Errorf("%w: %d", ErrPortTooHigh, portUint64)
	}

	return uint16(portUint64), nil
}

var (
	ErrPrivilegedPort = errors.New("listening on privileged port is not allowed")
)

func makePrivilegedPortError(port uint16, uid int,
	unprivilegedPortStart uint16) (err error) {
	errorDetails := []string{
		fmt.Sprintf("user id %d", uid),
	}
	const traditionalUnprivilegedPortStart = 1024
	if unprivilegedPortStart != traditionalUnprivilegedPortStart {
		errorDetails = append(errorDetails,
			fmt.Sprintf("unprivileged start port %d", unprivilegedPortStart))
	}

	return fmt.Errorf("%w: port %d (%s)", ErrPrivilegedPort,
		port, strings.Join(errorDetails, ", "))
}
