package tomato

import (
	"fmt"
	"log"
	"os/exec"
)

type MacOS struct{}

func (this *MacOS) Notify(message string) {
	notification := fmt.Sprintf("display notification \"%s\" with title \"tomato Timer\"", message)
	this.execute(exec.Command("osascript", "-e", notification))
}

func (this *MacOS) FocusApp(name string) {
	this.execute(exec.Command("open", "-a", name))
}

func (this *MacOS) LockScreen() {
	this.execute(exec.Command("/System/Library/CoreServices/Menu Extras/User.menu/Contents/Resources/CGSession", "-suspend"))
}

func (this *MacOS) execute(command *exec.Cmd) {
	if output, err := command.CombinedOutput(); err != nil {
		log.Println("[WARN]", string(output), err)
	}
}
