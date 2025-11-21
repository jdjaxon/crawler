package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		log.Printf("error parsing URL: %s", err)
		return "", err
	}

	constructedURL := fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.Path)
	normalizedURL := strings.TrimSuffix(constructedURL, "/")

	return normalizedURL, nil
}
