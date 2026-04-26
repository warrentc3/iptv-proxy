package validate

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"kernel.org/pub/linux/libs/security/libcap/cap"
)

// getUnprivilegedPortStart returns the first port that can be
// used by unprivileged users.
// Note the unprivileged port start can be changed with
// net.ipv4.ip_unprivileged_port_start. Thanks to the-maldridge
// for the explanation at
// https://github.com/qdm12/ddns-updater/issues/335#issuecomment-1868406779
func getUnprivilegedPortStart() (startPort uint16, err error) {
	const filepath = "/proc/sys/net/ipv4/ip_unprivileged_port_start"
	const maxFileSize = 6 // 65535\n
	buffer := make([]byte, maxFileSize)
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) || errors.Is(err, os.ErrPermission) {
			const defaultUnprivilegedPortStart = 1024
			return defaultUnprivilegedPortStart, nil
		}
		return 0, fmt.Errorf("opening file: %w", err)
	}

	n, err := file.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("reading file: %w", err)
	}
	buffer = buffer[:n]

	s := string(buffer)
	s = strings.TrimSuffix(s, "\n")
	return parsePortString(s)
}

func hasNetBindServiceCapability() (has bool, err error) {
	capabilitySet := cap.GetProc()
	flagsRequired := []cap.Flag{cap.Effective, cap.Permitted}
	for _, flagRequired := range flagsRequired {
		ok, err := capabilitySet.GetFlag(flagRequired, cap.NET_BIND_SERVICE)
		if err != nil {
			return false, fmt.Errorf("getting %s flag: %w", flagRequired, err)
		} else if !ok {
			return false, nil
		}
	}

	return true, nil
}
