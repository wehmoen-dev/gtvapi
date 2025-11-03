package gtvapi

import (
	"context"
	"fmt"
)

// LiveCheck checks the live status of Twitch channels associated with Gronkh.TV.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns a map of channel names to their live status information,
// or an error if the request fails.
//
// Example:
//
//	ctx := context.Background()
//	channels, err := client.LiveCheck(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for name, info := range channels {
//	    if info.IsLive {
//	        fmt.Printf("%s is live with %d viewers\n", name, info.ViewerCount)
//	    }
//	}
func (c *Client) LiveCheck(ctx context.Context) (map[ChannelName]*ChannelInfo, error) {
	data, err := c.get(ctx, string(externalTwitchLivecheckEndpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check live status: %w", err)
	}

	var response LiveCheck
	if err := unmarshalResponse(data, &response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}
