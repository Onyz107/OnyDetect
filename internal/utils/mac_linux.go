//go:build linux

package utils

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// GetMACAddresses retrieves a list of MAC addresses from linux's network interfaces.
// It executes the "ip link show" command to gather network interface details and parses
// the output to extract MAC addresses associated with the "link/ether" keyword.
//
// Returns:
//
//	[]string: A slice of strings containing the MAC addresses of the network interfaces.
//
// Note:
//
//	This function is designed to work on Linux systems and relies on the availability
//	of the "ip" command-line utility.
func GetMACAddresses() []string {
	var macAddresses []string

	cmd := exec.Command("ip", "link", "show")
	output, err := cmd.Output()
	if err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "link/ether") {
				parts := strings.Fields(line)
				if len(parts) > 1 {
					macAddresses = append(macAddresses, parts[1])
				}
			}
		}
	}

	return macAddresses
}
