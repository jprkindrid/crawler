package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "nil", fmt.Errorf("error getting url %v: %w", rawURL, err)
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error getting url:  %v, status code: %v %v", rawURL, res.StatusCode, res.Status)
	}

	contentType := res.Header.Get("Content-Type")

	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("incorrect content type: %v", contentType)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil

}
