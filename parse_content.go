package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	header := doc.Find("h1").First().Text()
	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	paragraph := doc.Find("main").Find("p").Text()
	if paragraph == "" {
		paragraph = doc.Find("p").First().Text()
	}

	return paragraph
}
