package main

import (
	"sync"
)

// Declare Global Variables
const magnetheader string = "magnet:?xt=urn:btih:"

var apistr string = "https://torrent-paradise.ml/api/search?q="
var trackerurl string        // tracker list in url
var torrents []Torrent       // Array of Torrent
var query string             // Query String
var trackerlisturl string    // Tracker list URL
var trackerquery string      // used in flag assignation
var selid int = -1           // Id of torrent selected
var errorui error            // error to be displayed in error ui
var datafilename string      // used in flag assignation
var datatype int             // used in flag assignation
var waitforme sync.WaitGroup // used in some functions to wait for gentorrents to complete

// Structure of Torrent
type Torrent struct {
	Id     string  `json:"id"`
	Text   string  `json:"text"`
	Length float64 `json:"len"`
	Seeds  int     `json:"s"`
	Leechs int     `json:"l"`
}

// color strings required to color text in terminal
var (
	term_red   string = "\u001b[0;31m" // Red color
	term_green string = "\u001b[0;32m" // Green color
	term_cyan  string = "\u001b[0;36m" // Cyan color
	term_blue  string = "\u001b[0;34m" // Blue color
	term_purp  string = "\u001b[0;35m" // Purple color
	term_yell  string = "\u001b[0;33m" // Yellow color
	term_res   string = "\u001b[0;00m" // Color reset
	term_grey  string = "\u001b[0;37m" // Color Grey
)
