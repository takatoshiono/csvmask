package main

import "testing"

func TestChecksum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"not empty", "foobar", "9ef61f95"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum(tt.s); got != tt.want {
				t.Errorf("checksum(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
