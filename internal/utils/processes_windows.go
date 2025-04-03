//go:build windows

package utils

import (
	"os/exec"
	"strings"
)

// GetRunningProcesses retrieves a list of names of currently running processes on a Windows system.
// It executes the "tasklist" command with specific flags to get the process list in CSV format,
// parses the output, and extracts the process names.
//
// Returns:
//
//	[]string: A slice of strings containing the names of the running processes.
//	          If an error occurs during command execution, it returns nil.
//
// Note:
//
//	This function is specific to Windows and uses the "tasklist" command,
//	so it will not work on other operating systems.
func GetRunningProcesses() []string {
	cmd := exec.Command("tasklist", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(output), "\n")
	processes := make([]string, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		start := strings.IndexByte(line, '"')
		if start == -1 {
			continue
		}
		end := strings.IndexByte(line[start+1:], '"')
		if end == -1 {
			continue
		}

		processes = append(processes, line[start+1:start+1+end])
	}

	return processes
}
