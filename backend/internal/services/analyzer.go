package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/models"
	"github.com/pubuduudara/Golang-assignment-home24-BXP/backend/internal/utils/logger"
	"golang.org/x/net/html"
)

// RequestError is used to handle errors from HTTP requests
type RequestError struct {
	StatusCode  int
	Description string
}

// limits the number of concurrent HTTP link accessibility checks.
// this prevents overwhelming destination web server
const maxConcurrentChecks = 50

// service function for analyzing URLs
func AnalyzeURL(targetURL string) (*models.PageAnalysis, error) {
	logger.Info("Analyzing URL:" + targetURL)
	resp, err := http.Get(targetURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, &RequestError{
			StatusCode:  resp.StatusCode,
			Description: http.StatusText(resp.StatusCode),
		}
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Extract structured data from DOM
	title, headings, links, hasLoginForm := traverseDOM(doc)

	// Detect base URL
	parsedBaseURL, _ := url.Parse(targetURL)

	// Classify links concurrently
	linkCounts := classifyAndCheckLinks(parsedBaseURL, links)
	// get HTML version
	htmlVersion := detectHTMLVersion(doc)

	logger.Info("Analysis complete for URL:" + targetURL)

	return &models.PageAnalysis{
		HTMLVersion:  htmlVersion,
		Title:        title,
		Headings:     headings,
		Links:        linkCounts,
		HasLoginForm: hasLoginForm,
	}, nil
}

// detectHTMLVersion used to get HTML version from the document
func detectHTMLVersion(doc *html.Node) string {
	var doctype string

	var findDoctype func(*html.Node)
	findDoctype = func(n *html.Node) {
		if n.Type == html.DoctypeNode {
			doctype = strings.ToLower(n.Data)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findDoctype(c)
		}
	}
	findDoctype(doc)

	switch {
	case doctype == "html":
		return "HTML5"
	case strings.Contains(doctype, "xhtml"):
		return "XHTML"
	case strings.Contains(doctype, "4.01"):
		return "HTML 4.01"
	default:
		return "Unknown"
	}
}

// walks the HTML tree and collects data
func traverseDOM(root *html.Node) (title string, headings map[string]int, links []string, hasLoginForm bool) {
	headings = make(map[string]int)
	links = []string{}
	titleFound := false

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			// Title
			if !titleFound && n.Data == "title" && n.FirstChild != nil {
				title = n.FirstChild.Data
				titleFound = true
			}

			// Headings (h1-h6)
			if strings.HasPrefix(n.Data, "h") && len(n.Data) == 2 {
				if n.Data[1] >= '1' && n.Data[1] <= '6' {
					headings[n.Data]++
				}
			}

			// Anchor links
			if n.Data == "a" {
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
						break
					}
				}
			}

			// Detect login form
			if n.Data == "input" {
				for _, attr := range n.Attr {
					if attr.Key == "type" && strings.ToLower(attr.Val) == "password" {
						hasLoginForm = true
					}
				}
			}
		}

		// Recursively check children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)

	return
}

// classifies links as internal or external based on base URL
// it also checks link accessibility concurrently using a semaphore to limit parallelism.
// returns a LinkCounts struct with internal, external, and inaccessible link counts.
func classifyAndCheckLinks(base *url.URL, links []string) models.LinkCounts {
	var counts models.LinkCounts
	var mu sync.Mutex
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrentChecks)

	for _, rawLink := range links {
		resolvedURL, err := base.Parse(rawLink)
		if err != nil || resolvedURL.Scheme == "" || resolvedURL.Host == "" {
			continue
		}

		// Classify internal/external
		isInternal := (resolvedURL.Host == base.Host)

		wg.Add(1)
		go func(link string, internal bool) {
			defer wg.Done()

			// acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			accessible := isLinkAccessible(link)

			mu.Lock()
			if internal {
				counts.Internal++
			} else {
				counts.External++
			}
			if !accessible {
				counts.Inaccessible++
			}
			mu.Unlock()
		}(resolvedURL.String(), isInternal)
	}

	wg.Wait()
	return counts
}

// performs a HEAD request to the given link to check if it is reachable.
// Returns true if the response is successful (< 400), false otherwise.
func isLinkAccessible(link string) bool {
	// keep the timeout short to avoid long waits for unresponsive links
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil || resp.StatusCode >= 400 {
		return false
	}
	return true
}

// Error implements the error interface for RequestError.
// Returns the standard HTTP status text for the given status code.
func (e *RequestError) Error() string {
	return fmt.Sprintf("%d %s", e.StatusCode, e.Description)
}
