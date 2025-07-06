package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concControl <- struct{}{}
	defer func() {
		<-cfg.concControl
		cfg.wg.Done()
	}()

	if len(cfg.pages) >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldnt parse url %s: %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error -  crawlpage: couldnt normalize url %s: %v\n", rawCurrentURL, err)
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("Getting HTML from %s...\n", rawCurrentURL)
	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlpage: couldn't get html from url %s: %v\n", rawCurrentURL, err)
		return
	}

	pageURLs, err := getURLsFromHTML(pageHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - crawlpage: couldn't get urls from html at %s: %v\n", rawCurrentURL, err)
	}

	for _, pageURL := range pageURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(pageURL)
	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1

	return true
}
