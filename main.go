package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("[x] Error - maxConcurrency: %v", err)
		return
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("[x] Error - maxPages: %v", err)
		return
	}

	cfg, err := initialize(rawBaseURL, maxPages, maxConcurrency)
	if err != nil {
		fmt.Printf("[x] Error - configure: %v", err)
		return
	}

	fmt.Println("===================================")
	fmt.Printf("  STARTING CRAWL for: %s...\n", rawBaseURL)
	fmt.Println("===================================")

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}
