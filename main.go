package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	ehenPostURL string
	folder      string
	checkURL    = "e-hentai.org/g/"
	test        = "http://g.e-hentai.org/g/1020609/7df75786ee/"
)

const (
	usage = `
Usage: go-Ehendown [options]

Options:
    -u, --URL <URL>             e-hentai post URL *Require
    -o, --out <path>            destination folder path
    -h, --help                  Show usage
`
)

func init() {
	flag.StringVar(&ehenPostURL, "u", "", "e-hentai post URL")
	flag.StringVar(&ehenPostURL, "URL", "", "e-hentai post URL")
	flag.StringVar(&folder, "o", "", "destination folder path")
	flag.StringVar(&folder, "out", "", "destination folder path")

	flag.Usage = func() {
		fmt.Printf("%s\n", usage)
		os.Exit(0)
	}
	flag.Parse()
}

func main() {
	if len(ehenPostURL) == 0 {
		flag.Usage()
	} else if strings.Contains(ehenPostURL, checkURL) == false {
		fmt.Println("Please key-in URL like as https://e-hentai.org/g/831308/855ebdd842/")
		os.Exit(0)
	}

	run()
}
