package checks

import "github.com/onyz107/OnyDetect/internal/utils"

// CheckHypervisorPresent checks if a hypervisor is present on the system.
// It utilizes the GetHypervisorInfo function from the utils package to determine
// the presence of a hypervisor. Returns true if a hypervisor is detected, otherwise false.
func CheckHypervisorPresent() bool {
	return utils.GetHypervisorInfo()
}
