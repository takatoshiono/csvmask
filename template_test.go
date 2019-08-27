package main

import (
	"testing"
)

func TestNewTemplate(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr bool
	}{
		{"no func", "{{.Field1}},{{.Field2}},{{.Field3}}", false},
		{"hash func", "{{.Field1}},{{hash .Field2}},{{.Field3}}", false},
		{"maskright func", `{{.Field1}},{{maskright "3" "x" .Field2}},{{.Field3}}`, false},
		{"checksum func", "{{.Field1}},{{checksum .Field2}},{{.Field3}}", false},
		{"invalid syntax", "{{.Field1},{{ .Field2}}", true},
		{"invalid func", "{{.Field1}},{{foofunc .Field2}}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := NewTemplate(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Fatal("want error, but got no error")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
			}
			if tt.wantErr {
				return
			}
			if tmpl == nil {
				t.Error("want template not nil, but got nil")
			}
		})
	}
}
