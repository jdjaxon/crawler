package main

import (
	"testing"
)

func TestGetH1FromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1></body></html>"
	actual := getH1FromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetMultiH1FromHTMLBasic(t *testing.T) {
	inputBody := `
		<html>
			<body>
				<h1>Test Title 1</h1>
				<h1>Test Title 2</h1>
			</body>
		</html>
	`
	actual := getH1FromHTML(inputBody)
	expected := "Test Title 1"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTMLNoMatch(t *testing.T) {
	inputBody := "<html><body>Test Title</body></html>"
	actual := getH1FromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTMLNested(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<h1>Inner header</h1>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getH1FromHTML(inputBody)
	expected := "Inner header"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLNoMain(t *testing.T) {
	inputBody := `<html><body>
		<p>First paragraph.</p>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "First paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphNoParagraph(t *testing.T) {
	inputBody := `<html><body>Hello</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphMainPriorityNoParagraph(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main><h1>Header</h1></main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Outside paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
