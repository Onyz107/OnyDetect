package checks

import (
	"strings"

	"github.com/Onyz107/OnyDetect/internal/utils"
)

// CheckGPU checks if the GPU name contains any indicators that suggest
// the system might be running in a virtualized environment. It retrieves
// the GPU name using the utility function `utils.GetGPUName`, converts it
// to lowercase, and compares it against a list of known VM-related indicators
// (`utils.VmIndicatorsLower`). If a match is found, it returns true, indicating
// the presence of a virtualized GPU. Otherwise, it returns false.
func CheckGPU() bool {
	gpu := strings.ToLower(utils.GetGPUName())
	for _, indicator := range utils.VmIndicatorsLower {
		if strings.Contains(gpu, indicator) {
			return true
		}
	}

	return false
}
