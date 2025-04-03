package checks

import (
	"strings"

	"github.com/onyz107/OnyDetect/internal/utils"
)

// CheckProcesses checks the currently running processes on the system
// to determine if any of them match known sandbox-related processes.
// It retrieves the list of running processes using the utils.GetRunningProcesses
// function and compares the base name of each process against a predefined
// map of sandbox process names (utils.SandboxProcesses).
//
// Returns true if a sandbox-related process is detected, otherwise false.
func CheckProcesses() bool {
	processes := utils.GetRunningProcesses()
	for _, process := range processes {
		baseName := process
		if lastSlash := strings.LastIndexAny(process, "/\\"); lastSlash != -1 {
			baseName = process[lastSlash+1:]
		}

		if utils.SandboxProcesses[strings.ToLower(baseName)] {
			return true
		}
	}
	return false
}
