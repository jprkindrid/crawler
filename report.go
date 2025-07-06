package main

import (
	"fmt"
	"sort"
)

type pageCountSorted struct {
	page  string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.page)
	}

	fmt.Println("...")

}

func sortPages(pages map[string]int) []pageCountSorted {

	var sortingStruct []pageCountSorted
	for page, count := range pages {
		sortingStruct = append(sortingStruct, pageCountSorted{page, count})
	}

	sort.Slice(sortingStruct, func(i, j int) bool {
		return sortingStruct[i].count > sortingStruct[j].count
	})

	sortedPages := make(map[string]int)
	for _, page := range sortingStruct {
		sortedPages[page.page] = page.count
	}

	return sortingStruct
}
