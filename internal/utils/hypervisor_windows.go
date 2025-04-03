//go:build windows

package utils

import (
	"os/exec"
	"strings"
)

// GetHypervisorInfo checks if virtualization is enabled in windows's firmware.
// It executes a PowerShell command to query the Win32_Processor class for the
// VirtualizationFirmwareEnabled property. If the property is set to "true",
// the function returns true, indicating that virtualization is enabled.
// Otherwise, it returns false.
//
// Returns:
//   - bool: true if virtualization is enabled, false otherwise.
func GetHypervisorInfo() bool {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject Win32_Processor | Select-Object VirtualizationFirmwareEnabled")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(output)), "true")
}
