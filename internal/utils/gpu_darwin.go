//go:build darwin

package utils

import (
	"fmt"
	"os/exec"
)

// GetGPUName retrieves the name of the GPU on a macOS (Darwin) system.
// It executes a shell command to query the system profiler for display information
// and extracts the "Chipset Model" field.
//
// Returns:
//   - A string containing the GPU name if the command executes successfully.
//   - An empty string if an error occurs during command execution.
//
// Note:
//
//	This function is specific to macOS and uses the `system_profiler` command,
//	which may not be available or behave differently on other operating systems.
func GetGPUName() string {
	cmd := exec.Command("sh", "-c", `system_profiler SPDisplaysDataType | grep "Chipset Model"`)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return string(output)
}
