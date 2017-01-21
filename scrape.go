package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/cheggaaa/pb.v1"
)

var (
	comicTitle string
	comicPages []string
)

func GetHtml(URL string) (res *http.Response) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.94 Safari/537.36")

	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func run() {
	runPost(ehenPostURL)

	count := len(comicPages)
	bar := pb.StartNew(count)
	for i, c := range comicPages {
		runPage(c, i+1)
		bar.Increment()
	}
	bar.FinishPrint(fmt.Sprintf("%s download complete", comicTitle))
}

func runPost(URL string) {
	res := GetHtml(URL)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatalln(err)
	}

	comicTitle = doc.Find("#gd2 #gj").Text()
	if len(comicTitle) == 0 {
		comicTitle = doc.Find("#gd2 #gn").Text()
	}

	doc.
		Find("#gdt .gdtm").
		Each(func(i int, this *goquery.Selection) {
			page, _ := this.Find("div").Find("a").Attr("href")
			comicPages = append(comicPages, page)
		})

	nextPage, exist := doc.Find("table.ptb tr td").Last().Find("a").Attr("href")
	if exist {
		runPost(nextPage)
	}
}

func runPage(URL string, index int) {
	res := GetHtml(URL)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatalln(err)
	}

	imageLink, _ := doc.Find("#img").Attr("src")
	downloadImage(imageLink, index)
}

func downloadImage(URL string, index int) {
	res := GetHtml(URL)
	defer res.Body.Close()

	// create folder
	var destination string
	if len(folder) > 0 {
		destination = fmt.Sprintf("%s/%s", folder, comicTitle)
	} else {
		destination = comicTitle
	}
	os.MkdirAll(destination, 0775)

	// create file
	file, err := os.Create(fmt.Sprintf("%s/%d.jpg", destination, index))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// copy body to file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
}
