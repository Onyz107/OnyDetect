//go:build windows

package utils

import (
	"os/exec"
	"syscall"
)

// GetDiskName retrieves the model name of the physical disk on a Windows system.
// It executes a PowerShell command to fetch the disk model information using
// the "Get-PhysicalDisk" cmdlet and returns the output as a string.
//
// The function hides the PowerShell window during execution by setting the
// SysProcAttr.HideWindow attribute. If an error occurs during the execution
// of the command, an empty string is returned.
//
// Returns:
//
//	string: The model name of the physical disk, or an empty string if an error occurs.
func GetDiskName() string {
	cmd := exec.Command("powershell", "-Command", "Get-PhysicalDisk | Select-Object Model")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
