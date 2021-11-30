package utils

import (
	"os"
	"os/exec"
	"time"
)

func Clrscr() {
	var clrcmd *exec.Cmd = exec.Command("clear")
	clrcmd.Stdout = os.Stdout
	clrcmd.Run()
	time.Sleep(time.Millisecond * 10)
}
