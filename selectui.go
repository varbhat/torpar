package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/cheynewallace/tabby"
	"github.com/jroimartin/gocui"
	"github.com/skratchdot/open-golang/open"
	"strconv"
	"strings"
	"text/tabwriter"
)

func selectfunc(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	maxX, maxY := g.Size()

	// Select Widget
	if selwid, err := g.SetView("selwid", -1, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		selwid.Wrap = true

		if err := g.SetKeybinding("selwid", gocui.KeySpace, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("selwid", gocui.KeyEnter, gocui.ModNone, openmagnet); err != nil {
			return err
		}
		if err := g.SetKeybinding("selwid", gocui.KeyArrowRight, gocui.ModNone, openmagnet); err != nil {
			return err
		}
		if err := g.SetKeybinding("selwid", gocui.KeyArrowLeft, gocui.ModNone, backtoTorlist); err != nil {
			return err
		}
		if err := g.SetKeybinding("selwid", gocui.KeySpace, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("selwid", 'q', gocui.ModNone, quit); err != nil {
			return err
		}
		if _, err := g.SetCurrentView("selwid"); err != nil {
			return err
		}

		selwid.Frame = true
		selwid.Wrap = true
		selwid.MoveCursor(maxX-1, 0, true)

		magnet := magnetheader + torrents[selid].Id + trackerurl

		fmt.Fprintln(selwid, term_yell+"Torrent Details for Torrent ", selid+1, " →"+term_res)
		tw := tabwriter.NewWriter(selwid, 0, 0, 2, ' ', 0)
		t := tabby.NewCustom(tw)
		t.AddLine(term_res+"\u2022 Name"+term_cyan, torrents[selid].Text)
		t.AddLine(term_res+"\u2022 Size"+term_purp, torrents[selid].Length)
		t.AddLine(term_res+"\u2022 Leechs"+term_red, torrents[selid].Leechs)
		t.AddLine(term_res+"\u2022 Infohash"+term_cyan, torrents[selid].Leechs)
		t.AddLine(term_res+"\u2022 Magnet"+term_green, magnetheader+torrents[selid].Id)
		t.AddLine(term_res+"\u2022 Trackerlist URL"+term_purp, trackerlisturl)
		t.AddLine(term_res, term_res)
		t.Print()
		fmt.Fprintln(selwid, term_cyan, "Press <enter> to Open this Magnet link in Preconfigured Application\n", term_res)

		// Write to clipboard
		cliperr := clipboard.WriteAll(magnet)
		if cliperr != nil {
			fmt.Fprintln(selwid, term_red, "Error copying Magnet link :", cliperr.Error(), term_res)
		} else {
			fmt.Fprintln(selwid, term_green, "Magnet Link with added trackers has been copied to your device's clipboard\n\n", term_res)
		}

	}

	if selhelp, err := g.SetView("selhelp", -1, maxY-2, maxX, maxY); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}

		selhelp.Frame = false
		selhelp.Wrap = true
		fmt.Fprintf(selhelp, "Press <enter> or → to Open , ← to go back , <space> to go to Search, (q) or Ctrl-c to quit")

	}

	return nil
}

func openmagnet(g *gocui.Gui, v *gocui.View) error {
	err := open.Run(magnetheader + torrents[selid].Id + trackerurl)
	selwid, _ := g.View("selwid")
	if err != nil {
		fmt.Fprintln(selwid, term_red, "Error opening the magnet: ", err.Error(), term_res)
		return nil
	} else {
		fmt.Fprintln(selwid, term_green, "Torrent has been opened in the configured Application")
	}
	return nil
}

func selectit(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	selline, err := v.Line(cy)
	selArr := strings.Fields(selline)
	if len(selArr) > 0 {
		if selArr[0] == "#" {
			return nil
		}
		selid, err = strconv.Atoi(selArr[0])
		if err != nil {
			return nil
		}
		selid = selid - 1

	} else {
		return nil
	}
	if len(datafilename) > 0 {
		switch datatype {
		case 4, 5:
			g.SetManagerFunc(torrentdatawrite)
			return nil
		default:
			g.SetManagerFunc(selectfunc)
			return nil

		}
	}
	g.SetManagerFunc(selectfunc)

	return nil
}
