package checks

import (
	"strings"

	"github.com/Onyz107/OnyDetect/internal/utils"
)

// CheckBIOS checks the system's BIOS information to determine if it contains
// any indicators that suggest the system is running in a virtualized environment.
// It retrieves the BIOS information using the utils.GetBIOSInfo function, converts
// it to lowercase, and compares it against a list of known virtualization-related
// indicators defined in utils.VmIndicatorsLower.
//
// Returns true if any virtualization indicators are found in the BIOS information,
// otherwise returns false.
func CheckBIOS() bool {
	biosInfo := utils.GetBIOSInfo()

	if biosInfo != "" {
		biosInfo = strings.ToLower(biosInfo)
		for _, indicator := range utils.VmIndicatorsLower {
			if strings.Contains(biosInfo, indicator) {
				return true
			}
		}
	}
	return false
}
