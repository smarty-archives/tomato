package tomato

import (
	"io"
	"time"
)

type TimeCounter interface {
	Countdown(duration time.Duration) chan time.Time
}

type OS interface {
	FocusApp(name string)
	LockScreen()
}

type Controller struct {
	terminal io.ReadWriter
	timer    TimeCounter
	os       OS
}
