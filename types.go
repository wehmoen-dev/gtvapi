package gtvapi

import "time"

type TwitchDetails struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type Tag struct {
	ID    int8   `json:"id"`
	Title string `json:"title"`
}

type Game struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	TwitchDetails TwitchDetails `json:"twitch_details"`
}
type Chapter struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Offset int    `json:"offset"`
	Game   Game   `json:"game"`
}

type PreviousVideo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Season      int       `json:"season"`
	Episode     int       `json:"episode"`
	Views       int       `json:"views"`
	VideoLength int       `json:"video_length"`
	PreviewURL  string    `json:"preview_url"`
}

type VideoInfo struct {
	ID           int            `json:"id"`
	Title        string         `json:"title"`
	CreatedAt    time.Time      `json:"created_at"`
	Season       int            `json:"season"`
	Episode      int            `json:"episode"`
	SourceFps    int            `json:"source_fps"`
	SourceLength int            `json:"source_length"`
	SourceWidth  int            `json:"source_width"`
	SourceHeight int            `json:"source_height"`
	Views        int            `json:"views"`
	PreviewURL   string         `json:"preview_url"`
	VttURL       string         `json:"vtt_url"`
	SpriteURL    string         `json:"sprite_url"`
	Tags         []Tag          `json:"tags"`
	ChatReplay   string         `json:"chat_replay"`
	Chapters     []Chapter      `json:"chapters"`
	Next         interface{}    `json:"next"`
	Previous     *PreviousVideo `json:"previous"`
}

type User struct {
	UserID      int         `json:"user_id"`
	Loginname   string      `json:"loginname"`
	Username    string      `json:"username"`
	Bio         string      `json:"bio"`
	AccessLevel string      `json:"access_level"`
	Badges      []UserBadge `json:"badges"`
	IsVerified  bool        `json:"is_verified"`
	AvatarURL   string      `json:"avatar_url"`
}

type UserBadge struct {
	ID          int       `json:"id"`
	Slot        int       `json:"slot"`
	ActiveSpace string    `json:"active_space"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	GrantedAt   time.Time `json:"granted_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	IsSelected  bool      `json:"is_selected"`
	Level       int       `json:"level"`
	Version     int       `json:"version"`
	Image       string    `json:"image"`
}

type Comment struct {
	ID                  int       `json:"id"`
	ParentCommentID     int       `json:"parent_comment_id"`
	CreatedAt           time.Time `json:"created_at"`
	EditedAt            time.Time `json:"edited_at"`
	VideoPlaybackOffset int       `json:"video_playback_offset"`
	RepliesCount        int       `json:"replies_count"`
	UpvoteCount         int       `json:"upvote_count"`
	DidUpvote           bool      `json:"did_upvote"`
	User                User      `json:"user"`
	Comment             string    `json:"comment"`
}

type DiscoveryType string

const (
	DiscoveryTypeRecent  DiscoveryType = "recent"
	DiscoveryTypeViews   DiscoveryType = "views"
	DiscoveryTypeSimilar DiscoveryType = "similar"
)

type VideoSearchResult struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Season      int       `json:"season"`
	Episode     int       `json:"episode"`
	CreatedAt   time.Time `json:"created_at"`
	VideoLength int       `json:"video_length"`
	Views       int       `json:"views"`
	PreviewURL  string    `json:"preview_url"`
	Tags        []Tag     `json:"tags"`
}
type GameSearchResult struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	TwitchDetails TwitchDetails `json:"twitch_details"`
	Videos        []struct {
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

type ChannelName string

type ChannelInfo struct {
	IsLive             bool   `json:"is_live"`
	ChannelID          string `json:"channel_id"`
	ChannelDisplayName string `json:"channel_display_name"`
	GameID             string `json:"game_id"`
	GameName           string `json:"game_name"`
	Title              string `json:"title"`
	ViewerCount        int    `json:"viewer_count"`
	StartedAt          string `json:"started_at"`
	ThumbnailURL       string `json:"thumbnail_url"`
	IsMature           bool   `json:"is_mature"`
}

type LiveCheck struct {
	Channels map[ChannelName]ChannelInfo `json:"channels"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

type SortBy string

const (
	SortByDate  SortBy = "date"
	SortByViews SortBy = "views"
)

type SearchQuery struct {
	// Number of results returned
	Limit *int16 `json:"first"`
	// Offset from the start of all results
	Offset    *int16         `json:"offset"`
	Direction *SortDirection `json:"direction"`
	// Search query
	Query *string `json:"query"`
	// Slice of Tag ids to include in results
	Tags *[]int8 `json:"tags"`
	// What to sort the results by
	Sort *SortBy `json:"sort"`
}
