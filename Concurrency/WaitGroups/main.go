package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// download data from a URL
func downloadData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading data:", err)
		return nil
	}
	defer resp.Body.Close()

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return nil
	}

	return data
}

// download data from multiple URLs concurrently using goroutines
func downloadDataConcurrently(urls []string) [][]byte {
	// wait group with size equal to the number of URLs
	var wg sync.WaitGroup
	wg.Add(len(urls))

	// slice to store the downloaded data
	results := make([][]byte, len(urls))

	// start a goroutine for each URL
	for i, url := range urls {
		go func(i int, url string) {
			// get the data from the URL and store it in the results slice
			data := downloadData(url)
			results[i] = data
			// signal that the goroutine has completed by calling Done on the wait group
			wg.Done()
		}(i, url)
	}

	// wait for all the goroutines to complete
	wg.Wait()

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

	//  get the data from the URLs concurrently
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
