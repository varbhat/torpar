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
		csvw.Write([]string{"Id", "Name", "Size", "Seeds", "Leeches", "Infohash", "Magnet"})
		for cidno, ceacht := range torrents {
			csvw.Write(
				[]string{
					strconv.Itoa(cidno),
					ceacht.Text,
					fmt.Sprintf("%f", ceacht.Length),
					strconv.Itoa(ceacht.Seeds),
					strconv.Itoa(ceacht.Leechs),
					ceacht.Id,
					magnetheader + ceacht.Id + trackerurl,
				})
		}
		csvw.Flush() // Flushing is necessary when csv.Writer.Write() is used
		if csverr := csvw.Error(); err != nil {
			errorui = csverr
			g.SetManagerFunc(errorfunc)
		}

	case 3:
		// Write Magnet Links to the file. That's it
		for _, ceacht := range torrents {
			if _, fprinterr := fmt.Fprintln(datafile, magnetheader+ceacht.Id+trackerurl); err != nil {
				errorui = fprinterr
				g.SetManagerFunc(errorfunc)
			}
		}

	case 4:
		// Write magnet link of selected torrent
		if _, fprinterr := fmt.Fprintln(datafile, magnetheader+torrents[selid].Id+trackerurl); err != nil {
			errorui = fprinterr
			g.SetManagerFunc(errorfunc)
		}
	case 5:
		// Write Torrent details of selected torrent
		if _, fprinterr := fmt.Fprint(datafile, "Name: "+torrents[selid].Text, "\nSize: ", torrents[selid].Length, "\nSeeds: ", torrents[selid].Seeds, "\nLeechs: ", torrents[selid].Leechs, "\nInfohash: ", torrents[selid].Id, "\n\nMagnet: \n\n", magnetheader+torrents[selid].Id+trackerurl); err != nil {
			errorui = fprinterr
			g.SetManagerFunc(errorfunc)
		}
	default:
		// Write CSV data of torrent to the file
		csvw := csv.NewWriter(datafile)
		csvw.Write([]string{"Id", "Name", "Size", "Seeds", "Leeches", "Infohash", "Magnet"})
		for cidno, ceacht := range torrents {
			csvw.Write(
				[]string{
					strconv.Itoa(cidno),
					ceacht.Text,
					fmt.Sprintf("%f", ceacht.Length),
					strconv.Itoa(ceacht.Seeds),
					strconv.Itoa(ceacht.Leechs),
					ceacht.Id,
					magnetheader + ceacht.Id,
				})
		}
		csvw.Flush() // Flushing is necessary when csv.Writer.Write() is used
		if csverr := csvw.Error(); err != nil {
			errorui = csverr
			g.SetManagerFunc(errorfunc)
		}
	}

	g.Close()
	fmt.Println("Data has been written to file ", datafilename)
	os.Exit(1)
	return gocui.ErrQuit
}
