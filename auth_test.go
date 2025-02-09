package main

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "correct password",
			password: "anirudhisreallycool",
			want:     true,
		},
		{
			name:     "incorrect password",
			password: "wrongpassword",
			want:     false,
		},
		{
			name:     "empty password",
			password: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authenticate(tt.password); got != tt.want {
				t.Errorf("authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
