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

// TestShallowCrawl calls main.crawl with a url and a depth of 1, checking
// for a valid mutation of VisitedUrls.
func TestShallowCrawl(t *testing.T) {
	visitedUrls := startCrawl("http://quotes.toscrape.com/author/Albert-Einstein/", 1)

	want := make(map[string]bool)
	want["http://quotes.toscrape.com/author/Albert-Einstein/"] = true
	want["https://www.goodreads.com/quotes"] = false
	want["https://www.zyte.com"] = false

	fail := false
	if len(visitedUrls) != len(want) {
		fail = true
	}
	for key, visitValue := range visitedUrls {
		wantValue, ok := want[key]
		if !ok || visitValue != wantValue {
			fail = true
			break
		}
	}
	if fail {
		t.Fatalf(`crawl("http://quotes.toscrape.com/author/Albert-Einstein/", 1, visitedUrls),
        did not result in visitedUrls = %#v, instead it was %#v`, want, visitedUrls)
	}
}

// TestCrawl calls main.crawl with a url and a depth of 3, checking
// for a valid mutation of VisitedUrls.
func TestCrawl(t *testing.T) {
	visitedUrls := startCrawl("http://www.example.com", 3)

	wantedLen := 105

	if len(visitedUrls) != wantedLen {
		t.Fatalf(`crawl("http://www.example.com", 3, visitedUrls), did not result in
        len(visitedUrls) = %#v, instead it was %#v`, wantedLen, len(visitedUrls))
	}
}
