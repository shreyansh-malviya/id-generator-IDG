package idgenerator

import (
	"errors"
	"sync"
	"time"
)

const SeqSize = 14
const MIDSize = 10
const TimeSize = 39

const RepTime = int64(10 * time.Millisecond)

var (
	ErrStartTimeAhead      = errors.New("start time is in the future")
	ErrNoMachineIDProvided = errors.New("settings.MachineID function must be provided")
	ErrInvalidMachineID    = errors.New("invalid machine ID")
	Err4                   = errors.New("one or more fields exceed their bit capacity")
)

const SeqMax = (1 << SeqSize) - 1 //11111111111111
const MIDMax = (1 << MIDSize) - 1 //1111111111

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

type IDGenerator struct {
	mu sync.Mutex

	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   uint16
}

func New(st Settings) (*IDGenerator, error) {

	now := time.Now().UTC()
	start := st.StartTime.UTC()
	if start.IsZero() {
		start = time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC)
	}
	if start.After(now) {
		return nil, ErrStartTimeAhead
	}
	if st.MachineID == nil {
		return nil, ErrNoMachineIDProvided
	}
	mid, err := st.MachineID()
	if err != nil {
		return nil, err
	}
	mid &= MIDMax
	if st.CheckMachineID != nil && !st.CheckMachineID(mid) {
		return nil, ErrInvalidMachineID
	}

	startGenUnits := toGeneratorTime(start)
	nowGenUnits := toGeneratorTime(now)

	return &IDGenerator{
		startTime:   startGenUnits,
		elapsedTime: nowGenUnits - startGenUnits,
		sequence:    0,
		machineID:   mid,
	}, nil
}

func toGeneratorTime(t time.Time) int64 {
	return t.UnixNano() / RepTime
}

func currentElapsedTime(startTime int64) int64 {
	nowUnits := toGeneratorTime(time.Now().UTC())
	return nowUnits - startTime
}

func sleepTime(overtime int64) time.Duration {
	if overtime <= 0 {
		return 0
	}
	return time.Duration(overtime*RepTime) * time.Nanosecond
}

func (g *IDGenerator) NextID() (uint64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for {
		curElapsed := currentElapsedTime(g.startTime)

		if curElapsed < g.elapsedTime {
			overtime := g.elapsedTime - curElapsed
			time.Sleep(sleepTime(overtime))
			continue
		}

		if curElapsed > g.elapsedTime {
			g.elapsedTime = curElapsed
			g.sequence = 0
		} else {
			g.sequence++
			g.sequence &= SeqMax
			if g.sequence == 0 {
				g.elapsedTime++
				overtime := g.elapsedTime - curElapsed
				time.Sleep(sleepTime(overtime))
				continue
			}
		}

		return g.toID()
	}
}

func (g *IDGenerator) toID() (uint64, error) {
	if g.elapsedTime < 0 || g.elapsedTime >= (1<<TimeSize) {
		return 0, Err4
	}
	if int(g.machineID) > MIDMax {
		return 0, Err4
	}
	if int(g.sequence) > SeqMax {
		return 0, Err4
	}
	id := (uint64(g.elapsedTime) << (SeqSize + MIDSize)) |
		(uint64(g.machineID) << SeqSize) |
		uint64(g.sequence)
	return id, nil
}
