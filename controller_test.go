package tomato

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/clock"
	"github.com/smartystreets/gunit"
)

func TestControllerTestFixture(t *testing.T) {
	gunit.Run(new(ControllerTestFixture), t)
}

type ControllerTestFixture struct {
	*gunit.Fixture

	timer          *Timer
	os             *FakeOS
	terminalWriter *bytes.Buffer
	terminalReader *bytes.Buffer
	terminal       *InterruptingTerminal
	controller     *Controller
	interrupt      chan os.Signal
}

var (
	start = time.Now()
	end   = start.Add(time.Minute * 25)
)

func (this *ControllerTestFixture) Setup() {
	this.interrupt = make(chan os.Signal, 10)
	this.timer = NewTimer()
	this.timer.sleeper = clock.StayAwake()
	this.os = &FakeOS{}
	this.terminalWriter = new(bytes.Buffer)
	this.terminalReader = new(bytes.Buffer)
	this.terminal = &InterruptingTerminal{
		interrupt: this.interrupt,
		SmartyTerminal: &SmartyTerminal{
			Reader: this.terminalReader, Writer: this.terminalWriter,
		},
	}
	this.controller = NewController(this.terminal, this.timer, this.os, this.interrupt)
}

func (this *ControllerTestFixture) TestGameOverUponInterruptSignalReceivedAndCausesRunToFinishNaturally() {
	this.interrupt <- new(FakeSignal)
	this.controller.Run()
}

func (this *ControllerTestFixture) TestFullTomato() {
	this.terminal.interruptCountdown = 1
	this.terminalReader.WriteString("\n")

	this.controller.Run()

	this.So(this.timer.sleeper.Naps, should.Resemble, []time.Duration{
		time.Minute * 24, // session
		time.Minute,      // last-minute reminder
		time.Second * 5,  // allow time to read "take a break"
		time.Minute * 5,  // break
	})
	this.So(this.terminalWriter.String(), should.StartWith, expectedOutput_1Tomato)
	this.So(this.os.notified, should.BeGreaterThan, 0)
	this.So(this.os.locked, should.BeGreaterThan, 0)
	this.So(this.os.focused, should.BeGreaterThan, 0)
}

func (this *ControllerTestFixture) TestFourFullTomatoesEndsWithLongerBreak() {
	this.terminal.interruptCountdown = 4
	this.terminalReader.WriteString("\n")
	this.terminalReader.WriteString("\n")
	this.terminalReader.WriteString("\n")
	this.terminalReader.WriteString("\n")

	this.controller.Run()

	this.So(this.terminalWriter.String(), should.StartWith, expectedOutput_4Tomatos)
	this.So(this.terminalWriter.String(), should.ContainSubstring, "Time's up! Go take a longer, 15 minute, break. Locking screen...")
}

const expectedOutput_1Tomato = `Setting timer for 25 minutes...Go!
1 minute remaining...
Time's up! Go take a 5 minute break. Locking screen...
<ENTER> to start a new 25 minute tomato...
`

const expectedOutput_4Tomatos = `Setting timer for 25 minutes...Go!
1 minute remaining...
Time's up! Go take a 5 minute break. Locking screen...
<ENTER> to start a new 25 minute tomato...

Setting timer for 25 minutes...Go!
1 minute remaining...
Time's up! Go take a 5 minute break. Locking screen...
<ENTER> to start a new 25 minute tomato...

Setting timer for 25 minutes...Go!
1 minute remaining...
Time's up! Go take a 5 minute break. Locking screen...
<ENTER> to start a new 25 minute tomato...

Setting timer for 25 minutes...Go!
1 minute remaining...
Time's up! Go take a longer, 15 minute, break. Locking screen...
<ENTER> to start a new 25 minute tomato...
`

///////////////////////////////////////////////////

type FakeOS struct {
	focused  int
	notified int
	locked   int
}

func (this *FakeOS) FocusApp(name string)         { this.focused++ }
func (this *FakeOS) LockScreen()                  { this.locked++ }
func (this *FakeOS) Notify(title, message string) { this.notified++ }

///////////////////////////////////////////////////

type FakeSignal struct{}

func (FakeSignal) String() string { panic("implement me") }
func (FakeSignal) Signal()        { panic("implement me") }

//////////////////////////////////////////////////

type InterruptingTerminal struct {
	*SmartyTerminal

	interruptCountdown int
	interrupt          chan os.Signal
}

func (this *InterruptingTerminal) Read(p []byte) (n int, err error) {
	this.interruptCountdown--
	if this.interruptCountdown <= 0 {
		this.interrupt <- new(FakeSignal)
	}
	return this.SmartyTerminal.Read(p)
}
