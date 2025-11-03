package gtvapi

import (
	"context"
	"fmt"
)

// VideoInfo retrieves detailed information about a specific video episode.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - episode: The episode number to retrieve information for
//
// Returns the video information or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	videoInfo, err := client.VideoInfo(ctx, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Title: %s\n", videoInfo.Title)
func (c *Client) VideoInfo(ctx context.Context, episode int) (*VideoInfo, error) {
	if episode <= 0 {
		return nil, fmt.Errorf("episode must be positive, got %d", episode)
	}

	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	data, err := c.get(ctx, string(videoInfoEndpoint), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %w", err)
	}

	var info VideoInfo
	if err := unmarshalResponse(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// VideoComments retrieves all comments for a specific video episode.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - episode: The episode number to retrieve comments for
//
// Returns a slice of comments or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	comments, err := client.VideoComments(ctx, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, comment := range comments {
//	    fmt.Printf("%s: %s\n", comment.User.Username, comment.Comment)
//	}
func (c *Client) VideoComments(ctx context.Context, episode int) ([]Comment, error) {
	if episode <= 0 {
		return nil, fmt.Errorf("episode must be positive, got %d", episode)
	}

	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	data, err := c.get(ctx, string(videoCommentsEndpoint), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get video comments: %w", err)
	}

	var response struct {
		Comments []Comment `json:"comments"`
	}
	if err := unmarshalResponse(data, &response); err != nil {
		return nil, err
	}

	return response.Comments, nil
}

// VideoPlaylist retrieves the playlist URL for streaming a specific video episode.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - episode: The episode number to retrieve the playlist URL for
//
// Returns the playlist URL or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	playlistURL, err := client.VideoPlaylist(ctx, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Playlist URL: %s\n", playlistURL)
func (c *Client) VideoPlaylist(ctx context.Context, episode int) (string, error) {
	if episode <= 0 {
		return "", fmt.Errorf("episode must be positive, got %d", episode)
	}

	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	data, err := c.get(ctx, string(videoPlaylistEndpoint), query)
	if err != nil {
		return "", fmt.Errorf("failed to get video playlist: %w", err)
	}

	var response struct {
		PlaylistURL string `json:"playlist_url"`
	}
	if err := unmarshalResponse(data, &response); err != nil {
		return "", err
	}

	return response.PlaylistURL, nil
}

// Discover retrieves videos based on the specified discovery type.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - discoveryType: The type of discovery (DiscoveryTypeRecent, DiscoveryTypeViews, or DiscoveryTypeSimilar)
//
// Returns a slice of video search results or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	videos, err := client.Discover(ctx, gtvapi.DiscoveryTypeRecent)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, video := range videos {
//	    fmt.Printf("Video: %s (Episode %d)\n", video.Title, video.Episode)
//	}
func (c *Client) Discover(ctx context.Context, discoveryType DiscoveryType) ([]VideoSearchResult, error) {
	var endpoint string

	switch discoveryType {
	case DiscoveryTypeRecent, DiscoveryTypeViews, DiscoveryTypeSimilar:
		endpoint = fmt.Sprintf("%s/%s", videoDiscoveryEndpoint, discoveryType)
	default:
		return nil, fmt.Errorf("invalid discovery type: %s", discoveryType)
	}

	data, err := c.get(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get discovery results: %w", err)
	}

	var response struct {
		Discovery []VideoSearchResult `json:"discovery"`
	}
	if err := unmarshalResponse(data, &response); err != nil {
		return nil, err
	}

	return response.Discovery, nil
}
