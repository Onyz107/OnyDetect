//go:build linux

package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetRunningProcesses retrieves a list of names of currently running processes
// on a Linux system. It reads the /proc directory to identify process IDs (PIDs)
// and extracts the process names from the "comm" file for each PID.
//
// Returns:
//
//	[]string: A slice of strings containing the names of running processes.
//	          If an error occurs while reading the /proc directory or process
//	          information, it returns nil.
//
// Notes:
//   - This function is specific to Linux systems and relies on the /proc filesystem.
//   - Processes without a valid "comm" file or with empty names are ignored.
func GetRunningProcesses() []string {
	files, err := os.ReadDir("/proc")
	if err != nil {
		return nil
	}

	processes := make([]string, 0, len(files)/2) // Estimate half are processes

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if _, err := strconv.Atoi(file.Name()); err != nil {
			continue
		}

		// Read comm file for process name
		cmdlinePath := filepath.Join("/proc", file.Name(), "comm")
		data, err := os.ReadFile(cmdlinePath)
		if err != nil {
			continue
		}

		processName := strings.TrimSpace(string(data))
		if processName != "" {
			processes = append(processes, processName)
		}
	}

	return processes
}
