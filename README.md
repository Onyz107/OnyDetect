# OnyDetect - Virtual Machine & Sandbox Detection

![Go](https://img.shields.io/badge/Go-1.24.0-blue) ![Cross-Platform](https://img.shields.io/badge/Cross--Platform-Windows%20%7C%20Linux%20%7C%20MacOS-green)

OnyDetect is a lightweight, cross-platform library designed to detect virtual machines (VMs) and sandboxes by leveraging multiple detection techniques. It helps identify whether the program is running in a monitored environment, which can be useful for security, malware analysis, and anti-analysis mechanisms.

## Features

- ✅ **Hostname Check** - Compares the system hostname against a predefined list of common VM names.
- ✅ **BIOS Information Check** - Analyzes BIOS details for known VM signatures.
- ✅ **MAC Address Check** - Detects virtual network interfaces.
- ✅ **Hypervisor Detection** - Identifies if a hypervisor is present.
- ✅ **Process Inspection** - Scans running processes for virtualization-related tasks.
- ✅ **System File Analysis** - Looks for system files typically associated with virtual environments.
- ✅ **CPU Speed Check** - Detects anomalies in CPU behavior that indicate a VM.
- ✅ **Cross-Platform** - Works seamlessly on **Windows, Linux, and macOS**.

## Installation

To install OnyDetect, use `go get`:

```sh
go get github.com/Onyz107/OnyDetect
```

## Usage Example

Here’s a basic example of how to use OnyDetect in your Go project:

```go
package main

import (
	"fmt"
	"os"

	"github.com/Onyz107/OnyDetect/detectVM"
)

func main() {
	for _, r := range detectVM.Run() {
		if r {
			fmt.Println("VM detected.")
			// Virtual machine or sandbox detected, exit immediately.
			os.Exit(-1)
		}
	}

	fmt.Println("No VM detected.")
}
```

## Use Cases

- **Malware Evasion** - Prevent execution in controlled environments.
- **Security Research** - Identify sandboxed execution.
- **Licensing Enforcement** - Restrict software usage on virtualized instances.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to open issues and submit pull requests.

---

🔗 **Follow for updates:** [GitHub Profile](https://github.com/Onyz107)
