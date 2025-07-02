package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		fmt.Println("Usage: crawler <URL>")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		fmt.Println("Usage: crawler <URL>")
		os.Exit(1)
	}

	inputURL := args[0]
	fmt.Printf("starting crawl of: %s\n", inputURL)
	rawHTML, err := getHTML(inputURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(rawHTML)

}
