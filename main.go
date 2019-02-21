package main

import (
	"os"
	"os/exec"
	"time"
	"github.com/fatih/color"
)

func main() {
	start := time.Now()
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	elapsed := time.Since(start)
	color.New(color.FgMagenta).Printf("Time Elapsed: %s", elapsed)
}
