package tomato

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/clock"
	"github.com/smartystreets/gunit"
)

func TestTimerFixture(t *testing.T) {
	gunit.Run(new(TimerFixture), t)
}

type TimerFixture struct {
	*gunit.Fixture

	timer *Timer
}

func (this *TimerFixture) Setup() {
	this.timer = NewTimer()
}

func (this *TimerFixture) TestTimerSendsCompletionMessageAfterElapsed() {
	countdown := time.Minute

	later := time.Now().Add(countdown)
	this.timer.clock = clock.Freeze(later)
	this.timer.sleeper = clock.StayAwake()

	zero := <-this.timer.Countdown(countdown)

	this.So(zero, should.Equal, later)
	this.So(this.timer.sleeper.Naps, should.Resemble, []time.Duration{countdown})
}
