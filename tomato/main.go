package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/smartystreets/tomato"
)

func main() {
	log.SetFlags(log.Ltime)
	log.SetOutput(os.Stdout)

	if runtime.GOOS != "darwin" {
		log.Fatal("This application only supports macos.")
	}

	tomato.NewController(os.Stdin, &tomato.MacOS{}, 32, time.Minute, time.Sleep).Run()
}
