package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawbaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawbaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldnt parse raw base url %w", err)
	}
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html reader: %w", err)
	}

	var urls []string
	var traverseNodes func(node *html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					hrefURL, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
						continue
					}

					fullURL := baseURL.ResolveReference(hrefURL)
					urls = append(urls, fullURL.String())
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil
}
