package gtvapi

import "time"

// TwitchDetails contains Twitch-specific information about a game.
type TwitchDetails struct {
	// ID is the Twitch game ID.
	ID string `json:"id"`
	// Title is the game title on Twitch.
	Title string `json:"title"`
	// ThumbnailURL is the URL to the game's thumbnail image.
	ThumbnailURL string `json:"thumbnail_url"`
}

// Tag represents a video tag/category.
type Tag struct {
	// ID is the unique identifier for the tag.
	ID int8 `json:"id"`
	// Title is the display name of the tag.
	Title string `json:"title"`
}

// Game represents a video game.
type Game struct {
	// ID is the unique identifier for the game.
	ID int `json:"id"`
	// Title is the game's name.
	Title string `json:"title"`
	// TwitchDetails contains Twitch-specific information about the game.
	TwitchDetails TwitchDetails `json:"twitch_details"`
}

// Chapter represents a chapter or segment within a video.
type Chapter struct {
	// ID is the unique identifier for the chapter.
	ID int `json:"id"`
	// Title is the chapter's name.
	Title string `json:"title"`
	// Offset is the time offset in seconds where the chapter starts.
	Offset int `json:"offset"`
	// Game is the game being played in this chapter.
	Game Game `json:"game"`
}

// PreviousVideo contains information about the previous video in a series.
type PreviousVideo struct {
	// ID is the unique identifier for the video.
	ID int `json:"id"`
	// Title is the video's title.
	Title string `json:"title"`
	// CreatedAt is when the video was created.
	CreatedAt time.Time `json:"created_at"`
	// Season is the season number.
	Season int `json:"season"`
	// Episode is the episode number.
	Episode int `json:"episode"`
	// Views is the total view count.
	Views int `json:"views"`
	// VideoLength is the video duration in seconds.
	VideoLength int `json:"video_length"`
	// PreviewURL is the URL to the video preview image.
	PreviewURL string `json:"preview_url"`
}

// VideoInfo contains comprehensive information about a video.
type VideoInfo struct {
	// ID is the unique identifier for the video.
	ID int `json:"id"`
	// Title is the video's title.
	Title string `json:"title"`
	// CreatedAt is when the video was created.
	CreatedAt time.Time `json:"created_at"`
	// Season is the season number.
	Season int `json:"season"`
	// Episode is the episode number.
	Episode int `json:"episode"`
	// SourceFps is the original frames per second of the video.
	SourceFps int `json:"source_fps"`
	// SourceLength is the video duration in seconds.
	SourceLength int `json:"source_length"`
	// SourceWidth is the video width in pixels.
	SourceWidth int `json:"source_width"`
	// SourceHeight is the video height in pixels.
	SourceHeight int `json:"source_height"`
	// Views is the total view count.
	Views int `json:"views"`
	// PreviewURL is the URL to the video preview image.
	PreviewURL string `json:"preview_url"`
	// VttURL is the URL to the WebVTT subtitle/caption file.
	VttURL string `json:"vtt_url"`
	// SpriteURL is the URL to the thumbnail sprite sheet.
	SpriteURL string `json:"sprite_url"`
	// Tags is the list of tags associated with the video.
	Tags []Tag `json:"tags"`
	// ChatReplay is the URL or identifier for the chat replay.
	ChatReplay string `json:"chat_replay"`
	// Chapters is the list of chapters in the video.
	Chapters []Chapter `json:"chapters"`
	// Next contains information about the next video, if any.
	Next interface{} `json:"next"`
	// Previous contains information about the previous video, if any.
	Previous *PreviousVideo `json:"previous"`
}

// User represents a user account.
type User struct {
	// UserID is the unique identifier for the user.
	UserID int `json:"user_id"`
	// Loginname is the user's login name.
	Loginname string `json:"loginname"`
	// Username is the user's display name.
	Username string `json:"username"`
	// Bio is the user's biography/description.
	Bio string `json:"bio"`
	// AccessLevel indicates the user's access level or role.
	AccessLevel string `json:"access_level"`
	// Badges is the list of badges earned by the user.
	Badges []UserBadge `json:"badges"`
	// IsVerified indicates if the user account is verified.
	IsVerified bool `json:"is_verified"`
	// AvatarURL is the URL to the user's avatar image.
	AvatarURL string `json:"avatar_url"`
}

// UserBadge represents a badge or achievement earned by a user.
type UserBadge struct {
	// ID is the unique identifier for the badge.
	ID int `json:"id"`
	// Slot is the badge's display slot.
	Slot int `json:"slot"`
	// ActiveSpace indicates where the badge is active.
	ActiveSpace string `json:"active_space"`
	// Title is the badge's name.
	Title string `json:"title"`
	// Description explains what the badge represents.
	Description string `json:"description"`
	// GrantedAt is when the badge was awarded.
	GrantedAt time.Time `json:"granted_at"`
	// ExpiresAt is when the badge expires, if applicable.
	ExpiresAt time.Time `json:"expires_at"`
	// IsSelected indicates if the badge is currently displayed.
	IsSelected bool `json:"is_selected"`
	// Level is the badge level.
	Level int `json:"level"`
	// Version is the badge version.
	Version int `json:"version"`
	// Image is the URL to the badge image.
	Image string `json:"image"`
}

// Comment represents a user comment on a video.
type Comment struct {
	// ID is the unique identifier for the comment.
	ID int `json:"id"`
	// ParentCommentID is the ID of the parent comment if this is a reply.
	ParentCommentID int `json:"parent_comment_id"`
	// CreatedAt is when the comment was created.
	CreatedAt time.Time `json:"created_at"`
	// EditedAt is when the comment was last edited.
	EditedAt time.Time `json:"edited_at"`
	// VideoPlaybackOffset is the video timestamp (in seconds) where the comment was made.
	VideoPlaybackOffset int `json:"video_playback_offset"`
	// RepliesCount is the number of replies to this comment.
	RepliesCount int `json:"replies_count"`
	// UpvoteCount is the number of upvotes the comment received.
	UpvoteCount int `json:"upvote_count"`
	// DidUpvote indicates if the current user upvoted this comment.
	DidUpvote bool `json:"did_upvote"`
	// User is the comment author.
	User User `json:"user"`
	// Comment is the text content of the comment.
	Comment string `json:"comment"`
}

// DiscoveryType specifies the type of video discovery algorithm to use.
type DiscoveryType string

const (
	// DiscoveryTypeRecent returns the most recently uploaded videos.
	DiscoveryTypeRecent DiscoveryType = "recent"
	// DiscoveryTypeViews returns videos sorted by view count.
	DiscoveryTypeViews DiscoveryType = "views"
	// DiscoveryTypeSimilar returns videos similar to previously watched content.
	DiscoveryTypeSimilar DiscoveryType = "similar"
)

// VideoSearchResult represents a video in search or discovery results.
type VideoSearchResult struct {
	// ID is the unique identifier for the video.
	ID int `json:"id"`
	// Title is the video's title.
	Title string `json:"title"`
	// Season is the season number.
	Season int `json:"season"`
	// Episode is the episode number.
	Episode int `json:"episode"`
	// CreatedAt is when the video was created.
	CreatedAt time.Time `json:"created_at"`
	// VideoLength is the video duration in seconds.
	VideoLength int `json:"video_length"`
	// Views is the total view count.
	Views int `json:"views"`
	// PreviewURL is the URL to the video preview image.
	PreviewURL string `json:"preview_url"`
	// Tags is the list of tags associated with the video.
	Tags []Tag `json:"tags"`
}

// GameSearchResult represents a game in search results along with related videos.
type GameSearchResult struct {
	// ID is the unique identifier for the game.
	ID int `json:"id"`
	// Title is the game's name.
	Title string `json:"title"`
	// TwitchDetails contains Twitch-specific information about the game.
	TwitchDetails TwitchDetails `json:"twitch_details"`
	// Videos is the list of videos featuring this game.
	Videos []struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		CreatedAt   time.Time `json:"created_at"`
		Episode     int       `json:"episode"`
		PreviewURL  string    `json:"preview_url"`
		VideoLength int       `json:"video_length"`
		Views       int       `json:"views"`
		Tags        []Tag     `json:"tags"`
	} `json:"videos"`
}

// ChannelName is a type alias for channel names.
type ChannelName string

// ChannelInfo contains information about a Twitch channel's live status.
type ChannelInfo struct {
	// IsLive indicates if the channel is currently live.
	IsLive bool `json:"is_live"`
	// ChannelID is the Twitch channel ID.
	ChannelID string `json:"channel_id"`
	// ChannelDisplayName is the channel's display name.
	ChannelDisplayName string `json:"channel_display_name"`
	// GameID is the Twitch game ID being played.
	GameID string `json:"game_id"`
	// GameName is the name of the game being played.
	GameName string `json:"game_name"`
	// Title is the stream title.
	Title string `json:"title"`
	// ViewerCount is the current number of viewers.
	ViewerCount int `json:"viewer_count"`
	// StartedAt is when the stream started.
	StartedAt string `json:"started_at"`
	// ThumbnailURL is the URL to the stream thumbnail.
	ThumbnailURL string `json:"thumbnail_url"`
	// IsMature indicates if the stream is marked as mature content.
	IsMature bool `json:"is_mature"`
}

// LiveCheck contains the live status of multiple channels.
type LiveCheck struct {
	// Channels maps channel names to their live status information.
	Channels map[ChannelName]ChannelInfo `json:"channels"`
}

// SortDirection specifies the sort order direction.
type SortDirection string

const (
	// SortDirectionAsc sorts in ascending order.
	SortDirectionAsc SortDirection = "asc"
	// SortDirectionDesc sorts in descending order.
	SortDirectionDesc SortDirection = "desc"
)

// SortBy specifies the field to sort by.
type SortBy string

const (
	// SortByDate sorts by date.
	SortByDate SortBy = "date"
	// SortByViews sorts by view count.
	SortByViews SortBy = "views"
)

// SearchQuery contains parameters for search queries.
type SearchQuery struct {
	// Limit is the number of results to return.
	Limit *int16 `json:"first"`
	// Offset is the number of results to skip.
	Offset *int16 `json:"offset"`
	// Direction specifies the sort order.
	Direction *SortDirection `json:"direction"`
	// Query is the search query string.
	Query *string `json:"query"`
	// Tags is the list of tag IDs to filter by.
	Tags *[]int8 `json:"tags"`
	// Sort specifies the field to sort by.
	Sort *SortBy `json:"sort"`
}
