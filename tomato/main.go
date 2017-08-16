package main

import (
	"log"
	"os"
	"time"

	"github.com/smartystreets/tomato"
)

const maxTomatoesPerDay = 16

func main() {
	log.SetFlags(log.Ltime)
	log.SetOutput(os.Stdout)

	tomato.NewController(os.Stdin, new(tomato.MacOS), maxTomatoesPerDay, time.Minute).Run()
}
