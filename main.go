package idgenerator

import (
	"time"
)

// Settings contains configuration for the ID generator.
type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

// IDGenerator represents the generator instance.
// TODO: define required fields for internal state.
type IDGenerator struct {
	// TODO: add necessary fields (mutex, startTime, elapsedTime, sequence, machineID, etc.)
}

// TODO: Define the required error variables (e.g., ErrStartTimeAhead, ErrNoMachineIDProvided, etc.)

// New creates and initializes an IDGenerator according to the settings.
// TODO: implement
func New(st Settings) (*IDGenerator, error) {
	// Steps to implement:
	// 1. Validate StartTime (if StartTime.After(now) -> return ErrStartTimeAhead).
	// 2. If StartTime.IsZero(), set default epoch to 2014-09-01 00:00:00 UTC.
	// 3. Validate that st.MachineID is not nil. If nil -> return ErrNoMachineIDProvided.
	// 4. Call st.MachineID(), return error if it fails.
	// 5. If st.CheckMachineID != nil and it returns false, return ErrInvalidMachineID.
	// 6. Initialize and return *IDGenerator with proper starting sequence & times.
	return nil, nil
}

// NextID generates the next unique ID. Must be concurrency-safe.
// TODO: implement
func (g *IDGenerator) NextID() (uint64, error) {
	// Hints:
	// - Use maskSequence := uint16(1<<BitLenSequence - 1)
	// - Lock g.mutex at the start of the critical section and defer Unlock.
	// - Compute current elapsed time units using currentElapsedTime(g.startTime).
	// - If sequence wraps (becomes zero after masking), increment elapsedTime and call time.Sleep(sleepTime(overtime)).
	// - Call g.toID() to pack fields and return the value.
	return 0, nil
}

// Helpers you must implement:

// toGeneratorTime converts a time.Time into generator units since Unix epoch.
func toGeneratorTime(t time.Time) int64 {
	// TODO: implement conversion to units of 10 ms
	return 0
}

// currentElapsedTime returns elapsed time units since startTime.
func currentElapsedTime(startTime int64) int64 {
	// TODO: implement
	return 0
}

// sleepTime returns the duration to sleep when the generator is ahead of the real current time.
func sleepTime(overtime int64) time.Duration {
	// TODO: implement: compute how long to sleep (in nanoseconds) so that the internal
	// elapsedTime will catch up to the real time unit boundary.
	return 0
}

// toID packs internal fields into a uint64 ID, or returns an error.
// TODO: implement
func (g *IDGenerator) toID() (uint64, error) {
	return 0, nil
}
