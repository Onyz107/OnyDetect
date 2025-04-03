//go:build darwin

package utils

import "os/exec"

// GetBIOSInfo retrieves information about the system's BIOS on macOS.
// It executes the "system_profiler SPHardwareDataType" command to gather
// hardware details and returns the output as a string. If an error occurs
// during command execution, an empty string is returned.
func GetBIOSInfo() string {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
