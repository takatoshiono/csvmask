package main

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"not empty", "foobar", "w6uP8Tcg6K2QR905Rms8iXTlksL6OD1KOWBxTK7wxPI"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.s); got != tt.want {
				t.Errorf("hash(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
