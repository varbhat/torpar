<h1 align="center">torpar</h1> 
<p align="center">TUI Client for <a href="https://knaben.org">Knaben Database</a></p>

## Introduction
`torpar` is a TUI client for [Knaben Database](https://knaben.org), a torrent search engine aggregating multiple indexers via Elasticsearch.

## Installation

You can download binary for your OS from [Releases](https://github.com/varbhat/torpar/releases/latest) . Also , if you have [Go](https://golang.org/) installed , you can install `torpar` by typing this in terminal.

```bash
go install github.com/varbhat/torpar@latest
```

## Features

* Search the [Knaben Database](https://knaben.org)
* Navigate the Search Results
* Get Details of the Search Result
* Copy Magnet Link of Search Result into Device Clipboard
* Open Magnet Link in Preconfigured Application 
* Write Data to File
* Navigate TUI interface very easily
* Supports Most Platforms and Terminals (including tty)

## Usage

```
Usage of torpar:
  -a, --apiurl string   API Endpoint URL
                         (default https://api.knaben.org/v1)
  -f, --file string     File Path to write Data(csv) into
  -q, --query string    Search Query
  -l, --tlist string    URL to tracker list
                         (default https://newtrackon.com/api/stable)
  -t, --type int        Type of Data(csv) to write into file
                         1 → Results
                         2 → Results with Magnet links (with trackers)
                         3 → Only Magnet links of Results
                         4 → Magnet link of Selected Torrent
                         5 → Data of Selected Torrent
                         (default 1)
  -h, --help            Help
```

`torpar` is TUI and intuitive to understand . 

`torpar` also shows every navigation keys at each TUI views so that you can use `torpar` very easily.

## License
[GPL-v3](LICENSE)
