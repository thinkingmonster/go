package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usages: URL TotalRequestPerSecond Duration\n Example:\n" +
			" pingming https://a.com 15 20 \n" +
			"This will trigger 15 requests per second on https://a.com for 20 seconds")
		os.Exit(1)
	}
	var Url string
	var TotalRequestPerSecond, Duration int
	var SampleData []string
	start := time.Now()
	defer func() {
		ExecutionTime := int(time.Since(start).Seconds())
		TotalRequestsSent := TotalRequestPerSecond * Duration
		fmt.Printf("Execution Time: %d\n", ExecutionTime)
		fmt.Printf("Total Requests sent: %d\n", TotalRequestsSent)
		fmt.Printf("Request Rate: %d\n", TotalRequestsSent/ExecutionTime)
	}()

	Url = os.Args[1]
	TotalRequestPerSecond, _ = strconv.Atoi(os.Args[2])
	Duration, _ = strconv.Atoi(os.Args[3])

	// Building Sample data
	for i := 1; i <= TotalRequestPerSecond; i++ {
		SampleData = append(SampleData, Url)
	}

	for i := 1; i <= Duration; i++ {
		checkUrls(SampleData)
	}

}

func checkUrls(urls []string) {
	c := make(chan string)
	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1)
		go checkUrl(link, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}
func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	output, _ := http.Get(url)
	if output.Status != "200 OK" {
		c <- "ERROR:   " + output.Status
	} else {
		c <- "SUCCESS: " + output.Status + " Running on " + output.Header["X-Infa-Server-Agent"][0]
	}
}
