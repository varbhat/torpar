package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func aboutfunc(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	if err := g.SetKeybinding("aboutwid", gocui.KeyEsc, gocui.ModNone, backtoSearch); err != nil {
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

		abouttptext := `• TorPar is TUI client for Torrent Paradise
• Torrent Paradise is Open-Source DHT Torrent Search Engine
• TorPar is FLOSS and is licensed under GPLv3
• Source code at https://github.com/varbhat/torpar`

		fmt.Fprint(aboutwid, abouttptext)

		fmt.Fprint(aboutwid, term_green, "\n\nAbout Torrent Paradise Search Engine ->\n", term_res)

		abouttpstext := `• Fresh and rich torrent index
• New torrents identified quickly via multiple RSS feeds
• Obscure torrents discovered through DHT
• Seed/Leech counts constantly refreshed
• privacy preserving, not-in-your-face ads
• donate and vote on future features
• Source Code at https://github.com/urbanguacamole/torrent-paradise
• Send suggestions to urban-guacamole (at) protonmail.com
• Want to report a copyright violation? See copyright at https://torrent-paradise.ml/copyright.html`
		fmt.Fprint(aboutwid, abouttpstext)

		aboutwid.MoveCursor(maxX-1, 0, true)
	}
	// Help Widget
	if abouthelpwid, err := g.SetView("abouthelpwid", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		abouthelpwid.Frame = true
		abouthelpwid.Wrap = true
		fmt.Fprintln(abouthelpwid, "Press <esc> / <enter> / ← to go back to Search , Ctrl-c / (q) to Quit")
	}
	return nil
}
