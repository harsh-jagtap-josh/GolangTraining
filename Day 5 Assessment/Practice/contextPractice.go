// Contexts provide a mechanism to control the lifecycle, cancellation and propogation of requests across multiple go routines
// --> These basically help in achieving concurrency (for building efficient and robust applications).

// example copied from Medium for reference
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
