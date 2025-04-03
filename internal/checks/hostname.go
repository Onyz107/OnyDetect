package checks

import (
	"os"
	"strings"

	"github.com/onyz107/OnyDetect/internal/utils"
)

// CheckHostname checks the system's hostname to determine if it contains
// any indicators that suggest the system is running in a virtual machine.
// It retrieves the hostname, converts it to lowercase, and compares it
// against a predefined list of VM-related substrings (utils.VmIndicatorsLower).
//
// Returns true if the hostname contains any of the VM indicators, otherwise false.
func CheckHostanme() bool {
	hostname, err := os.Hostname()
	if err == nil {
		hostname = strings.ToLower(hostname)
		for _, indicator := range utils.VmIndicatorsLower {
			if strings.Contains(hostname, indicator) {
				return true
			}
		}
	}

	return false
}
