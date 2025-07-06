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

	var pagesCountSorted []pageCountSorted
	for page, count := range pages {
		pagesCountSorted = append(pagesCountSorted, pageCountSorted{page, count})
	}

	sort.Slice(pagesCountSorted, func(i, j int) bool {
		if pagesCountSorted[i].count == pagesCountSorted[j].count {
			return pagesCountSorted[i].page < pagesCountSorted[j].page
		}
		return pagesCountSorted[i].count > pagesCountSorted[j].count
	})

	return pagesCountSorted
}
