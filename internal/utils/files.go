package utils

import (
	"os"
	"runtime"
)

// CheckFiles checks for the existence of specific files on the system
// based on the operating system. It uses a predefined map `VmFilesByOS`
// that associates operating systems with a list of file paths.
//
// The function retrieves the relevant file paths for the current OS
// using the `runtime.GOOS` value. If the OS is not found in the map,
// the function returns false. Otherwise, it iterates through the file
// paths and checks if any of the files exist.
//
// Returns:
//   - true: If at least one of the files exists.
//   - false: If none of the files exist or the OS is not supported.
func CheckFiles() bool {
	relevantFiles, exists := VmFilesByOS[runtime.GOOS]
	if !exists {
		return false
	}

	for _, filePath := range relevantFiles {
		if _, err := os.Stat(filePath); err == nil {
			return true
		}
	}
	return false
}
