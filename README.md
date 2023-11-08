# Go Crawler

**go-crawler** is a simple web crawling program written in Go in an afternoon. It allows you to start from a given URL and crawl through web pages, collecting links up to a specified depth and printing what is has found to stdout. This tool could be a great foundation for various web scraping and web data collection applications. 

## Getting Started

To get started with the Go Crawler, follow these simple steps:

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-crawler.git
cd go-crawler
```

2. Build the executable:

```bash
go get go-crawler
go build
```

3. Run the program:

```bash
./go-crawler <startURL> <maxDepth>
```

- \<startURL>: The URL from which the crawling will begin.
- \<maxDepth> (optional): The maximum depth for crawling. Default is 1 if not specified.

## Usage

1. The program takes at least one command-line argument, which is the starting URL for crawling. You can optionally provide a second argument for the maximum depth of the crawl.

2. The crawler will start from the specified URL and collect links up to the specified depth.

3. The crawled URLs and any errors encountered during the process will be printed to the console.

### Example

```bash
./go-crawler https://example.com 2
```

This command will start crawling from "https://example.com" up to a depth of 2.

## Features

- Recursive web crawling starting from a given URL.
- Specify the maximum depth for the crawl.
- Handle and report HTTP errors.
- Extract links from HTML content.

## Dependencies

The program uses the following Go packages:

- net/http: For making HTTP requests.
- golang.org/x/net/html: For parsing HTML content and extracting links.
