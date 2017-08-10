package tomato

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TimeCounter interface {
	Countdown(duration time.Duration) chan time.Time
}

type OS interface {
	Notify(title, message string)
	FocusApp(name string)
	LockScreen()
}

type Controller struct {
	terminal  io.ReadWriter
	timer     TimeCounter
	os        OS
	interrupt chan os.Signal
	gameOver  bool
	duration  time.Duration
}

func NewController(terminal io.ReadWriter, timer TimeCounter, system OS, signals chan os.Signal) *Controller {
	return &Controller{
		terminal:  terminal,
		timer:     timer,
		os:        system,
		interrupt: signals,
	}
}

func (this *Controller) Run() {
	for tomato := 1; !this.gameOver; tomato++ {
		this.doTomato()
		this.takeBreak(tomato)
		this.prepareForNextTomato()
	}
}

func (this *Controller) doTomato() {
	this.Println("Setting timer for 25 minutes...Go!")
	this.countdown(time.Minute * 24)
	this.Println("1 minute remaining...")
	this.os.Notify("Tomato Timer", "1 minute remaining...")
	this.countdown(time.Minute)
}

func (this *Controller) takeBreak(tomatoCount int) {
	this.os.FocusApp("Terminal")
	if tomatoCount%4 == 0 {
		this.takeLongerBreak()
	} else {
		this.takeShorterBreak()
	}
}
func (this *Controller) takeLongerBreak() {
	this.executeBreak(time.Minute*15, "Time's up! Go take a longer, 15 minute, break. Locking screen...")
}
func (this *Controller) takeShorterBreak() {
	this.executeBreak(time.Minute*5, "Time's up! Go take a 5 minute break. Locking screen...")
}
func (this *Controller) executeBreak(duration time.Duration, message string) {
	this.Println(message)
	this.countdown(time.Second * 5)
	this.os.LockScreen()
	this.countdown(duration)
}
func (this *Controller) prepareForNextTomato() {
	this.Println("<ENTER> to start a new 25 minute tomato...")
	if !this.gameOver {
		fmt.Fscanln(this.terminal)
		fmt.Fprintln(this.terminal)
	}
}

func (this *Controller) Println(message string) {
	if this.gameOver {
		return
	}
	fmt.Fprintln(this.terminal, message)
}

func (this *Controller) countdown(duration time.Duration) {
	if this.gameOver {
		return
	}
	select {
	case <-this.timer.Countdown(duration):
		return
	case <-this.interrupt:
		this.gameOver = true
	}
}
