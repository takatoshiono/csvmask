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
	if len(chars) <= n {
		return strings.Repeat(c, len(chars))
	}
	return string(chars[0:len(chars)-n]) + strings.Repeat(c, n)
}

// left masks n characters to the left of s with c.
func left(n int, c, s string) string {
	if s == "" {
		return s
	}
	chars := []rune(s)
	if len(chars) <= n {
		return strings.Repeat(c, len(chars))
	}
	return strings.Repeat(c, n) + string(chars[n:])
}
