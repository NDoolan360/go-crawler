package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("This program expects at least 1 string argument.")
	}
	startUrl := os.Args[1]

	maxDepth := 1
	if len(os.Args) >= 3 {
		var err error
		maxDepth, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	startCrawl(startUrl, maxDepth)
}

func startCrawl(startUrl string, maxDepth int) map[string]bool {
	visitedUrls := make(map[string]bool)
	crawl(startUrl, maxDepth, visitedUrls)
	return visitedUrls
}

func crawl(url string, depth int, visitedUrls map[string]bool) {
	if depth <= 0 {
		return
	}
	fmt.Println("Crawling", url)

	content, err := scrape(url)
	if err != nil {
		fmt.Println("Error scraping", url, ":", err)
		return
	}

	visitedUrls[url] = true

	links, err := extractLinks(content)
	if err != nil {
		fmt.Println("Error extracting links from", url, ":", err)
	}

	for _, link := range links {
		if !visitedUrls[link] {
			fmt.Println("Found", link)
			visitedUrls[link] = false
			crawl(link, depth-1, visitedUrls)
		}
	}
}

func scrape(url string) ([]byte, error) {
	res, getErr := http.Get(url)
	if getErr != nil {
		return nil, getErr
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP response error: %s", res.Status)
	}
	content, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return content, nil
}

func extractLinks(content []byte) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(bytes.NewReader(content))

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" && strings.HasPrefix(attr.Val, "http") {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
