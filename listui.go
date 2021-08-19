package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	"github.com/jroimartin/gocui"
)

func maxstring(str string, length int) string {
	strlen := len(str)
	if strlen > length {
		return str[:length]
	} else {
		return str
	}
}

func torparadise(g *gocui.Gui) error {

	if strings.TrimSpace(query) == "" {
		errorui = errors.New("No Query")
		g.SetManagerFunc(errorfunc)
		return nil
	}

	// construct API URL
	finalquery := apistr + url.QueryEscape(query)

	// Query the Torrent Paradise API to get result as response
	response, err := http.Get(finalquery)
	if err != nil {
		errorui = err
		g.SetManagerFunc(errorfunc)
		return nil
	}

	// Read JSON Response
	respbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorui = err
		g.SetManagerFunc(errorfunc)
		return nil
	}

	// Parse the json response into []Torrent
	err = json.Unmarshal([]byte(respbody), &torrents)
	if err != nil {
		errorui = err
		g.SetManagerFunc(errorfunc)

	}

	if len(torrents) == 0 {
		errorui = errors.New("No Results")
		g.SetManagerFunc(errorfunc)
		return nil
	}
	query = ""

	if len(datafilename) > 0 {
		switch datatype {
		case 4, 5:
			g.SetManagerFunc(searchlisttor)
		default:
			g.SetManagerFunc(torrentdatawrite)
		}
	} else {
		g.SetManagerFunc(searchlisttor)
	}
	return nil
}

func searchlisttor(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	maxX, maxY := g.Size()

	// Torrent List Box
	if tllist, err := g.SetView("tl", -1, -1, maxX, maxY-2); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}

		tllist.Highlight = true
		tllist.SelBgColor = gocui.ColorBlue
		tllist.SelFgColor = gocui.ColorWhite
		tllist.Wrap = true
		tllist.MoveCursor(maxX-1, 1, true)

		if err := g.SetKeybinding("tl", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
			return err
		}
		if err := g.SetKeybinding("tl", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
			return err
		}
		if err := g.SetKeybinding("tl", gocui.KeyArrowLeft, gocui.ModNone, backtoSearch); err != nil {
			return err
		}
		if err := g.SetKeybinding("tl", gocui.KeyEnter, gocui.ModNone, selectit); err != nil {
			return err
		}
		if err := g.SetKeybinding("tl", gocui.KeyArrowRight, gocui.ModNone, selectit); err != nil {
			return err
		}
		if err := g.SetKeybinding("tl", 'q', gocui.ModNone, quit); err != nil {
			return err
		}
		var cv *gocui.View
		if cv, err = g.SetCurrentView("tl"); err != nil {
			return err
		}

		if len(torrents) != 0 {
			maxlentl := int(float64(maxX) / 1.35)
			tw := tabwriter.NewWriter(tllist, 0, 0, 1, ' ', 0)
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

		for mc := 0; mc <= selid+1; mc++ {
			if ccerr := cv.SetCursor(0, mc); ccerr != nil {
				ox, oy := cv.Origin()
				if soerr := cv.SetOrigin(ox, oy+1); soerr != nil {
					return soerr
				}
			}
		}
	}

	if tlhelp, err := g.SetView("tlhelp", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		tlhelp.Frame = false
		tlhelp.Wrap = true
		fmt.Fprintf(tlhelp, "use ↑ ↓ to navigate, <enter> or → to select , ← to go back to search , (q) or Ctrl-c to quit")

	}

	return nil
}
