package auth

import (
	"fmt"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string][]string
		want    string
		wantErr error
	}{
		{
			name:    "no auth header",
			headers: map[string][]string{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed auth header",
			headers: map[string][]string{
				"Authorization": {"notApiKey"},
			},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
		{
			name: "valid auth header",
			headers: map[string][]string{
				"Authorization": {"ApiKey my-api-key"},
			},
			want:    "my-api-key",
			wantErr: nil,
		},
	}

	fmt.Println(tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if gotErr != tt.wantErr {
				t.Errorf("GetAPIKey() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
