package main

import (
	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if _, lineerr := v.Line(cy + 2); lineerr != nil {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if _, lineerr := v.Line(cy - 2); lineerr != nil {
			if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
				if err := v.SetOrigin(ox, oy-1); err != nil {
					return err
				}
			}
			if err := v.SetCursor(cx, cy); err != nil && oy > 0 {
				if err := v.SetOrigin(ox, oy); err != nil {
					return err
				}
			}
			return nil
		}
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func backtoSearch(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(searchui)
	return nil
}

func backtoTorlist(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(searchlisttor)
	return nil
}

func doSearch(g *gocui.Gui, v *gocui.View) error {
	dstv, _ := g.View("inputbox")
	query = dstv.Buffer()
	err := torparadise(g)
	return err
}

func doAbout(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(aboutfunc)
	return nil
}
