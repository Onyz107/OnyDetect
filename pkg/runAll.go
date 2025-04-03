package pkg

import "github.com/onyz107/OnyDetect/internal/checks"

var checksToRun = []func() bool{
	checks.CheckHostanme,
	checks.CheckBIOS,
	checks.CheckMACAddress,
	checks.CheckHypervisorPresent,
	checks.CheckProcesses,
	checks.CheckSystemFiles,
	checks.CheckCPUSpeed,
}

// Run executes a series of predefined system checks and returns their results.
// Each check is a function that returns a boolean value indicating whether
// the check passed or failed. The results are returned as a slice of booleans,
// where each element corresponds to the result of a specific check.
//
// Returns:
//   []bool: A slice containing the results of the executed checks. Each element
//           is `true` if the corresponding check passed, and `false` otherwise.
func Run() []bool {
	results := make([]bool, 7)

	for _, check := range checksToRun {
		results = append(results, check())
	}

	return results
}
