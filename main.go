package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
)

func main() {
	urls := []string{"https://google.com", "https://www.cartier.com"}
	var client http.Client
	var totalResponseTime float64 = 0
	fmt.Println("Response time:")
	for _, url := range urls {
		var start = time.Now()
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			totalResponseTime += time.Since(start).Seconds()
			fmt.Println(url, ":", time.Since(start))
		}
	}

	fmt.Println("-----------------------")
	fmt.Println("Average response time :", fmt.Sprintf("%.2f", totalResponseTime/float64(len(urls))), "seconds")
}