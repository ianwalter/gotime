package main

import (
	"os"
	"os/exec"
	"time"
	"github.com/fatih/color"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	start := time.Now()
	cmd.Run()
	elapsed := time.Since(start)
	color.New(color.FgMagenta).Printf("Time Elapsed: %s", elapsed)
}
