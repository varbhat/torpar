package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func searchui(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Println(err)
	}

	// Header
	if header, err := g.SetView("header", maxX/2-9, maxY/2-6, maxX/2+9, maxY/2-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		header.Frame = false
		fmt.Fprintln(header, "   Search the")
		fmt.Fprintln(header, term_yell, "Torrent Paradise", term_res)
	}

	// Input Box
	if editbox, err := g.SetView("inputbox", maxX/2-17, maxY/2-2, maxX/2+17, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		editbox.Editable = true

		if err := g.SetKeybinding("inputbox", gocui.KeyEnter, gocui.ModNone, doSearch); err != nil {
			return err
		}

		if err := g.SetKeybinding("inputbox", gocui.KeyArrowRight, gocui.ModNone, doSearch); err != nil {
			return err
		}

		if err := g.SetKeybinding("inputbox", gocui.KeyHome, gocui.ModNone, doAbout); err != nil {
			return err
		}

		if err := g.SetKeybinding("inputbox", gocui.KeyDelete, gocui.ModNone, resetInput); err != nil {
			return err
		}
		if err := g.SetKeybinding("inputbox", gocui.KeyArrowLeft, gocui.ModNone, quit); err != nil {
			return err
		}
		if _, err := g.SetCurrentView("inputbox"); err != nil {
			return err
		}
	}

	// Enter to Search View
	if srview, err := g.SetView("srview", maxX/2-9, maxY/2+1, maxX/2+9, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		srview.Frame = false
		fmt.Fprintln(srview, term_grey, "Enter to Search", term_res)
	}
	// Help Widget
	if srhelp, err := g.SetView("srhelp", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		srhelp.Frame = true
		srhelp.Wrap = true
		fmt.Fprintln(srhelp, "Press <enter> / → to Search , ← / Ctrl-c  to Quit , <home> for About Menu")

	}

	return nil
}

func resetInput(g *gocui.Gui, v *gocui.View) error {
	inpbox, _ := g.View("inputbox")
	inpbox.Clear()
	inpbox.SetCursor(0, 0)
	inpbox.SetOrigin(0, 0)
	return nil
}
