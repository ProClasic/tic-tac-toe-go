package utils

import (
	"os"
	"time"
)

func Clrscr() {
	// var clrcmd *exec.Cmd = exec.Command("clear")
	// clrcmd.Stdout = os.Stdout
	// clrcmd.Run()
	time.Sleep(time.Millisecond * 10)
	os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})
}
