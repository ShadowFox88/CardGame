package main

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{
			name:     "Test Valid Username",
			username: "John",
			want:     true,
		},
		{
			name:     "Test Invalid Username",
			username: "123456789123456",
			want:     false,
		},
		{
			name:     "Test Empty Username",
			username: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidUsername(tt.username); got != tt.want {
				t.Errorf("ValidUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
