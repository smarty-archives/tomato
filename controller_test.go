package tomato

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestController(t *testing.T) {
	output := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(output)
	tomatoes := 8
	sleeper := func(duration time.Duration) { log.Println("[Sleep]", duration) }

	controller := NewController(
		strings.NewReader(strings.Repeat("\n", tomatoes)),
		new(NopSystem),
		tomatoes,
		time.Nanosecond,
		sleeper,
	)

	controller.Run()
	assertions.New(t).So(output.String(), should.Equal, expectedTomatoSession)
}

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

///////////////////////////////////////////////////

type NopSystem struct{}

func (NopSystem) Notify(message string) { log.Println("[Notify]", message) }
func (NopSystem) FocusApp(name string)  { log.Println("[FocusApp]", name) }
func (NopSystem) LockScreen()           { log.Println("[LockScreen]") }
