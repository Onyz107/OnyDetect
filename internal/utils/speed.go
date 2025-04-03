package utils

import "time"

// MeasureLoop measures the time taken to execute a loop that performs
// a simple computation (squaring the loop index) for a fixed number
// of iterations. It returns the duration of the loop execution.
func MeasureLoop() time.Duration {
	start := time.Now()
	for i := range 50000000 {
		_ = i * i
	}
	return time.Since(start)
}
