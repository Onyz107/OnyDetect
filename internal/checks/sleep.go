package checks

import "time"

// CheckTimeAnomaly detects potential time anomalies by measuring the elapsed time
// for a sleep operation. It compares the measured elapsed time against an expected
// range to determine if there is a significant deviation.
//
// The function starts by recording the current time, then sleeps for 1 second.
// After the sleep, it calculates the elapsed time and checks if it falls outside
// the range of 900ms to 1100ms. If the elapsed time is outside this range, it
// returns true, indicating a potential anomaly. Otherwise, it returns false.
//
// This function can be useful for detecting discrepancies in time measurement,
// which might occur in virtualized environments or under certain system conditions.
func CheckTimeAnomaly() bool {
	start := time.Now()
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)

	// In a real machine, elapsed should be close to 1s, in a VM, it may be off.
	if elapsed.Milliseconds() < 900 || elapsed.Milliseconds() > 1100 {
		return true
	}

	return false
}
