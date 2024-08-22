package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("[-] Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - normalizedURL: %v", err)
	}

	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}
	pages[normalizedURL] = 1

	fmt.Printf("[*] crawling %s\n", rawCurrentURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - getHTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(body, rawBaseURL)
	if err != nil {
		fmt.Printf("[-] Error - getURLsFromHTML: %v", err)
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
