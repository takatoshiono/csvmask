package main

import "testing"

func TestRight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		c, s string
		want string
	}{
		{"6 chars with x", 6, "x", "東京都港区芝公園4丁目2-8", "東京都港区芝公園xxxxxx"},
		{"0 chars with x", 0, "x", "東京都港区芝公園4丁目2-8", "東京都港区芝公園4丁目2-8"},
		{"6 chars with x lack of str", 6, "x", "東京都", "xxx"},
		{"3 chars with empty", 3, "", "東京都港区芝公園4丁目2-8", "東京都港区芝公園4丁目"},
		{"empty", 1, "x", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := right(tt.n, tt.c, tt.s); got != tt.want {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		name string
		n    int
		c, s string
		want string
	}{
		{"12 chars with x", 12, "x", "4 Chome-2-8 Shibakoen, Minato City, Tokyo", "xxxxxxxxxxxxShibakoen, Minato City, Tokyo"},
		{"0 chars with x", 0, "x", "4 Chome-2-8 Shibakoen, Minato City, Tokyo", "4 Chome-2-8 Shibakoen, Minato City, Tokyo"},
		{"6 chars with x lack of str", 6, "x", "Tokyo", "xxxxx"},
		{"12 chars with empty", 12, "", "4 Chome-2-8 Shibakoen, Minato City, Tokyo", "Shibakoen, Minato City, Tokyo"},
		{"empty", 1, "x", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := left(tt.n, tt.c, tt.s); got != tt.want {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}
