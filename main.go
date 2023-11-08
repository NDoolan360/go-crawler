package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("This program expects at least 1 string argument.")
	}
	fmt.Println(scrape(os.Args[1]))
}

func scrape(url string) string {
	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	content, readErr := io.ReadAll(res.Body)
	res.Body.Close()
	if readErr != nil {
		log.Fatal(readErr)
	}
	return string(content)
}
