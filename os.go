package tomato

import (
	"log"
	"os/exec"
)

type MacOS struct {
}

func (this *MacOS) FocusApp(name string) {
	err := exec.Command("open", "-a", "Terminal").Run()
	if err != nil {
		log.Println("[WARN]", err)
	}
}

func (this *MacOS) LockScreen() {
	err := exec.Command("/System/Library/CoreServices/Menu Extras/User.menu/Contents/Resources/CGSession", "-suspend").Run()
	if err != nil {
		log.Println("[WARN]", err)
	}
}
