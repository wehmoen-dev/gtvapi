package gtvapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"resty.dev/v3"
)

const userAgent = "gtvapi/1.0 (+https://github.com/wehmoen/gtvapi)"
const apiBaseUrl = "https://api.gronkh.tv"
const apiVersion = "v1"

type GronkhTV struct {
	client *resty.Client
}

func NewClient(debug bool) *GronkhTV {

	return &GronkhTV{
		client: resty.
			New().
			SetBaseURL(fmt.Sprintf("%s/%s", apiBaseUrl, apiVersion)).
			SetHeader("User-Agent", userAgent).
			SetDebug(debug),
	}
}

func (c *GronkhTV) AllTags() (*[]Tag, error) {
	result, err := c.get(tagsAllEndpoint, nil)

	if err != nil {
		return nil, err
	}

	var tags []Tag

	err = json.Unmarshal(result, &tags)

	if err != nil {
		return nil, err
	}

	return &tags, nil
}

func (c *GronkhTV) LiveCheck() (*map[ChannelName]ChannelInfo, error) {
	var livecheck LiveCheck
	result, err := c.get(externalTwitchLivecheckEndpoint, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &livecheck)

	if err != nil {
		return nil, err
	}

	return &livecheck.Channels, nil

}

func (c *GronkhTV) Discover(kind DiscoveryType) (*[]VideoSearchResult, error) {

	var discovery struct {
		Discovery []VideoSearchResult `json:"discovery"`
	}

	var result []byte
	var err error

	switch kind {

	case DiscoveryTypeRecent:
		result, err = c.get(Endpoint(fmt.Sprintf("%s/%s", videoDiscoveryEndpoint, DiscoveryTypeRecent)), nil)
		break
	case DiscoveryTypeViews:
		result, err = c.get(Endpoint(fmt.Sprintf("%s/%s", videoDiscoveryEndpoint, DiscoveryTypeViews)), nil)
		break
	case DiscoveryTypeSimilar:
		result, err = c.get(Endpoint(fmt.Sprintf("%s/%s", videoDiscoveryEndpoint, DiscoveryTypeSimilar)), nil)
		break
	default:
		return nil, fmt.Errorf("unknown discovery type: %s", kind)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &discovery)

	if err != nil {
		return nil, err
	}

	return &discovery.Discovery, nil

}

func (c *GronkhTV) VideoComments(episode int16) (*[]Comment, error) {
	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	result, err := c.get(videoCommentsEndpoint, query)

	if err != nil {
		return nil, err
	}

	var comments struct {
		Comments []Comment `json:"comments"`
	}

	err = json.Unmarshal(result, &comments)

	if err != nil {
		return nil, err
	}

	return &comments.Comments, nil
}

func (c *GronkhTV) VideoPlaylist(episode int16) (*string, error) {
	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	result, err := c.get(videoPlaylistEndpoint, query)

	if err != nil {
		return nil, err
	}

	var playlistResult struct {
		PlaylistUrl string `json:"playlist_url"`
	}

	err = json.Unmarshal(result, &playlistResult)

	if err != nil {
		return nil, err
	}

	return &playlistResult.PlaylistUrl, nil
}

func (c *GronkhTV) VideoInfo(episode int16) (*VideoInfo, error) {

	query := map[string]string{
		"episode": fmt.Sprintf("%d", episode),
	}

	result, err := c.get(videoInfoEndpoint, query)

	if err != nil {
		return nil, err
	}

	var info *VideoInfo

	err = json.Unmarshal(result, &info)

	return info, err
}

func (c *GronkhTV) get(endpoint Endpoint, query map[string]string) ([]byte, error) {

	requesturl, err := url.Parse(fmt.Sprintf("%s/%s", c.client.BaseURL(), endpoint))

	if err != nil {
		return nil, err
	}

	if query != nil {
		q := requesturl.Query()
		for key, value := range query {
			q.Set(key, value)
		}
		requesturl.RawQuery = q.Encode()
	}

	res, err := c.client.R().Get(requesturl.String())

	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
