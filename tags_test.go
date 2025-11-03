package gtvapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_AllTags(t *testing.T) {
	tests := []struct {
		name       string
		response   interface{}
		statusCode int
		wantErr    bool
		wantLen    int
	}{
		{
			name: "successful request",
			response: []Tag{
				{ID: 1, Title: "Action"},
				{ID: 2, Title: "Adventure"},
				{ID: 3, Title: "RPG"},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
			wantLen:    3,
		},
		{
			name:       "empty tags",
			response:   []Tag{},
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
			tags, err := client.AllTags(ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("AllTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(tags) != tt.wantLen {
				t.Errorf("AllTags() len = %v, want %v", len(tags), tt.wantLen)
			}

			if !tt.wantErr && tt.wantLen > 0 {
				expected := tt.response.([]Tag)
				if tags[0].ID != expected[0].ID {
					t.Errorf("AllTags() first tag ID = %v, want %v", tags[0].ID, expected[0].ID)
				}
			}
		})
	}
}

// Benchmark test
func BenchmarkClient_AllTags(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]Tag{
			{ID: 1, Title: "Action"},
			{ID: 2, Title: "Adventure"},
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
		_, _ = client.AllTags(ctx)
	}
}

// Example test
func ExampleClient_AllTags() {
	client := NewClient(nil)
	ctx := context.Background()

	tags, err := client.AllTags(ctx)
	if err != nil {
		// Handle error
		return
	}

	_ = tags // Use the tags
}
