package main

import (
	"fmt"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
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

func sortAsc(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := g.Size()
	sort.Sort(AscTorrents(torrents))
	v.Clear()
	if len(torrents) != 0 {
		maxlentl := int(float64(maxX) / 1.35)
		tw := tabwriter.NewWriter(v, 0, 0, 1, ' ', 0)
		t := tabby.NewCustom(tw)
		t.AddLine(term_res+"#", term_res+"Name", term_res+"Size", term_res+"Seeds", term_res+"Leeches"+term_res)
		for idno, eachtorrent := range torrents {
			t.AddLine(term_yell+strconv.Itoa(idno+1),
				term_cyan+maxstring(eachtorrent.Text, maxlentl),
				term_purp+fmt.Sprintf("%f", 0.000000001*eachtorrent.Length)+" GB",
				term_green+"S:"+strconv.Itoa(eachtorrent.Seeds),
				term_red+"L:"+strconv.Itoa(eachtorrent.Leechs)+term_res)
		}
		t.Print()
	}
	return nil
}

func sortDec(g *gocui.Gui, v *gocui.View) error {
	maxX, _ := g.Size()
	sort.Sort(DeTorrents(torrents))
	v.Clear()
	if len(torrents) != 0 {
		maxlentl := int(float64(maxX) / 1.35)
		tw := tabwriter.NewWriter(v, 0, 0, 1, ' ', 0)
		t := tabby.NewCustom(tw)
		t.AddLine(term_res+"#", term_res+"Name", term_res+"Size", term_res+"Seeds", term_res+"Leeches"+term_res)
		for idno, eachtorrent := range torrents {
			t.AddLine(term_yell+strconv.Itoa(idno+1),
				term_cyan+maxstring(eachtorrent.Text, maxlentl),
				term_purp+fmt.Sprintf("%f", 0.000000001*eachtorrent.Length)+" GB",
				term_green+"S:"+strconv.Itoa(eachtorrent.Seeds),
				term_red+"L:"+strconv.Itoa(eachtorrent.Leechs)+term_res)
		}
		t.Print()
	}
	return nil
}
