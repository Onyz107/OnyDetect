//go:build windows

package utils

import (
	"os/exec"
	"syscall"
)

// GetGPUName retrieves the name of the GPU installed on a Windows system.
// It executes a PowerShell command to query the Win32_VideoController class
// and extracts the Name property of the GPU using the Get-CimInstance cmdlet.
// The function hides the PowerShell window during execution for a seamless user experience.
// Returns the GPU name as a string if successful, or an empty string in case of an error.
func GetGPUName() string {
	cmd := exec.Command("powershell", "-Command", "Get-CimInstance Win32_VideoController | Select-Object Name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
