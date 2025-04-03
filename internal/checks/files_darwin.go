//go:build darwin

package checks

import (
	"os/exec"
	"strings"

	"github.com/onyz107/OnyDetect/internal/utils"
)

// CheckSystemFiles checks for the presence of system files or indicators
// that might suggest the system is running in a virtualized environment.
// It executes the "ioreg -l" command to retrieve system information and
// searches the output for known virtualization indicators defined in
// utils.VmIndicatorsLower. If any indicator is found, it returns true.
// If no indicators are found or the command fails, it falls back to
// calling utils.CheckFiles() to perform additional checks.
//
// Returns:
//   - true: if virtualization indicators are detected or utils.CheckFiles()
//     returns true.
//   - false: if no indicators are found and utils.CheckFiles() returns false.
func CheckSystemFiles() bool {
	cmd := exec.Command("ioreg", "-l")
	output, err := cmd.Output()
	if err == nil {
		outputStr := strings.ToLower(string(output))
		for _, indicator := range utils.VmIndicatorsLower {
			if strings.Contains(outputStr, indicator) {
				return true
			}
		}
	}

	return utils.CheckFiles()
}
