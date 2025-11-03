package gtvapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_LiveCheck(t *testing.T) {
	tests := []struct {
		name       string
		response   interface{}
		statusCode int
		wantErr    bool
		wantLen    int
	}{
		{
			name: "successful request with live channels",
			response: LiveCheck{
				Channels: map[ChannelName]*ChannelInfo{
					"gronkh": {
						IsLive:             true,
						ChannelID:          "12345",
						ChannelDisplayName: "Gronkh",
						GameName:           "Minecraft",
						ViewerCount:        5000,
					},
					"pandorya": {
						IsLive:             false,
						ChannelID:          "67890",
						ChannelDisplayName: "Pandorya",
					},
				},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    2,
		},
		{
			name: "empty channels",
			response: LiveCheck{
				Channels: map[ChannelName]*ChannelInfo{},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    0,
		},
		{
			name:       "server error",
			response:   map[string]string{"error": "internal error"},
			statusCode: http.StatusInternalServerError,
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
			channels, err := client.LiveCheck(ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("LiveCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(channels) != tt.wantLen {
				t.Errorf("LiveCheck() len = %v, want %v", len(channels), tt.wantLen)
			}

			if !tt.wantErr && tt.wantLen > 0 {
				expected := tt.response.(LiveCheck)
				gronkhInfo, ok := channels["gronkh"]
				if !ok {
					t.Error("LiveCheck() missing 'gronkh' channel")
				} else {
					expectedGronkh := expected.Channels["gronkh"]
					if gronkhInfo.IsLive != expectedGronkh.IsLive {
						t.Errorf("LiveCheck() gronkh.IsLive = %v, want %v", gronkhInfo.IsLive, expectedGronkh.IsLive)
					}
					if gronkhInfo.ViewerCount != expectedGronkh.ViewerCount {
						t.Errorf("LiveCheck() gronkh.ViewerCount = %v, want %v", gronkhInfo.ViewerCount, expectedGronkh.ViewerCount)
					}
				}
			}
		})
	}
}

func TestClient_LiveCheck_ContextCancellation(t *testing.T) {
	server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LiveCheck{
			Channels: map[ChannelName]*ChannelInfo{},
		})
	})

	client := NewClient(&ClientConfig{
		BaseURL:    server.URL,
		MaxRetries: 0,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := client.LiveCheck(ctx)
	if err == nil {
		t.Error("expected error from canceled context, got nil")
	}
}

// Benchmark test
func BenchmarkClient_LiveCheck(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LiveCheck{
			Channels: map[ChannelName]*ChannelInfo{
				"gronkh": {
					IsLive:      true,
					ViewerCount: 5000,
				},
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
		_, _ = client.LiveCheck(ctx)
	}
}

// Example test
func ExampleClient_LiveCheck() {
	client := NewClient(nil)
	ctx := context.Background()

	channels, err := client.LiveCheck(ctx)
	if err != nil {
		// Handle error
		return
	}

	for name, info := range channels {
		if info.IsLive {
			_ = name // Channel is live
		}
	}
}
