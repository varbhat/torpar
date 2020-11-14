package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Fetch List of Best Public Torrent Trackers Available and create trackerlist out of it.
func gentrackers() {
	// Different URLs and options to fetch trackers from
	if trackerquery != "" {
		switch trackerquery {
		case "tl":
			trackerlisturl = "https://ngosang.github.io/trackerslist/trackers_all.txt"
			break
		case "nt":
			trackerlisturl = "https://newtrackon.com/api/stable"
		default:
			trackerlisturl = trackerquery
		}
	} else {
		trackerlisturl = "https://newtrackon.com/api/stable"
	}
	if trackerurl != "" {
		return
	}

	trackersresponse, err := http.Get(trackerlisturl)
	if err != nil {
		return
	}
	trackers, err := ioutil.ReadAll(trackersresponse.Body)
	if err != nil {
		return
	}
	var trackerlist strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(string(trackers)))
	for scanner.Scan() {
		trackerlist.WriteString("&tr=" + url.QueryEscape(scanner.Text()))
	}
	trackerurl = trackerlist.String()
	waitforme.Done()
}
