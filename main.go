package main

import (
	"github.com/taion809/bottest/cmd/bottest"
)

func main() {
	// trace.Start(os.Stderr)
	// defer trace.Stop()

	bottest.Execute()
}
