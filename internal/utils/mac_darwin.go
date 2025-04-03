//go:build darwin

package utils

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// GetMACAddresses retrieves a list of MAC addresses from MacOS's network interfaces.
// It executes the "ifconfig" command to gather network interface details and parses the output
// to extract MAC addresses associated with the "ether" keyword.
//
// Returns:
//
//	[]string: A slice of strings containing the MAC addresses of the network interfaces.
//	          If an error occurs while executing the command, an empty slice is returned.
func GetMACAddresses() []string {
	var macAddresses []string

	cmd := exec.Command("ifconfig")
	output, err := cmd.Output()
	if err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "ether") {
				parts := strings.Fields(line)
				if len(parts) > 1 {
					macAddresses = append(macAddresses, parts[1])
				}
			}
		}
	}

	return macAddresses
}
