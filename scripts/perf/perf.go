package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

const (
	outerLoop = 10
	// innerLoop = 100_000
	innerLoop = 1000
)

var (
	queryParameters = []string{
		"",
		"?shirt_name=Campbell",
		"?shirt_name=Campbell&country=Netherlands",
		"?shirt_name=Campbell&country=Netherlands&club_id=3e0f6c25-c31d-4bc8-9cbf-b99cd7f08281",
		"?country=Indonesia",
		"?name=Hello world",
		"?name=Hello world&shirt_name=world&country=Netherlands",
		"?name=Hello world&shirt_name=world&country=Netherlands&club_id=3e0f6c25-c31d-4bc8-9cbf-b99cd7f08281",
	}
)

func main() {
	maxResponseTime := -1

	// Fetch.
	fmt.Println("Fetching...")
	intChan := make(chan int)
	var readWg sync.WaitGroup

	for it := 0; it < outerLoop; it++ {
		readWg.Add(1)
		go readResponseTime(intChan, &readWg, &maxResponseTime)
	}

	var fetchWg sync.WaitGroup

	for it := 0; it < outerLoop; it++ {
		fetchWg.Add(innerLoop)
		go fetch(it, intChan, &fetchWg)
	}

	fetchWg.Wait()

	close(intChan)
	readWg.Wait()
	fmt.Println("Done! With max response time", maxResponseTime)
}

func fetch(it int, ch chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Iteration", it)

	queryParamsLen := len(queryParameters)

	for i := 0; i < innerLoop; i++ {
		queryParams := queryParameters[rand.Intn(queryParamsLen)]

		resp, _ := http.Get(fmt.Sprintf("http://localhost:3000%s", queryParams))
		responseTimeStr := resp.Header.Get("response-time")
		responseTime, _ := strconv.Atoi(responseTimeStr)

		wg.Done()
		ch <- responseTime
	}
}

func readResponseTime(ch <-chan int, wg *sync.WaitGroup, maxResponseTime *int) {
	defer wg.Done()

	for responseTime := range ch {
		if responseTime > *maxResponseTime {
			*maxResponseTime = responseTime
		}
	}
}