//go:build windows

package utils

import "os/exec"

// GetBIOSInfo retrieves information about windows's BIOS using a PowerShell command.
// It executes the "Get-WmiObject Win32_BIOS" command to fetch details such as the
// Manufacturer and SMBIOSBIOSVersion, and formats the output as a list.
//
// Returns:
//   - a string containing the formatted BIOS information if the command executes successfully.
//   - an empty string is returned if an error occurs during command execution.
func GetBIOSInfo() string {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject Win32_BIOS | Select-Object Manufacturer, SMBIOSBIOSVersion | Format-List")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
