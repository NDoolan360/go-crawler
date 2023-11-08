package main

import (
	"testing"
)

// TestScrapeAndExtractLinks calls main.scrape with a url, then passes
// that to main.extractLinks checking for a valid return value.
func TestScrapeAndExtractLinks(t *testing.T) {
	content, scrapeErr := scrape("http://quotes.toscrape.com/author/Albert-Einstein/")
	if scrapeErr != nil {
		t.Fatalf(`scrape("http://quotes.toscrape.com/author/Albert-Einstein/") = %q, %v`, content, scrapeErr)
	}

	links, parseErr := extractLinks(content)

	want := []string{"https://www.goodreads.com/quotes", "https://www.zyte.com"}

	if parseErr != nil || len(links) != len(want) || links[0] != want[0] {
		t.Fatalf(`extractLinks(scrape("http://www.example.com")) = %q, %v, want match for %#q, nil`,
			links, parseErr, want)
	}
}
