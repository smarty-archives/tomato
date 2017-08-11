package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/smartystreets/tomato"
)

const maxTomatoesPerDay = 16

func main() {
	log.SetFlags(log.Ltime)
	log.SetOutput(os.Stdout)

	if runtime.GOOS != "darwin" {
		log.Fatal("This application only supports macos.")
	}

	tomato.NewController(os.Stdin, new(tomato.MacOS), maxTomatoesPerDay, time.Minute).Run()
}
