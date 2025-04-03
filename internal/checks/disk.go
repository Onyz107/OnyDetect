package checks

import (
	"strings"

	"github.com/Onyz107/OnyDetect/internal/utils"
)

// CheckDisk checks if the disk name contains any of the predefined VM indicators.
// It retrieves the disk name using the `utils.GetDiskName` function, converts it to lowercase,
// and compares it against a list of VM indicators (also in lowercase) from `utils.VmIndicatorsLower`.
// If a match is found, it returns true, indicating the presence of a VM-related disk.
// Otherwise, it returns false.
func CheckDisk() bool {
	disk := strings.ToLower(utils.GetDiskName())
	for _, indicator := range utils.VmIndicatorsLower {
		if strings.Contains(disk, indicator) {
			return true
		}
	}

	return false
}
