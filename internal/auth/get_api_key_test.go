package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
	}{
		{
			name: "Valid key",
			headers: http.Header{
				"Authorization": []string{"ApiKey secret-token-123"},
			},
			expectedKey: "secret-token-123",
		},
		{
			name: "Missing prefix",
			headers: http.Header{
				"Authorization": []string{"secret-token-123"},
			},
			expectedKey: "",
		},
		{
			name: "Wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer secret-token-123"},
			},
			expectedKey: "",
		},
		{
			name: "No key",
			headers: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			expectedKey: "",
		},
		{
			name:        "Empty header",
			headers:     http.Header{},
			expectedKey: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, _ := GetAPIKey(tt.headers)

			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.expectedKey)
			}
		})
	}
}
