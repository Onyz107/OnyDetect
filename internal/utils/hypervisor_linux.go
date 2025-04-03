//go:build linux

package utils

import (
	"os"
	"strings"
)

// GetHypervisorInfo checks if linux is running on a hypervisor by
// reading the contents of the "/proc/cpuinfo" file and searching for
// the presence of the term "hypervisor" (case-insensitive).
//
// Returns:
// - true: if the term "hypervisor" is found in the CPU information.
// - false: if the term is not found or if there is an error reading the file.
func GetHypervisorInfo() bool {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), "hypervisor")
}
