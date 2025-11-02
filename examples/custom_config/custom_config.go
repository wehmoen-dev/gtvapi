// Package main demonstrates advanced configuration options for the gtvapi library.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wehmoen-dev/gtvapi"
)

func main() {
	// Example 1: Client with debug logging
	fmt.Println("=== Example 1: Debug Logging ===")
	debugClient := gtvapi.NewClient(&gtvapi.ClientConfig{
		Debug:   true,
		Timeout: 10 * time.Second,
	})

	ctx := context.Background()
	_, err := debugClient.AllTags(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Example 2: Client with custom timeout
	fmt.Println("\n=== Example 2: Custom Timeout ===")
	fastClient := gtvapi.NewClient(&gtvapi.ClientConfig{
		Timeout: 5 * time.Second,
	})

	ctx2, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = fastClient.VideoInfo(ctx2, 1000)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Example 3: Client with retry configuration
	fmt.Println("\n=== Example 3: Retry Configuration ===")
	resilientClient := gtvapi.NewClient(&gtvapi.ClientConfig{
		MaxRetries:   5,
		RetryWaitMin: 2 * time.Second,
		RetryWaitMax: 60 * time.Second,
		Timeout:      30 * time.Second,
	})

	_, err = resilientClient.Discover(ctx, gtvapi.DiscoveryTypeRecent)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Example 4: Client with custom HTTP client
	fmt.Println("\n=== Example 4: Custom HTTP Client ===")
	customHTTPClient := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: 2,
		},
	}

	customClient := gtvapi.NewClient(&gtvapi.ClientConfig{
		HTTPClient: customHTTPClient,
	})

	channels, err := customClient.LiveCheck(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Checked %d channels\n", len(channels))
	}

	// Example 5: Client with custom user agent
	fmt.Println("\n=== Example 5: Custom User Agent ===")
	customUAClient := gtvapi.NewClient(&gtvapi.ClientConfig{
		UserAgent: "MyApp/1.0 (+https://example.com)",
	})

	tags, err := customUAClient.AllTags(ctx)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d tags\n", len(tags))
	}

	fmt.Println("\n=== All configuration examples completed ===")
}
