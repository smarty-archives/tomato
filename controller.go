package tomato

import (
	"fmt"
	"io"
	"log"
	"time"
)

type System interface {
	Notify(message string)
	FocusApp(name string)
	LockScreen()
	Sleep(time.Duration)
}

type Controller struct {
	terminal io.Reader
	system   System

	maxTomatoCount int
	tomatoCount    int

	tomato     time.Duration
	coolDown   time.Duration
	shortBreak time.Duration
	longBreak  time.Duration
}

func NewController(terminal io.Reader, system System, sessions int, scale time.Duration) *Controller {
	return &Controller{
		terminal:       terminal,
		system:         system,
		maxTomatoCount: sessions,

		tomato:     scale * 24,
		coolDown:   scale,
		shortBreak: scale * 5,
		longBreak:  scale * 15,
	}
}

func (this *Controller) Run() {
	for this.tomatoCount = 1; this.tomatoCount <= this.maxTomatoCount; this.tomatoCount++ {
		this.doTomato()
		this.takeBrake()
		this.prepareNextTomato()
	}
}

func (this *Controller) doTomato() {
	log.Printf("--- Tomato #%d: %v ---", this.tomatoCount, this.tomato+this.coolDown)
	this.system.Sleep(this.tomato)
	this.doCoolDown()
}
func (this *Controller) doCoolDown() {
	soon := fmt.Sprintf("%s remaining until %s break...", this.coolDown, this.breakDuration())
	log.Println(soon)
	this.system.Notify(soon)
	this.system.Sleep(this.coolDown)
}

func (this *Controller) breakDuration() time.Duration {
	if this.tomatoCount%4 == 0 {
		return this.longBreak
	} else {
		return this.shortBreak
	}
}

func (this *Controller) takeBrake() {
	duration := this.breakDuration()
	this.system.FocusApp("Terminal")
	log.Printf("Break: %v", duration)
	this.system.LockScreen()
	this.system.Sleep(duration)
}

func (this *Controller) prepareNextTomato() {
	log.Println("Press <ENTER> to continue...")
	this.awaitEnterKeyPress()
}

func (this *Controller) awaitEnterKeyPress() {
	fmt.Fscanln(this.terminal)
}
