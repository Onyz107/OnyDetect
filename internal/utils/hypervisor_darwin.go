//go:build darwin

package utils

import (
	"os/exec"
	"strings"
)

// GetHypervisorInfo checks if the current system is running under a hypervisor
// on macOS (Darwin). It uses system commands to detect virtualization-related
// features and flags.
//
// The function performs the following checks:
//  1. Executes the "sysctl -n machdep.cpu.features" command to retrieve CPU features
//     and looks for the "vmm" flag, which indicates the presence of a hypervisor.
//  2. Executes the "sysctl -n kern.hv_support" command to check if the Hypervisor
//     Framework is supported. If supported, it further checks the CPU brand string
//     for known virtualization indicators.
//
// Returns:
// - true: If the system is detected to be running under a hypervisor.
// - false: If no hypervisor is detected or if an error occurs during the checks.
func GetHypervisorInfo() bool {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.features")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	// Check for VMM flag which indicates hypervisor presence
	if strings.Contains(strings.ToLower(string(output)), "vmm") {
		return true
	}

	// Additional check for hvf_tid which is related to Hypervisor Framework
	cmd = exec.Command("sysctl", "-n", "kern.hv_support")
	output, err = cmd.Output()
	if err == nil && string(output) == "1\n" {
		// Check if running under virtualization
		cmd = exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
		output, err = cmd.Output()
		if err == nil {
			brand := strings.ToLower(string(output))
			for _, indicator := range VmIndicatorsLower {
				if strings.Contains(brand, indicator) {
					return true
				}
			}
		}
	}

	return false
}
