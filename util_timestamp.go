package vrage

import (
	"fmt"
	"strconv"
	"time"
)

const (
	ticksPerSecond   = 10_000_000
	dotNetEpochTicks = 621_355_968_000_000_000
	nanosPerTick     = 100
)

// DotNetTimestampToTime converts a .NET timestamp string to a time.Time object.
func DotNetTimestampToTime(dotNetTimestamp string) (time.Time, error) {
	ticks, err := strconv.ParseInt(dotNetTimestamp, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid dotNetTimestamp %q: %w", dotNetTimestamp, err)
	}

	unixTicks := ticks - dotNetEpochTicks
	sec := unixTicks / ticksPerSecond
	nsec := (unixTicks % ticksPerSecond) * nanosPerTick

	return time.Unix(sec, nsec).UTC(), nil
}
