package tomato

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type GenericOS struct{}

func (this *GenericOS) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

type MacOS struct{ *GenericOS }

func (this *MacOS) Notify(message string) {
	notification := fmt.Sprintf("display notification \"%s\" with title \"tomato Timer\"", message)
	this.execute(exec.Command("osascript", "-e", notification))
}

func (this *MacOS) FocusApp(name string) {
	this.execute(exec.Command("open", "-a", name))
}

func (this *MacOS) LockScreen() {
	const screenSaver = "/System/Library/Frameworks/ScreenSaver.framework/Versions/A/Resources/ScreenSaverEngine.app"
	this.FocusApp(screenSaver)
}

func (this *MacOS) execute(command *exec.Cmd) {
	if output, err := command.CombinedOutput(); err != nil {
		log.Println("[WARN]", string(output), err)
	}
}
