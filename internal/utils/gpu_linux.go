//go:build linux

package utils

import (
	"os/exec"
)

// GetGPUName retrieves the name of the GPU on a Linux system by executing a shell command
// that searches for VGA-compatible devices using `lspci`. It returns the output of the
// command as a string. If an error occurs during the execution of the command, an empty
// string is returned.
//
// Note: This function is specific to Linux systems and relies on the `lspci` command
// being available in the system's PATH.
func GetGPUName() string {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
