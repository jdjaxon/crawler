package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Test Title",
		FirstParagraph: "This is the first paragraph.",
		OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
		ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestExtractPageDataDuplicateTagsComplex(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `
		<html>
			<body>
				<h1>Test Title 1</h1>
				<h1>Test Title 2</h1>
				<p>This is the first paragraph.</p>
				<p>This is the second paragraph.</p>
				<a href="/link1">Link 1</a>
				<a>No Link</a>
				<a href="https://blog.boot.dev/link2">Link 2</a>
				<img src="/image1.jpg" alt="Image 1">
				<img alt="No image">
				<img src="https://blog.boot.dev/image2.jpg" alt="Image 2">
			</body>
		</html>
	`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Test Title 1",
		FirstParagraph: "This is the first paragraph.",
		OutgoingLinks: []string{
			"https://blog.boot.dev/link1",
			"https://blog.boot.dev/link2",
		},
		ImageURLs: []string{
			"https://blog.boot.dev/image1.jpg",
			"https://blog.boot.dev/image2.jpg",
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
