package tomato

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestControllerFixture(t *testing.T) {
	gunit.RunSequential(new(ControllerFixture), t)
}

type ControllerFixture struct {
	*gunit.Fixture

	controller *Controller
	output     *bytes.Buffer
}

func (this *ControllerFixture) Setup() {
	this.output = new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(this.output)
}
func (this *ControllerFixture) initializeController(sessions int) {
	consoleInput := strings.NewReader(strings.Repeat("\n", sessions))
	this.controller = NewController(consoleInput, new(NopSystem), sessions, time.Nanosecond)
}

func (this *ControllerFixture) TestSessionInteractionsAndOutput() {
	this.initializeController(8)
	this.controller.Run()
	this.So(this.output.String(), should.Equal, expectedTomatoSession)
}

type NopSystem struct{}

func (NopSystem) Notify(message string)        { log.Println("[Notify]", message) }
func (NopSystem) FocusApp(name string)         { log.Println("[FocusApp]", name) }
func (NopSystem) LockScreen()                  { log.Println("[LockScreen]") }
func (NopSystem) Sleep(duration time.Duration) { log.Println("[Sleep]", duration) }

const expectedTomatoSession = `--- Tomato #1: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #2: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #3: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #4: 25ns ---
[Sleep] 24ns
1ns remaining until 15ns break...
[Notify] 1ns remaining until 15ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 15ns
[LockScreen]
[Sleep] 15ns
Press <ENTER> to continue...
--- Tomato #5: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #6: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #7: 25ns ---
[Sleep] 24ns
1ns remaining until 5ns break...
[Notify] 1ns remaining until 5ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 5ns
[LockScreen]
[Sleep] 5ns
Press <ENTER> to continue...
--- Tomato #8: 25ns ---
[Sleep] 24ns
1ns remaining until 15ns break...
[Notify] 1ns remaining until 15ns break...
[Sleep] 1ns
[FocusApp] Terminal
Break: 15ns
[LockScreen]
[Sleep] 15ns
Press <ENTER> to continue...
`
