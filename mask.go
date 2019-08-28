package main

import (
	"fmt"
	"strconv"
	"strings"
)

// right masks n characters to the right of s with c.
func right(n, c, s string) (string, error) {
	if s == "" {
		return s, nil
	}
	cnt, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return "", fmt.Errorf("failed to parse int: %v", err)
	}
	chars := []rune(s)
	return string(chars[0:len(chars)-int(cnt)]) + strings.Repeat(c, int(cnt)), nil
}

// left masks n characters to the left of s with c.
func left(n, c, s string) (string, error) {
	if s == "" {
		return s, nil
	}
	cnt, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return "", fmt.Errorf("failed to parse int: %v", err)
	}
	chars := []rune(s)
	return strings.Repeat(c, int(cnt)) + string(chars[cnt:]), nil
}
