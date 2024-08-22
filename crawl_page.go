package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.lenPages() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - normalizedURL: %v", err)
	}

	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	fmt.Printf("[*] crawling %s\n", rawCurrentURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("[-] Error - getHTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(body, cfg.baseURL)
	if err != nil {
		fmt.Printf("[-] Error - getURLsFromHTML: %v", err)
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
