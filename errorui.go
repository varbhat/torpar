package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func errorfunc(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	maxX, maxY := g.Size()

	// Error Widget
	if errorwid, err := g.SetView("errorwid", -1, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding("errorwid", gocui.KeyArrowLeft, gocui.ModNone, backtoSearch); err != nil {
			return err
		}

		if err := g.SetKeybinding("errorwid", gocui.KeyEnter, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("errorwid", 'q', gocui.ModNone, quit); err != nil {
			return err
		}
		if _, err := g.SetCurrentView("errorwid"); err != nil {
			return err
		}
		errorwid.Frame = true
		errorwid.Wrap = true
		fmt.Fprint(errorwid, term_red, "Error! ->", term_res, "\n\n")
		fmt.Fprintln(errorwid, string(errorui.Error()))
		errorwid.MoveCursor(maxX-1, 0, true)
	}
	// Help Widget
	if errorhelpwid, err := g.SetView("errorhelpwid", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		errorhelpwid.Frame = true
		errorhelpwid.Wrap = true
		fmt.Fprintln(errorhelpwid, "Press <enter> / ← to go back to Search , Ctrl-c / (q) to Quit")
	}
	return nil
}
