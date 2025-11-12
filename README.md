# Go Challenge: Concurrent URL Processor

This project is a **hands-on Go (Golang) challenge** focused on using **goroutines**, **wait groups**, and **channels**.  
The goal is to practice handling **concurrency** efficiently by simulating multiple HTTP requests being processed at the same time.

---

## Objective

Create a program that takes a list of URLs and performs an HTTP `GET` request to each one **concurrently**.  
Each request should measure the response time and send the result through a **channel**.

---

## Features

- Performs concurrent HTTP requests using `goroutines`
- Synchronizes execution with `sync.WaitGroup`
- Communicates results through `chan Result`
- Displays how long each site took to respond
- Handles request errors gracefully (invalid URLs, timeouts, etc.)

---

## Example Output

```bash
$ go run main.go
https://golang.org -> 212.1948ms
https://google.com -> 113.9301ms
https://youtube.com -> 356.7106ms
https://facebook.com -> 302.4821ms
https://instagram.com -> 417.3342ms
