package tomato

import (
	"time"

	"github.com/smartystreets/clock"
)

type Timer struct {
	clock   *clock.Clock
	sleeper *clock.Sleeper

	completion chan time.Time
}

func NewTimer() *Timer {
	return &Timer{
		completion: make(chan time.Time),
	}
}

func (this *Timer) Countdown(duration time.Duration) chan time.Time {
	go this.start(duration)
	return this.completion
}

func (this *Timer) start(duration time.Duration) {
	this.sleeper.Sleep(duration)
	this.completion <- this.clock.UTCNow()
}
