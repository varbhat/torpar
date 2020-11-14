package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"os"
)

func main() {
	// Generate tracker part of Magnet URL in another goroutine
	go gentrackers()

	// Start our TUI - new gocui instance
	tui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tui.Close()

	tui.Cursor = true // Enable cursor
	if len(query) > 0 {
		if err := torparadise(tui); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		tui.SetManagerFunc(searchui) // Layout is handled by Mainlayout function
	}

	if err := tui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := tui.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		os.Exit(1)
	}
	tui.Close()

}
