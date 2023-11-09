package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapeAndExtractUrls(t *testing.T) {
	mockServer := createMockServer("<a href='http://example.com'></a><a href='http://example.org'></a>")
	defer mockServer.Close()

	content, err := scrape(mockServer.URL)
	if err != nil {
		t.Fatalf("Error scraping URL %s: %v", mockServer.URL, err)
	}

	links, err := extractUrls(content)
	if err != nil {
		t.Fatalf("Error extracting links from URL %s: %v", mockServer.URL, err)
	}

	expectedLinks := []string{"http://example.com", "http://example.org"}
	if len(links) != len(expectedLinks) {
		t.Fatalf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	for i, expectedLink := range expectedLinks {
		if links[i] != expectedLink {
			t.Errorf("Expected link at index %d to be %s, got %s", i, expectedLink, links[i])
		}
	}
}

func TestCrawl(t *testing.T) {
	mockServer3 := createMockServer("<a href='http://example.com'></a>")
	defer mockServer3.Close()
	mockServer2 := createMockServer("<a href='" + mockServer3.URL + "'></a>")
	defer mockServer2.Close()
	mockServer1 := createMockServer("<a href='" + mockServer2.URL + "'></a>")
	defer mockServer1.Close()

	maxDepth := 3
	result := startCrawl(mockServer1.URL, maxDepth)

	expectedUrls := []string{mockServer1.URL, mockServer2.URL, mockServer3.URL, "http://example.com"}

	for index, expectedUrl := range expectedUrls {
		if result.Url != expectedUrl {
			t.Errorf("Expected parent URL to be %s, got %s", expectedUrl, result.Url)
		}

		if index != len(expectedUrls)-1 {
			break
		}

		if len(result.Children) != 1 {
			t.Errorf("Expected parent URL to have 1 child link, got %d", len(result.Children))
		} else {
			result = *result.Children[0]
		}

	}
}

func createMockServer(responseBody string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseBody))
	}))
}
