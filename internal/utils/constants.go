package utils

var VmIndicatorsLower = []string{
	"vbox",      // VirtualBox
	"vmware",    // VMware
	"kvm",       // Kernel-based Virtual Machine
	"hyper-v",   // Microsoft Hyper-V
	"xen",       // Xen Hypervisor
	"parallels", // Parallels Desktop
	"qemu",      // QEMU Emulator
	"virtual",   // Generic Virtualization Indicator
	"bhyve",     // FreeBSD Hypervisor
	"vmm",       // Virtual Machine Monitor (common VM-related term)
	"qnx",       // QNX Real-Time OS (often virtualized)
	"openvz",    // OpenVZ Containers
	"lxc",       // Linux Containers (LXC)
	"wsl",       // Windows subsystem for linux (WSL)
	"malware",
	"maltest",
	"currentuser",
	"sandbox",
	"virus",
	"john doe",
	"test user",
	"sand box",
	"wdagutilityaccount",
}

var VmFilesByOS = map[string][]string{
	"windows": {
		// VirtualBox
		"C:\\Program Files\\Oracle\\VirtualBox\\VBoxTray.exe",
		"C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxService.exe",

		// VMware
		"C:\\Program Files\\VMware\\VMware Tools\\vmtoolsd.exe",
		"C:\\Program Files\\VMware\\VMware Tools\\vmwaretray.exe",

		// Parallels
		"C:\\Program Files\\Parallels\\Parallels Tools Services.exe",

		// QEMU/KVM
		"C:\\Program Files\\qemu-ga\\qemu-ga.exe",

		// Microsoft Sandbox
		"C:\\ProgramData\\Microsoft\\Windows\\Containers\\Sandbox",
		"C:\\Windows\\System32\\VmComputeAgent.exe",
	},

	"linux": {
		// VirtualBox
		"/usr/bin/VBoxService",
		"/usr/bin/VBoxClient",
		"/usr/sbin/VBoxTray",
		"/lib/modules/*/misc/vboxguest.ko",
		"/lib/modules/*/misc/vboxsf.ko",
		"/lib/modules/*/misc/vboxvideo.ko",

		// VMware
		"/usr/bin/vmtoolsd",
		"/usr/bin/vmware-user",
		"/usr/sbin/vmware-checkvm",
		"/lib/modules/*/misc/vmci.ko",
		"/lib/modules/*/misc/vmhgfs.ko",
		"/lib/modules/*/misc/vmxnet.ko",

		// Hyper-V
		"/dev/hv_vmbus",
		"/lib/modules/*/kernel/drivers/hv/hv_vmbus.ko",
		"/lib/modules/*/kernel/drivers/hv/hv_utils.ko",
		"/lib/modules/*/kernel/drivers/hv/hv_storvsc.ko",

		// Parallels
		"/usr/bin/prltoolsd",
		"/usr/bin/prlcc",
		"/usr/bin/prl_fs_freeze",
		"/lib/modules/*/misc/prl_tg.ko",
		"/lib/modules/*/misc/prl_vid.ko",

		// QEMU/KVM
		"/usr/sbin/qemu-ga",
		"/dev/kvm",
		"/usr/libexec/qemu-kvm",
	},

	"darwin": {
		// VMware
		"/Library/Application Support/VMware Tools/vmware-tools-daemon",
		"/Library/Application Support/VMware Tools/vmware-user",

		// VirtualBox
		"/Library/Application Support/VirtualBox Guest Additions/VBoxService",
		"/Library/Application Support/VirtualBox Guest Additions/VBoxTray",

		// Parallels
		"/Library/Parallels/Tools/prl_cc",
		"/Library/Parallels/Tools/prl_toolsd",
	},
}

var SandboxProcesses = map[string]bool{
	"sandbox.exe":        true, // Generic Sandbox
	"firejail":           true, // Linux sandboxing tool
	"cuckoo":             true, // Cuckoo Sandbox
	"vmwaretray":         true, // VMware Tray
	"vboxservice":        true, // VirtualBox Service
	"vboxtray":           true, // VirtualBox Tray
	"processhacker":      true, // Process Hacker (used for debugging/reversing)
	"wine":               true, // Wine (Windows emulator for Linux)
	"regedit":            true, // Windows Registry Editor (could indicate debugging)
	"vmtoolsd":           true, // VMware Tools
	"prl_tools":          true, // Parallels Tools
	"qemu-ga":            true, // QEMU Guest Agent
	"vmsrvc.exe":         true, // VMware Service
	"vmusrvc.exe":        true, // VMware User Service
	"df5serv.exe":        true, // Deep Freeze Service (sandbox-related)
	"vgauthservice.exe":  true, // VMware Authorization Service
	"vmacthlp.exe":       true, // VMware Tools Activity Helper
	"xenservice.exe":     true, // Xen Guest Service
	"vboxguest.exe":      true, // VirtualBox Guest Additions
	"vboxrun.exe":        true, // VirtualBox Run Process
	"wireshark":          true, // Wireshark (often used in analysis environments)
	"ida.exe":            true, // IDA Pro (debugging tool)
	"x64dbg.exe":         true, // x64dbg Debugger
	"ollydbg.exe":        true, // OllyDbg Debugger
	"windbg.exe":         true, // Windows Debugger
	"tcpview.exe":        true, // Sysinternals TCP View (network analysis tool)
	"autoruns.exe":       true, // Sysinternals Autoruns (process analysis tool)
	"procmon.exe":        true, // Sysinternals Process Monitor
	"procexp.exe":        true, // Sysinternals Process Explorer
	"fakenet":            true, // FakeNet (used to simulate network traffic)
	"windowssandbox.exe": true,
	"appvshnotify.exe":   true,
	"sandboxed.dll":      true,
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
	"00:15:5D", // Hyper-V
	"00:03:FF", // Microsoft Virtual PC
	"00:0F:4B", // Virtual Iron
	"00:25:AE", // Microsoft Hyper-V
}
