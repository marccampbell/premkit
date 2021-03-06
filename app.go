package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/premkit/premkit/commands"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()

	<-make(chan int)
}
