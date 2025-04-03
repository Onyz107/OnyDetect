//go:build linux

package utils

import (
	"os"
	"strings"
)

// GetBIOSInfo retrieves the BIOS vendor and version information from linux.
// It reads the data from the Linux-specific files located at "/sys/class/dmi/id/bios_vendor"
// and "/sys/class/dmi/id/bios_version". If any error occurs while reading these files,
// the function returns an empty string.
//
// Returns:
//   - a string containing the BIOS vendor and version concatenated with a space in between.
//   - an empty string if the information cannot be retrieved.
func GetBIOSInfo() string {
	data, err := os.ReadFile("/sys/class/dmi/id/bios_vendor")
	if err != nil {
		return ""
	}
	vendor := strings.TrimSpace(string(data))

	data, err = os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		return ""
	}
	version := strings.TrimSpace(string(data))

	return vendor + " " + version
}
