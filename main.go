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

	rootNode := startCrawl(startUrl, maxDepth)
	fmt.Println(rootNode.toString())
}

func startCrawl(startUrl string, maxDepth int) Link {
	initial := Link{Url: startUrl, Depth: 0, Children: []*Link{}}
	return crawl(initial, maxDepth)
}

func crawl(link Link, maxDepth int) Link {
	if link.Depth >= maxDepth {
		return link
	}

	content, err := scrape(link.Url)
	if err != nil {
		fmt.Println("Error scraping", link.Url, ":", err)
		return link
	}

	urls, err := extractUrls(content)
	if err != nil {
		fmt.Println("Error extracting links from", link.Url, ":", err)
	}

	for _, url := range urls {
		newChild := Link{Url: url, Depth: link.Depth + 1, Children: []*Link{}}
		link.Children = append(link.Children, &newChild)
		newChild = crawl(newChild, maxDepth)
	}
	return link
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

func extractUrls(content []byte) ([]string, error) {
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
