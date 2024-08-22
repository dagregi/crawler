package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("===================================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("===================================")

	pagesSlice := make([]Page, 0, len(pages))
	for l, c := range pages {
		pagesSlice = append(pagesSlice, Page{URL: l, Count: c})
	}

	pagesSlice = sortPages(pagesSlice)
	for _, pg := range pagesSlice {
		fmt.Printf("Found %d internal links to https://%s\n", pg.Count, pg.URL)
	}
}

func sortPages(pages []Page) []Page {
	sort.Slice(pages, func(i, j int) bool {
		if pages[i].Count == pages[j].Count {
			return pages[i].URL < pages[j].URL
		}
		return pages[i].Count > pages[j].Count
	})

	return pages
}
