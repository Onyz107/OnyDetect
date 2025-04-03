package detectVM

import (
	"sync"

	"github.com/Onyz107/OnyDetect/internal/checks"
)

var checksToRun = [10]func() bool{
	checks.CheckHostanme,
	checks.CheckBIOS,
	checks.CheckMACAddress,
	checks.CheckHypervisorPresent,
	checks.CheckProcesses,
	checks.CheckSystemFiles,
	checks.CheckCPUSpeed,
	checks.CheckDisk,
	checks.CheckGPU,
	checks.CheckTimeAnomaly,
}

// Run executes a series of predefined system checks and returns their results.
// Each check is a function that returns a boolean value indicating whether
// the check passed or failed. The results are returned as a slice of booleans,
// where each element corresponds to the result of a specific check.
//
// Returns:
//
//	[]bool: A slice containing the results of the executed checks. Each element
//	        is `true` if the corresponding check passed, and `false` otherwise.
func Run() []bool {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]bool, len(checksToRun))

	for i, check := range checksToRun {
		wg.Add(1)
		go func(idx int, fn func() bool) {
			defer wg.Done()

			result := fn()

			mu.Lock()
			results[idx] = result
			mu.Unlock()
		}(i, check)
	}

	wg.Wait()
	return results
}
