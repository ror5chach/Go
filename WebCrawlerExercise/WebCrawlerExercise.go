package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetcherResult struct {
	body string
	urls []string
	err error
}

func CrawlHelper(url string, depth int, fetcher Fetcher, ch chan *FetcherResult) {
	fmt.Println("start crawl")
	
	if (depth <= 0) {
		close(ch)
		return
	}
	
	var fr FetcherResult
	fr.body, fr.urls, fr.err = fetcher.Fetch(url)
	
	ch <-&fr
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	
	ch := make(chan *FetcherResult)
	
	mapFetcherResult := make(map[string]bool)
	mapFetcherResult[url] = true
	
	go CrawlHelper(url, depth, fetcher, ch)
	
	for res := range ch {
		for _, urlVal := range res.urls {
			_, found := mapFetcherResult[urlVal]
			if !found {
				mapFetcherResult[urlVal] = true
				
				fmt.Println("URL: ", urlVal, " ,Body: ", res.body)
				depth--
				go CrawlHelper(urlVal, depth, fetcher, ch)
			}
		}
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
