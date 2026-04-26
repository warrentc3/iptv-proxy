//go:build !linux

package validate

func getUnprivilegedPortStart() (startPort uint16, err error) {
	return 0, nil
}

func hasNetBindServiceCapability() (has bool, err error) {
	return false, nil
}
