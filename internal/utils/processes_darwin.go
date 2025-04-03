//go:build darwin

package utils

import (
	"os/exec"
	"strings"
)

// GetRunningProcesses retrieves a list of currently running processes on a Darwin (macOS) system.
// It executes the "ps" command with the options "-e" (to list all processes) and "-o comm"
// (to display the command name of each process).
//
// Returns:
// - A slice of strings containing the names of running processes.
// - If an error occurs while executing the command, it returns nil.
//
// Note:
// - The first line of the "ps" command output is a header and is skipped.
// - Empty lines in the output are ignored.
func GetRunningProcesses() []string {
	cmd := exec.Command("ps", "-e", "-o", "comm")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(output), "\n")
	processes := make([]string, 0, len(lines)-1) // -1 for header

	// Skip the header line
	for i := 1; i < len(lines); i++ {
		if line := strings.TrimSpace(lines[i]); line != "" {
			processes = append(processes, line)
		}
	}

	return processes
}
