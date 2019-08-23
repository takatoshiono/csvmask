package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

// hash returns hashed string of s.
func hash(s string) string {
	if s == "" {
		return s
	}
	sum := sha256.Sum256([]byte(s))
	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding.WithPadding(base64.NoPadding), &buf)
	encoder.Write(sum[:])
	encoder.Close()
	return buf.String()
}
