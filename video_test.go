package gtvapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_VideoInfo(t *testing.T) {
	tests := []struct {
		name       string
		episode    int
		response   interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:    "successful request",
			episode: 1000,
			response: VideoInfo{
				ID:           2596,
				Title:        "Test Video",
				Episode:      1000,
				Season:       1,
				Views:        12345,
				SourceLength: 3600,
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "invalid episode number",
			episode:    -1,
			response:   nil,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
		{
			name:       "zero episode number",
			episode:    0,
			response:   nil,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
		{
			name:       "server error",
			episode:    1000,
			response:   map[string]string{"error": "internal error"},
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
				if tt.episode > 0 {
					expectedEpisode := fmt.Sprintf("%d", tt.episode)
					if r.URL.Query().Get("episode") != expectedEpisode {
						t.Errorf("expected episode=%s, got %s", expectedEpisode, r.URL.Query().Get("episode"))
					}
				}
				w.WriteHeader(tt.statusCode)
				if tt.response != nil {
					_ = json.NewEncoder(w).Encode(tt.response)
				}
			})

			client := NewClient(&ClientConfig{
				BaseURL:    server.URL,
				MaxRetries: 0,
			})

			ctx := context.Background()
			info, err := client.VideoInfo(ctx, tt.episode)

			if (err != nil) != tt.wantErr {
				t.Errorf("VideoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && info == nil {
				t.Error("VideoInfo() returned nil info without error")
			}

			if !tt.wantErr && info != nil {
				expected := tt.response.(VideoInfo)
				if info.ID != expected.ID {
					t.Errorf("VideoInfo() ID = %v, want %v", info.ID, expected.ID)
				}
				if info.Title != expected.Title {
					t.Errorf("VideoInfo() Title = %v, want %v", info.Title, expected.Title)
				}
			}
		})
	}
}

func TestClient_VideoComments(t *testing.T) {
	tests := []struct {
		name       string
		episode    int
		response   interface{}
		statusCode int
		wantErr    bool
		wantLen    int
	}{
		{
			name:    "successful request with comments",
			episode: 1000,
			response: map[string]interface{}{
				"comments": []Comment{
					{
						ID:      1,
						Comment: "Great video!",
						User: User{
							UserID:   123,
							Username: "testuser",
						},
					},
					{
						ID:      2,
						Comment: "Thanks for sharing",
						User: User{
							UserID:   456,
							Username: "anotheruser",
						},
					},
				},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    2,
		},
		{
			name:    "empty comments",
			episode: 1000,
			response: map[string]interface{}{
				"comments": []Comment{},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    0,
		},
		{
			name:       "invalid episode",
			episode:    -5,
			response:   nil,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
				if tt.response != nil {
					_ = json.NewEncoder(w).Encode(tt.response)
				}
			})

			client := NewClient(&ClientConfig{
				BaseURL:    server.URL,
				MaxRetries: 0,
			})

			ctx := context.Background()
			comments, err := client.VideoComments(ctx, tt.episode)

			if (err != nil) != tt.wantErr {
				t.Errorf("VideoComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(comments) != tt.wantLen {
				t.Errorf("VideoComments() len = %v, want %v", len(comments), tt.wantLen)
			}
		})
	}
}

func TestClient_VideoPlaylist(t *testing.T) {
	tests := []struct {
		name       string
		episode    int
		response   interface{}
		statusCode int
		wantErr    bool
		wantURL    string
	}{
		{
			name:    "successful request",
			episode: 1000,
			response: map[string]string{
				"playlist_url": "https://example.com/playlist.m3u8",
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantURL:    "https://example.com/playlist.m3u8",
		},
		{
			name:       "invalid episode",
			episode:    0,
			response:   nil,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
				if tt.response != nil {
					_ = json.NewEncoder(w).Encode(tt.response)
				}
			})

			client := NewClient(&ClientConfig{
				BaseURL:    server.URL,
				MaxRetries: 0,
			})

			ctx := context.Background()
			url, err := client.VideoPlaylist(ctx, tt.episode)

			if (err != nil) != tt.wantErr {
				t.Errorf("VideoPlaylist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && url != tt.wantURL {
				t.Errorf("VideoPlaylist() url = %v, want %v", url, tt.wantURL)
			}
		})
	}
}

func TestClient_Discover(t *testing.T) {
	mockVideos := []VideoSearchResult{
		{
			ID:          1,
			Title:       "Video 1",
			Episode:     100,
			VideoLength: 3600,
			Views:       5000,
		},
		{
			ID:          2,
			Title:       "Video 2",
			Episode:     101,
			VideoLength: 4200,
			Views:       7500,
		},
	}

	tests := []struct {
		name          string
		discoveryType DiscoveryType
		response      interface{}
		statusCode    int
		wantErr       bool
		wantLen       int
	}{
		{
			name:          "recent videos",
			discoveryType: DiscoveryTypeRecent,
			response: map[string]interface{}{
				"discovery": mockVideos,
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    2,
		},
		{
			name:          "popular videos",
			discoveryType: DiscoveryTypeViews,
			response: map[string]interface{}{
				"discovery": mockVideos,
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    2,
		},
		{
			name:          "similar videos",
			discoveryType: DiscoveryTypeSimilar,
			response: map[string]interface{}{
				"discovery": mockVideos,
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    2,
		},
		{
			name:          "invalid discovery type",
			discoveryType: DiscoveryType("invalid"),
			response:      nil,
			statusCode:    http.StatusOK,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
				if tt.response != nil {
					_ = json.NewEncoder(w).Encode(tt.response)
				}
			})

			client := NewClient(&ClientConfig{
				BaseURL:    server.URL,
				MaxRetries: 0,
			})

			ctx := context.Background()
			videos, err := client.Discover(ctx, tt.discoveryType)

			if (err != nil) != tt.wantErr {
				t.Errorf("Discover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(videos) != tt.wantLen {
				t.Errorf("Discover() len = %v, want %v", len(videos), tt.wantLen)
			}
		})
	}
}

// Benchmark tests
func BenchmarkClient_VideoInfo(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(VideoInfo{
			ID:      2596,
			Title:   "Test Video",
			Episode: 1000,
		})
	}))
	defer server.Close()

	client := NewClient(&ClientConfig{
		BaseURL:    server.URL,
		MaxRetries: 0,
	})
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.VideoInfo(ctx, 1000)
	}
}

func BenchmarkClient_VideoComments(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"comments": []Comment{
				{ID: 1, Comment: "Test comment"},
			},
		})
	}))
	defer server.Close()

	client := NewClient(&ClientConfig{
		BaseURL:    server.URL,
		MaxRetries: 0,
	})
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.VideoComments(ctx, 1000)
	}
}

// Example tests
func ExampleClient_VideoInfo() {
	client := NewClient(nil)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	videoInfo, err := client.VideoInfo(ctx, 1000)
	if err != nil {
		// Handle error
		return
	}

	_ = videoInfo.Title // Use the video info
}

func ExampleClient_Discover() {
	client := NewClient(nil)
	ctx := context.Background()

	videos, err := client.Discover(ctx, DiscoveryTypeRecent)
	if err != nil {
		// Handle error
		return
	}

	_ = videos // Use the discovered videos
}
