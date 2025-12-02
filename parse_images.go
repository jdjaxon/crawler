package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Printf("HTML parsing failed: %v\n", err)
	}

	var imgURLs []string
	doc.Find("img").Each(func(_ int, sel *goquery.Selection) {
		src, ok := sel.Attr("src")
		if !ok {
			return
		}

		relURL, err := url.Parse(src)
		if err != nil {
			log.Printf("Failed to parse href: %v\n", err)
			return
		}

		absoluteURL := baseURL.ResolveReference(relURL)
		imgURLs = append(imgURLs, absoluteURL.String())
	})

	return imgURLs, nil
}
