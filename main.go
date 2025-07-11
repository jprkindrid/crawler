package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages       map[string]int
	baseURL     *url.URL
	mu          *sync.Mutex
	concControl chan struct{}
	wg          *sync.WaitGroup
	maxPages    int
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		fmt.Println("Usage: crawler <URL>")
		os.Exit(1)
	}

	inputURL := args[0]

	// Set defaults
	maxConcurrency := 3
	maxPages := 10

	if len(args) > 1 && args[1] != "" {
		if i, err := strconv.Atoi(args[1]); err == nil {
			maxConcurrency = i
		} else {
			fmt.Println("Error parsing maxConcurrency argument, using default:", maxConcurrency)
		}
	}

	if len(args) > 2 && args[2] != "" {
		if i, err := strconv.Atoi(args[2]); err == nil {
			maxPages = i
		} else {
			fmt.Println("Error parsing maxPages argument, using default:", maxPages)
		}
	}

	cfg, err := configure(inputURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
	}

	fmt.Printf("starting crawl of: %s\n", inputURL)
	fmt.Printf("Maximum Concurrency: %d, Maximum Pages: %d\n", maxConcurrency, cfg.maxPages)

	cfg.wg.Add(1)
	go cfg.crawlPage(inputURL)
	cfg.wg.Wait()

	printReport(cfg.pages, inputURL)

}

func configure(rawbaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawbaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldnt parse raw base URL: %w", err)
	}

	return &config{
		pages:       make(map[string]int),
		baseURL:     baseURL,
		mu:          &sync.Mutex{},
		concControl: make(chan struct{}, maxConcurrency),
		wg:          &sync.WaitGroup{},
		maxPages:    maxPages,
	}, nil
}
