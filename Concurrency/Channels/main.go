package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// downloadData downloads data from a URL
func downloadData(url string) []byte {
	// perform a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error downloading data:", err)
		return nil
	}
	defer resp.Body.Close()

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading data:", err)
		return nil
	}

	return data
}

// downloadDataConcurrently downloads data from multiple URLs concurrently using goroutines

func downloadDataConcurrently(urls []string) [][]byte {
	// channel to receive downloaded data
	ch := make(chan []byte, len(urls))
	// slice to store the downloaded data
	results := make([][]byte, len(urls))

	// start a goroutine for each URL
	for _, url := range urls {
		go func(url string) {
			// get the data from the URL and send it to the channel
			data := downloadData(url)
			ch <- data
		}(url) // pass the url to the goroutine.
	}

	// wait for all the goroutines to complete
	for i := range urls {
		results[i] = <-ch
	}

	return results
}

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://github.com",
		"https://youtube.com",
		"https://wikipedia.org",
		"https://twitter.com",
		"https://google.com",
		"https://golang.org",
		"https://github.com",
	}
	// get the data from the URLs concurrently
	start := time.Now()
	downloadDataConcurrently(urls)
	fmt.Println("with concurrency:", time.Since(start))

	// get the data from the URLs sequentially
	start = time.Now()
	for _, url := range urls {
		downloadData(url)
	}
	fmt.Println("without concurrency:", time.Since(start))
}
