package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	parsedPageURL, err := url.Parse(pageURL)
	if err != nil {
		log.Printf("Failed to parse a valid URL from %s: %v\n", pageURL, err)
	}

	urls, err := getURLsFromHTML(html, parsedPageURL)
	if err != nil {
		log.Printf("Failed to parse URLs from %s: %v\n", pageURL, err)
	}

	images, err := getImagesFromHTML(html, parsedPageURL)
	if err != nil {
		log.Printf("Failed to parse images from %s: %v\n", pageURL, err)
	}

	pageData := PageData{
		URL:            pageURL,
		H1:             getH1FromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  urls,
		ImageURLs:      images,
	}

	return pageData
}
