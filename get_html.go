package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	// I need to create a request here so I can add the User-Agent
	// to the header.
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		return "", err
	}
	req.Header.Set("User-Agent", "Crawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request failed: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return "", fmt.Errorf("%d error", resp.StatusCode)
	}

	respContentType := resp.Header.Get("content-type")
	if respContentType != "text/html" {
		return "", fmt.Errorf("invalid content-type: %s", respContentType)
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response: %v", err)
		return "", err
	}

	return string(html), nil
}
