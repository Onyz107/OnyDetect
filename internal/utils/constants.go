package utils

var VmIndicatorsLower = []string{
	"vbox", "vmware", "kvm", "hyper-v", "xen", "parallels", "qemu", "virtual", "bhyve",
}

var VmFilesByOS = map[string][]string{
	"windows": {
		"C:\\Program Files\\Oracle\\VirtualBox\\VBoxTray.exe",
		"C:\\Program Files\\VMware\\VMware Tools\\vmtoolsd.exe",
	},
	"linux": {
		"/usr/bin/VBoxService",
		"/usr/sbin/vmtoolsd",
	},
	"darwin": {
		"/Library/Application Support/VMware Tools",
		"/Library/Application Support/VirtualBox Guest Additions",
		"/Library/Parallels/Tools",
	},
}

var SandboxProcesses = map[string]bool{
	"sandbox.exe":   true,
	"firejail":      true,
	"cuckoo":        true,
	"vmwaretray":    true,
	"vboxservice":   true,
	"vboxtray":      true,
	"processhacker": true,
	"wine":          true,
	"regedit":       true,
	"vmtoolsd":      true,
	"prl_tools":     true,
}

var VmMACPrefixes = []string{
	"00:05:69", // VMware
	"00:0C:29", // VMware
	"00:1C:14", // VMware
	"00:50:56", // VMware
	"08:00:27", // VirtualBox
	"52:54:00", // QEMU/KVM
	"00:16:3E", // Xen
	"00:1C:42", // Parallels
}
