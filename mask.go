package main

import (
	"strings"
)

// right masks n characters to the right of s with c.
func right(n int, c, s string) string {
	if s == "" {
		return s
	}
	chars := []rune(s)
	return string(chars[0:len(chars)-int(n)]) + strings.Repeat(c, int(n))
}

// left masks n characters to the left of s with c.
func left(n int, c, s string) string {
	if s == "" {
		return s
	}
	chars := []rune(s)
	return strings.Repeat(c, int(n)) + string(chars[int(n):])
}
