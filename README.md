# gtvapi - Gronkh.TV API Client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/wehmoen-dev/gtvapi.svg)](https://pkg.go.dev/github.com/wehmoen-dev/gtvapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/wehmoen-dev/gtvapi)](https://goreportcard.com/report/github.com/wehmoen-dev/gtvapi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, production-ready Go client library for the [Gronkh.TV](https://gronkh.tv) API. This library provides easy access to video information, comments, playlists, discovery features, and Twitch live status.

## Features

- 🚀 **Simple and intuitive API** - Easy-to-use client with clear method names
- 🔄 **Automatic retry logic** - Built-in retry with exponential backoff for resilient API calls
- ⏱️ **Context support** - Full context support for cancellation and timeouts
- 📝 **Comprehensive documentation** - Complete GoDoc comments for all public APIs
- ✅ **Well-tested** - High test coverage with unit and integration tests
- 🛡️ **Type-safe** - Strongly typed responses with proper error handling
- 🔍 **Request logging** - Optional debug logging for troubleshooting

## Installation

```bash
go get github.com/wehmoen-dev/gtvapi
```

## Requirements

- Go 1.20 or higher

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/wehmoen-dev/gtvapi"
)

func main() {
    // Create a new client
    client := gtvapi.NewClient(nil)

    // Set a timeout context
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get video information
    videoInfo, err := client.VideoInfo(ctx, 1000)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Video: %s (Episode %d)\n", videoInfo.Title, videoInfo.Episode)
    fmt.Printf("Views: %d\n", videoInfo.Views)
}
```

## Usage Examples

### Creating a Client

#### Basic Client
```go
client := gtvapi.NewClient(nil)
```

#### Client with Custom Options
```go
config := &gtvapi.ClientConfig{
    Debug:          true,                    // Enable debug logging
    Timeout:        30 * time.Second,        // Request timeout
    MaxRetries:     3,                       // Maximum retry attempts
    RetryWaitMin:   1 * time.Second,         // Minimum wait between retries
    RetryWaitMax:   30 * time.Second,        // Maximum wait between retries
}

client := gtvapi.NewClient(config)
```

### Getting Video Information

```go
ctx := context.Background()
videoInfo, err := client.VideoInfo(ctx, 1000)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Title: %s\n", videoInfo.Title)
fmt.Printf("Season: %d, Episode: %d\n", videoInfo.Season, videoInfo.Episode)
fmt.Printf("Duration: %d seconds\n", videoInfo.SourceLength)
```

### Fetching Video Comments

```go
ctx := context.Background()
comments, err := client.VideoComments(ctx, 1000)
if err != nil {
    log.Fatal(err)
}

for _, comment := range comments {
    fmt.Printf("%s: %s\n", comment.User.Username, comment.Comment)
}
```

### Getting Video Playlist URL

```go
ctx := context.Background()
playlistURL, err := client.VideoPlaylist(ctx, 1000)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Playlist URL: %s\n", playlistURL)
```

### Discovering Videos

```go
ctx := context.Background()

// Get most recent videos
recent, err := client.Discover(ctx, gtvapi.DiscoveryTypeRecent)
if err != nil {
    log.Fatal(err)
}

// Get most viewed videos
popular, err := client.Discover(ctx, gtvapi.DiscoveryTypeViews)
if err != nil {
    log.Fatal(err)
}

// Get similar videos
similar, err := client.Discover(ctx, gtvapi.DiscoveryTypeSimilar)
if err != nil {
    log.Fatal(err)
}
```

### Fetching All Tags

```go
ctx := context.Background()
tags, err := client.AllTags(ctx)
if err != nil {
    log.Fatal(err)
}

for _, tag := range tags {
    fmt.Printf("Tag #%d: %s\n", tag.ID, tag.Title)
}
```

### Checking Twitch Live Status

```go
ctx := context.Background()
channels, err := client.LiveCheck(ctx)
if err != nil {
    log.Fatal(err)
}

for name, info := range channels {
    if info.IsLive {
        fmt.Printf("%s is live! Viewers: %d\n", name, info.ViewerCount)
        fmt.Printf("Playing: %s\n", info.GameName)
    }
}
```

## API Reference

### Client Methods

#### `NewClient(config *ClientConfig) *Client`
Creates a new Gronkh.TV API client with optional configuration.

#### `VideoInfo(ctx context.Context, episode int) (*VideoInfo, error)`
Retrieves detailed information about a specific video episode.

#### `VideoComments(ctx context.Context, episode int) ([]Comment, error)`
Fetches all comments for a specific video episode.

#### `VideoPlaylist(ctx context.Context, episode int) (string, error)`
Gets the playlist URL for streaming a specific video episode.

#### `Discover(ctx context.Context, discoveryType DiscoveryType) ([]VideoSearchResult, error)`
Discovers videos based on the specified discovery type (recent, views, or similar).

#### `AllTags(ctx context.Context) ([]Tag, error)`
Retrieves all available video tags.

#### `LiveCheck(ctx context.Context) (map[ChannelName]ChannelInfo, error)`
Checks the live status of Twitch channels associated with Gronkh.TV.

### Types

See the [GoDoc](https://pkg.go.dev/github.com/wehmoen-dev/gtvapi) for detailed type definitions and field descriptions.

## Error Handling

The library provides structured error handling:

```go
videoInfo, err := client.VideoInfo(ctx, 1000)
if err != nil {
    // Check for specific error types
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Request timed out")
    } else if errors.Is(err, context.Canceled) {
        log.Println("Request was canceled")
    } else {
        log.Printf("API error: %v", err)
    }
    return
}
```

## Best Practices

1. **Always use context** - Pass context to all API methods for proper timeout and cancellation handling
2. **Set reasonable timeouts** - Configure appropriate timeouts based on your use case
3. **Handle errors properly** - Check and handle all errors returned by the API
4. **Reuse clients** - Create one client and reuse it for multiple requests
5. **Enable retry logic** - Use the built-in retry mechanism for production environments

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

Run the test suite:

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...

# Run tests with race detection
go test -v -race ./...

# Run benchmarks
go test -bench=. ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gronkh.TV](https://gronkh.tv) for providing the API
- [Resty](https://resty.dev) for the excellent HTTP client library

## Support

For questions, issues, or feature requests, please open an issue on [GitHub](https://github.com/wehmoen-dev/gtvapi/issues).
