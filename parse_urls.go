package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, sel *goquery.Selection) {
		href, ok := sel.Attr("href")
		if !ok {
			return
		}

		relURL, err := url.Parse(href)
		if err != nil {
			log.Printf("Failed to parse href: %v\n", err)
			return
		}

		absoluteURL := baseURL.ResolveReference(relURL)
		urls = append(urls, absoluteURL.String())
	})

	return urls, nil
}
