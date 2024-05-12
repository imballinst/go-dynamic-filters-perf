package main

import (
	"fmt"
	helper "go-dynamic-filters-perf/pkg"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	outerLoop = 10
)

var (
	innerLoop, _ = strconv.Atoi(os.Getenv("LOOPS"))
)

func main() {
	if innerLoop == 0 {
		panic("please set LOOPS environment variable")
	}

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

	for i := 0; i < innerLoop; i++ {
		id, clubId, name, country, shirtName := helper.GetRandomValues()
		queryParams := []string{
			fmt.Sprintf("id=%s", id),
			fmt.Sprintf("club_id=%s", clubId),
			fmt.Sprintf("name=%s", name),
			fmt.Sprintf("country=%s", country),
			fmt.Sprintf("shirt_name=%s", shirtName),
		}

		queryParamsSliceStart := rand.Intn(len(queryParams))
		queryParamsSliceEnd := rand.Intn(len(queryParams))

		if queryParamsSliceEnd < queryParamsSliceStart {
			tmp := queryParamsSliceEnd
			queryParamsSliceEnd = queryParamsSliceStart
			queryParamsSliceStart = tmp
		}

		queryParamsJoined := strings.Join(queryParams[queryParamsSliceStart:queryParamsSliceEnd], "&")

		resp, _ := http.Get(fmt.Sprintf("http://localhost:3000?%s", queryParamsJoined))
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
