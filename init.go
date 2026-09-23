package main

import (
	"errors"
	"os"

	flag "github.com/spf13/pflag"
)

func init() {
	searchquery := flag.StringP("query", "q", "", "Search Query")
	trquery := flag.StringP("tlist", "l", "", "URL to tracker list\n (default https://newtrackon.com/api/stable)")
	datafilequery := flag.StringP("file", "f", "", "File Path to write Data(csv) into")
	datatypequery := flag.IntP("type", "t", 1, "Type of Data(csv) to write into file\n 1 → Results\n 2 → Results with Magnet links (with trackers)\n 3 → Only Magnet links of Results\n 4 → Magnet link of Selected Torrent\n 5 → Data of Selected Torrent\n")
	flag.ErrHelp = errors.New("  -h, --help            Help\n\n" + os.Args[0] + " - Knaben torrent search TUI client.\nSource: https://github.com/varbhat/torpar\nAPI: https://api.knaben.org/v1")
	mainapiquery := flag.StringP("apiurl", "a", "", "API Endpoint URL\n (default https://api.knaben.org/v1)")

	flag.Parse()
	query = *searchquery
	trackerquery = *trquery
	datafilename = *datafilequery
	datatype = *datatypequery

	// If specified , change API Endpoint
	if *mainapiquery != "" {
		apistr = *mainapiquery
	}
	waitforme.Add(1)
}
