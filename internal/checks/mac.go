package checks

import (
	"strings"

	"github.com/Onyz107/OnyDetect/internal/utils"
)

// CheckMACAddress checks if any of the MAC addresses on the system
// match a known prefix associated with virtual machines. It retrieves
// the list of MAC addresses using the utils.GetMACAddresses function
// and compares each address (converted to lowercase) against a list
// of known virtual machine MAC address prefixes (utils.VmMACPrefixes).
//
// Returns true if a match is found, indicating the presence of a
// virtual machine MAC address. Otherwise, returns false.
func CheckMACAddress() bool {
	for _, mac := range utils.GetMACAddresses() {
		mac = strings.ToLower(mac)
		for _, prefix := range utils.VmMACPrefixes {
			if strings.HasPrefix(mac, strings.ToLower(prefix)) {
				return true
			}
		}
	}

	return false
}
