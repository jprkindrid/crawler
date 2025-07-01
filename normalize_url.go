package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	combinedPath := parsedURL.Host + parsedURL.Path
	combinedPath = strings.ToLower(combinedPath)
	combinedPath = strings.TrimSuffix(combinedPath, "/")
	return combinedPath, nil

}
