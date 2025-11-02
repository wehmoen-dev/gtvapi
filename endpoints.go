package gtvapi

// endpoint represents an API endpoint path.
type endpoint string

const (
	// videoInfoEndpoint is the endpoint for retrieving video information.
	videoInfoEndpoint endpoint = "video/info"

	// videoPlaylistEndpoint is the endpoint for retrieving video playlist URLs.
	videoPlaylistEndpoint endpoint = "video/playlist"

	// videoCommentsEndpoint is the endpoint for retrieving video comments.
	videoCommentsEndpoint endpoint = "video/comments"

	// videoDiscoveryEndpoint is the base endpoint for video discovery.
	videoDiscoveryEndpoint endpoint = "video/discovery"

	// externalTwitchLivecheckEndpoint is the endpoint for checking Twitch live status.
	externalTwitchLivecheckEndpoint endpoint = "external/twitch/livecheck"

	// tagsAllEndpoint is the endpoint for retrieving all tags.
	tagsAllEndpoint endpoint = "tags/all"

	// searchEndpoint is the endpoint for searching content.
	searchEndpoint endpoint = "search"
)
