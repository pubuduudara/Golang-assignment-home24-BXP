package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
	"golang.org/x/net/html"
)

// test HTML version detection from doctype
func TestDetectHTMLVersion(t *testing.T) {
	cases := []struct {
		doctype string
		expect  string
	}{
		{`<!DOCTYPE html>`, "HTML5"},
		{`<html>`, "Unknown"},
		{`<!DOCTYPE notarealdoctype>`, "Unknown"},
	}

	for _, tc := range cases {
		doc, _ := html.Parse(strings.NewReader(tc.doctype + "<html><head></head><body></body></html>"))
		version := detectHTMLVersion(doc)
		if version != tc.expect {
			logger.Error(fmt.Errorf("got %s, expected %s", version, tc.expect), "HTML version mismatch")
			t.Fail()
		}
	}
}

// test DOM traversal with headings, title, link, login form
func TestTraverseDOM(t *testing.T) {
	htmlStr := `
		<!DOCTYPE html>
		<html>
			<head><title>Test Page</title></head>
			<body>
				<h1>Main</h1>
				<h2>Sub</h2>
				<a href="/internal">Internal</a>
				<form><input type="password" /></form>
			</body>
		</html>`
	doc, _ := html.Parse(strings.NewReader(htmlStr))
	title, headings, links, hasLogin := traverseDOM(doc)

	if title != "Test Page" {
		logger.Error(fmt.Errorf("expected title 'Test Page', got %s", title))
		t.Fail()
	}
	if headings["h1"] != 1 || headings["h2"] != 1 {
		logger.Error(fmt.Errorf("incorrect heading counts: %v", headings))
		t.Fail()
	}
	if len(links) != 1 {
		logger.Error(fmt.Errorf("expected 1 link, got %d", len(links)))
		t.Fail()
	}
	if !hasLogin {
		logger.Error(fmt.Errorf("expected login form to be detected"))
		t.Fail()
	}
}

// test analysis on a fake HTML page served via httptest
func TestAnalyzeURL(t *testing.T) {
	// Mock server with known HTML
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `
			<!DOCTYPE html>
			<html>
				<head><title>Mock Site</title></head>
				<body>
					<h1>Hello</h1>
					<a href="https://external.com">Link</a>
					<form><input type="password" /></form>
				</body>
			</html>`
		w.Write([]byte(html))
	}))
	defer server.Close()

	result, err := AnalyzeURL(server.URL)
	if err != nil {
		logger.Error(err, "analyze URL failed")
		t.FailNow()
	}

	if result.Title != "Mock Site" {
		t.Errorf("Expected title 'Mock Site', got %s", result.Title)
	}
	if result.HTMLVersion != "HTML5" {
		t.Errorf("Expected HTML5, got %s", result.HTMLVersion)
	}
	if result.Headings["h1"] != 1 {
		t.Errorf("Expected 1 h1, got %d", result.Headings["h1"])
	}
	if !result.HasLoginForm {
		t.Errorf("Expected login form detection")
	}
}

// test isLinkAccessible returns false for fake 404 link
func TestIsLinkAccessible(t *testing.T) {
	// Fake server that always returns 404
	badServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer badServer.Close()

	ok := isLinkAccessible(badServer.URL)
	if ok {
		logger.Error(fmt.Errorf("expected link to be inaccessible"))
		t.Fail()
	}
}

// test RequestError Error() method formatting
func TestRequestError_Error(t *testing.T) {
	e := &RequestError{StatusCode: 404, Description: "Not Found"}
	expected := "404 Not Found"
	if e.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, e.Error())
	}
}
