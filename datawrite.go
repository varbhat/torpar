package main

import (
	"encoding/csv" // csv encoding support
	"fmt"
	"os"
	"strconv"

	"github.com/jroimartin/gocui"
)

func torrentdatawrite(g *gocui.Gui) error {

	// Try to open specified file
	datafile, err := os.Create(datafilename)
	if err != nil {
		errorui = err
		g.SetManagerFunc(errorfunc)
	}
	defer datafile.Close()

	waitforme.Wait() // Wait for gentrackers to fetch trackers from trackerlist

	switch datatype {

	case 2:
		// Write CSV with torrent details(with magnet with trackers) to the file
		csvw := csv.NewWriter(datafile)
		csvw.Write([]string{"Id", "Name", "Size", "Seeders", "Peers", "Infohash", "Magnet"})
		for cidno, ceacht := range torrents {
			csvw.Write(
				[]string{
					strconv.Itoa(cidno),
					ceacht.Title,
					fmt.Sprintf("%d", ceacht.Bytes),
					strconv.Itoa(ceacht.Seeders),
					strconv.Itoa(ceacht.Peers),
					ceacht.Hash,
					magnetFor(ceacht),
				})
		}
		csvw.Flush()
		if csverr := csvw.Error(); csverr != nil {
			errorui = csverr
			g.SetManagerFunc(errorfunc)
		}

	case 3:
		// Write Magnet Links to the file. That's it
		for _, ceacht := range torrents {
			if _, fprinterr := fmt.Fprintln(datafile, magnetFor(ceacht)); fprinterr != nil {
				errorui = fprinterr
				g.SetManagerFunc(errorfunc)
			}
		}

	case 4:
		// Write magnet link of selected torrent
		if _, fprinterr := fmt.Fprintln(datafile, magnetFor(torrents[selid])); fprinterr != nil {
			errorui = fprinterr
			g.SetManagerFunc(errorfunc)
		}
	case 5:
		// Write Torrent details of selected torrent
		if _, fprinterr := fmt.Fprint(datafile, "Name: "+torrents[selid].Title, "\nSize: ", torrents[selid].Bytes, "\nSeeders: ", torrents[selid].Seeders, "\nPeers: ", torrents[selid].Peers, "\nInfohash: ", torrents[selid].Hash, "\n\nMagnet: \n\n", magnetFor(torrents[selid])); fprinterr != nil {
			errorui = fprinterr
			g.SetManagerFunc(errorfunc)
		}
	default:
		// Write CSV data of torrent to the file
		csvw := csv.NewWriter(datafile)
		csvw.Write([]string{"Id", "Name", "Size", "Seeders", "Peers", "Infohash", "Magnet"})
		for cidno, ceacht := range torrents {
			csvw.Write(
				[]string{
					strconv.Itoa(cidno),
					ceacht.Title,
					fmt.Sprintf("%d", ceacht.Bytes),
					strconv.Itoa(ceacht.Seeders),
					strconv.Itoa(ceacht.Peers),
					ceacht.Hash,
					magnetFor(ceacht),
				})
		}
		csvw.Flush()
		if csverr := csvw.Error(); csverr != nil {
			errorui = csverr
			g.SetManagerFunc(errorfunc)
		}
	}

	g.Close()
	fmt.Println("Data has been written to file ", datafilename)
	os.Exit(1)
	return gocui.ErrQuit
}
