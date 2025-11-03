package gtvapi

import (
	"context"
	"fmt"
)

// AllTags retrieves all available video tags from the Gronkh.TV API.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns a slice of tags or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	tags, err := client.AllTags(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, tag := range tags {
//	    fmt.Printf("Tag #%d: %s\n", tag.ID, tag.Title)
//	}
func (c *Client) AllTags(ctx context.Context) ([]Tag, error) {
	data, err := c.get(ctx, string(tagsAllEndpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}

	var tags []Tag
	if err := unmarshalResponse(data, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
