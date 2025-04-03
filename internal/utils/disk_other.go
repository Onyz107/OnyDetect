//go:build !windows

package utils

import "os/exec"

// GetDiskName retrieves the disk model name by executing the "lsblk" command
// with the "-o MODEL" option. This function is intended to be used on non-Windows
// systems as indicated by the build constraint at the top of the file.
//
// It returns the output of the "lsblk" command as a string. If an error occurs
// while executing the command, an empty string is returned.
func GetDiskName() string {
	cmd := exec.Command("lsblk", "-o", "MODEL")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
