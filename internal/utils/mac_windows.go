//go:build windows

package utils

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// GetMACAddresses retrieves a list of MAC addresses from the system.
// It executes the "getmac" command with specific arguments to obtain
// the MAC addresses in CSV format, processes the output, and extracts
// the MAC addresses from the relevant column.
//
// Returns:
//
//	[]string: A slice containing the MAC addresses as strings. If an
//	error occurs during command execution or parsing, an empty slice
//	is returned.
func GetMACAddresses() []string {
	var macAddresses []string

	cmd := exec.Command("getmac", "/v", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ",")
			if len(parts) > 3 {
				mac := strings.Trim(parts[2], "\"")
				macAddresses = append(macAddresses, mac)
			}
		}
	}

	return macAddresses
}
