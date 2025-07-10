package gtvapi

import "testing"

const testEpisodeNumber = 1000
const testVideoId = 2596

var client *GronkhTV

func init() {
	client = NewClient(false)
}

func TestGronkhTV_VideoInfo(t *testing.T) {

	info, err := client.VideoInfo(testEpisodeNumber)

	if err != nil {
		t.Error(err)
	}

	if info.ID != testVideoId {

		t.Errorf("VideoInfo does not match. Got %d, expected %d", info.ID, testVideoId)
	}

}

func TestGronkhTV_VideoComments(t *testing.T) {
	comments, err := client.VideoComments(testEpisodeNumber)
	if err != nil {
		t.Error(err)
	}

	if len(*comments) == 0 || (*comments)[0].Comment == "" {
		t.Errorf("No comments found for the video %d or first comment empty/invalid", testEpisodeNumber)
	}

}

func TestGronkhTV_VideoDiscovery(t *testing.T) {
	tags, err := client.AllTags()

	if err != nil {
		t.Error(err)
	}

	if len(*tags) == 0 {
		t.Errorf("No tags found")
	}
}

func TestGronkhTV_ExternalTwitchLivecheck(t *testing.T) {
	channels, err := client.LiveCheck()

	if err != nil {
		t.Error(err)
	}

	if len(*channels) == 0 {
		t.Errorf("No channels found")
	}
}

func TestGronkhTV_ExternalTwitchLiveCheck(t *testing.T) {
	playlist, err := client.VideoPlaylist(testEpisodeNumber)

	if err != nil {
		t.Error(err)
	}

	if playlist == nil {
		t.Errorf("No playlist found for video %d", testEpisodeNumber)
	}
}

func TestGronkhTV_Discover(t *testing.T) {
	for _, kind := range []DiscoveryType{DiscoveryTypeViews, DiscoveryTypeRecent, DiscoveryTypeSimilar} {
		results, err := client.Discover(kind)

		if err != nil {
			t.Error(err)
		}

		if len(*results) == 0 {
			t.Errorf("No results found for discovery type %s", kind)
		}
	}

}
