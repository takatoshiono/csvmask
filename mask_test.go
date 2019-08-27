package main

import "testing"

func TestRight(t *testing.T) {
	tests := []struct {
		name    string
		n, c, s string
		want    string
		wantErr bool
	}{
		{"6 chars with x", "6", "x", "東京都港区芝公園4丁目2-8", "東京都港区芝公園xxxxxx", false},
		{"0 chars with x", "0", "x", "東京都港区芝公園4丁目2-8", "東京都港区芝公園4丁目2-8", false},
		{"3 chars with empty", "3", "", "東京都港区芝公園4丁目2-8", "東京都港区芝公園4丁目", false},
		{"invalid n", "invalid", "x", "東京都港区芝公園4丁目2-8", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := right(tt.n, tt.c, tt.s)
			if tt.wantErr {
				if err == nil {
					t.Errorf("want error, but got no error")
				}
			} else {
				if err != nil {
					t.Errorf("want no error, but got %v", err)
				}
			}
			if got != tt.want {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}
