package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("This program expects at least 1 string argument.")
	}
	
	content, err := scrape(os.Args[1])
	if err != nil {
		fmt.Println("Error scraping:", err)
		return
	}
	fmt.Println(extractLinks(content))
}

func scrape(url string) ([]byte, error) {
	res, getErr := http.Get(url)
	if getErr != nil {
		return nil, getErr
	}
	content, readErr := io.ReadAll(res.Body)
	res.Body.Close()
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
