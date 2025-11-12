package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Duration time.Duration
	Err      error
}

func requestGET(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	ch <- Result{
		URL:      url,
		Duration: time.Since(start),
	}
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	wg := sync.WaitGroup{}

	ch := make(chan Result, len(urls))

	for _, u := range urls {
		wg.Add(1)
		go requestGET(u, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		if res.Err != nil {
			fmt.Printf("%s -> error: %v\n", res.URL, res.Err)
			continue
		}
		fmt.Printf("%s -> %v\n", res.URL, res.Duration)
	}

}
