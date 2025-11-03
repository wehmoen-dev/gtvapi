package gtvapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// mockServer creates a test HTTP server that returns predefined responses.
func mockServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return server
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name   string
		config *ClientConfig
		want   string // expected baseURL
	}{
		{
			name:   "nil config uses defaults",
			config: nil,
			want:   DefaultBaseURL,
		},
		{
			name: "custom config",
			config: &ClientConfig{
				BaseURL:   "https://custom.api.com/v2",
				UserAgent: "custom-agent/1.0",
				Timeout:   10 * time.Second,
			},
			want: "https://custom.api.com/v2",
		},
		{
			name: "config with defaults filled",
			config: &ClientConfig{
				Debug: true,
			},
			want: DefaultBaseURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.config)
			if client == nil {
				t.Fatal("NewClient returned nil")
			}
			if client.baseURL != tt.want {
				t.Errorf("baseURL = %v, want %v", client.baseURL, tt.want)
			}
		})
	}
}

func TestClient_get(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		query      map[string]string
		response   string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful request",
			endpoint:   "test",
			query:      nil,
			response:   `{"success": true}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "with query parameters",
			endpoint:   "test",
			query:      map[string]string{"episode": "1000"},
			response:   `{"episode": 1000}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "404 error",
			endpoint:   "notfound",
			query:      nil,
			response:   `{"error": "not found"}`,
			statusCode: http.StatusNotFound,
			wantErr:    true,
		},
		{
			name:       "500 error",
			endpoint:   "error",
			query:      nil,
			response:   `{"error": "internal server error"}`,
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.response))
			})

			client := NewClient(&ClientConfig{
				BaseURL:    server.URL,
				MaxRetries: 0, // Disable retries for testing
			})

			ctx := context.Background()
			data, err := client.get(ctx, tt.endpoint, tt.query)

			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && string(data) != tt.response {
				t.Errorf("get() data = %v, want %v", string(data), tt.response)
			}
		})
	}
}

func TestClient_get_WithContext(t *testing.T) {
	t.Run("context cancellation", func(t *testing.T) {
		server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"success": true}`))
		})

		client := NewClient(&ClientConfig{
			BaseURL:    server.URL,
			MaxRetries: 0,
		})

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := client.get(ctx, "test", nil)
		if err == nil {
			t.Error("expected error from canceled context, got nil")
		}
	})

	t.Run("context timeout", func(t *testing.T) {
		server := mockServer(t, func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(200 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"success": true}`))
		})

		client := NewClient(&ClientConfig{
			BaseURL:    server.URL,
			MaxRetries: 0,
			Timeout:    50 * time.Millisecond,
		})

		ctx := context.Background()
		_, err := client.get(ctx, "test", nil)
		if err == nil {
			t.Error("expected timeout error, got nil")
		}
	})
}

func TestUnmarshalResponse(t *testing.T) {
	type testStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name:    "valid JSON",
			data:    []byte(`{"name": "test", "value": 42}`),
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			data:    []byte(`{"name": "test", "value": `),
			wantErr: true,
		},
		{
			name:    "empty data",
			data:    []byte(``),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testStruct
			err := unmarshalResponse(tt.data, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_UserAgent(t *testing.T) {
	var receivedUA string
	server := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		receivedUA = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	})

	client := NewClient(&ClientConfig{
		BaseURL:    server.URL,
		UserAgent:  "test-agent/1.0",
		MaxRetries: 0,
	})

	ctx := context.Background()
	_, err := client.get(ctx, "test", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedUA != "test-agent/1.0" {
		t.Errorf("User-Agent = %v, want %v", receivedUA, "test-agent/1.0")
	}
}

// Benchmark tests
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClient(nil)
	}
}

func BenchmarkClient_get(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":    1000,
			"title": "Test Video",
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
		_, _ = client.get(ctx, "test", nil)
	}
}
