// Package main demonstrates basic usage of the gtvapi library.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wehmoen-dev/gtvapi"
)

func main() {
	// Create a new client with default configuration
	client := gtvapi.NewClient(nil)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Example 1: Get video information
	fmt.Println("=== Example 1: Video Information ===")
	videoInfo, err := client.VideoInfo(ctx, 1000)
	if err != nil {
		log.Printf("Error getting video info: %v\n", err)
	} else {
		fmt.Printf("Video ID: %d\n", videoInfo.ID)
		fmt.Printf("Title: %s\n", videoInfo.Title)
		fmt.Printf("Season %d, Episode %d\n", videoInfo.Season, videoInfo.Episode)
		fmt.Printf("Views: %d\n", videoInfo.Views)
		fmt.Printf("Duration: %d seconds\n", videoInfo.SourceLength)
		fmt.Printf("Tags: %d\n", len(videoInfo.Tags))
	}

	// Example 2: Get all tags
	fmt.Println("\n=== Example 2: All Tags ===")
	tags, err := client.AllTags(ctx)
	if err != nil {
		log.Printf("Error getting tags: %v\n", err)
	} else {
		fmt.Printf("Total tags: %d\n", len(tags))
		fmt.Println("First 5 tags:")
		for i, tag := range tags {
			if i >= 5 {
				break
			}
			fmt.Printf("  - %s (ID: %d)\n", tag.Title, tag.ID)
		}
	}

	// Example 3: Discover recent videos
	fmt.Println("\n=== Example 3: Recent Videos ===")
	recentVideos, err := client.Discover(ctx, gtvapi.DiscoveryTypeRecent)
	if err != nil {
		log.Printf("Error discovering videos: %v\n", err)
	} else {
		fmt.Printf("Found %d recent videos\n", len(recentVideos))
		if len(recentVideos) > 0 {
			fmt.Println("First video:")
			fmt.Printf("  Title: %s\n", recentVideos[0].Title)
			fmt.Printf("  Episode: %d\n", recentVideos[0].Episode)
			fmt.Printf("  Views: %d\n", recentVideos[0].Views)
		}
	}

	// Example 4: Check Twitch live status
	fmt.Println("\n=== Example 4: Twitch Live Status ===")
	channels, err := client.LiveCheck(ctx)
	if err != nil {
		log.Printf("Error checking live status: %v\n", err)
	} else {
		fmt.Printf("Monitoring %d channels\n", len(channels))
		for name, info := range channels {
			if info.IsLive {
				fmt.Printf("  🔴 %s is LIVE!\n", name)
				fmt.Printf("     Game: %s\n", info.GameName)
				fmt.Printf("     Viewers: %d\n", info.ViewerCount)
				fmt.Printf("     Title: %s\n", info.Title)
			} else {
				fmt.Printf("  ⚫ %s is offline\n", name)
			}
		}
	}

	// Example 5: Get video comments
	fmt.Println("\n=== Example 5: Video Comments ===")
	comments, err := client.VideoComments(ctx, 1000)
	if err != nil {
		log.Printf("Error getting comments: %v\n", err)
	} else {
		fmt.Printf("Total comments: %d\n", len(comments))
		if len(comments) > 0 {
			fmt.Println("First comment:")
			fmt.Printf("  User: %s\n", comments[0].User.Username)
			fmt.Printf("  Comment: %s\n", comments[0].Comment)
			fmt.Printf("  Upvotes: %d\n", comments[0].UpvoteCount)
		}
	}

	// Example 6: Get video playlist URL
	fmt.Println("\n=== Example 6: Video Playlist ===")
	playlistURL, err := client.VideoPlaylist(ctx, 1000)
	if err != nil {
		log.Printf("Error getting playlist: %v\n", err)
	} else {
		fmt.Printf("Playlist URL: %s\n", playlistURL)
	}

	fmt.Println("\n=== All examples completed ===")
}
