package checks

import (
	"time"

	"github.com/onyz107/OnyDetect/internal/utils"
)

// CheckCPUSpeed measures the average execution time of a loop over a fixed number
// of iterations and determines if the CPU speed is below a certain threshold.
// It uses the MeasureLoop function from the utils package to measure the time
// taken for each iteration.
//
// Returns:
//   - true if the average execution time exceeds 200 milliseconds, indicating
//     potentially slow CPU performance.
//   - false otherwise.
func CheckCPUSpeed() bool {
	const iterations = 5
	var total time.Duration
	for range iterations {
		total += utils.MeasureLoop()
	}
	avgTime := total / iterations

	return avgTime > 200*time.Millisecond
}
