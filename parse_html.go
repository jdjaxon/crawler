package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	header := doc.Find("h1").Text()
	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	paragraph := doc.Find("main").Find("p").Text()
	if paragraph == "" {
		paragraph = doc.Find("p").Text()
	}

	return paragraph
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var links []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		// TODO: extract link instead of text
		relURL, err := url.Parse(s.Text())
		if err != nil {
			log.Printf("HTML URL parsing failed: %v\n", err)
			return
		}

		resolvedURL := baseURL.ResolveReference(relURL)
		links = append(links, resolvedURL.String())
	})

	return links, nil
}
