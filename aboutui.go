package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func aboutfunc(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	maxX, maxY := g.Size()

	// About Widget
	if aboutwid, err := g.SetView("aboutwid", -1, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetKeybinding("aboutwid", gocui.KeyArrowLeft, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("aboutwid", gocui.KeyEnter, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("aboutwid", gocui.KeyHome, gocui.ModNone, backtoSearch); err != nil {
			return err
		}

		if err := g.SetKeybinding("aboutwid", 'q', gocui.ModNone, quit); err != nil {
			return err
		}
		if _, err := g.SetCurrentView("aboutwid"); err != nil {
			return err
		}
		aboutwid.Frame = true
		aboutwid.Wrap = true
		fmt.Fprint(aboutwid, term_green, "About! ->\n\n", term_res)
		fmt.Fprint(aboutwid, term_green+"About Torpar ->\n"+term_res)

		abouttptext := `• TorPar is a TUI client for searching torrents
• Now powered by the Knaben Database API
• TorPar is FLOSS and is licensed under GPLv3
• Source code at https://github.com/varbhat/torpar`

		fmt.Fprint(aboutwid, abouttptext)

		fmt.Fprint(aboutwid, term_green, "\n\nAbout Knaben Database ->\n", term_res)

		aboutknabentext := `• Aggregates torrents from multiple indexers
• Fast Elasticsearch-backed search
• API docs at https://knaben.org/api/v1/
• Website at https://knaben.org`
		fmt.Fprint(aboutwid, aboutknabentext)

		aboutwid.MoveCursor(maxX-1, 0, true)
	}
	// Help Widget
	if abouthelpwid, err := g.SetView("abouthelpwid", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		abouthelpwid.Frame = true
		abouthelpwid.Wrap = true
		fmt.Fprintln(abouthelpwid, "Press <enter> / ← to go back to Search , Ctrl-c / (q) to Quit")
	}
	return nil
}
