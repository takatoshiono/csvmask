package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name       string
		file       string
		template   string
		skipHeader bool
		wants      []string
	}{
		{
			name:       "test.csv",
			file:       "./testdata/test.csv",
			template:   "{{.Field1}},{{hash .Field2}},{{.Field3}}",
			skipHeader: false,
			wants: []string{
				"ID,3NHVIj9zs6llwH4/9dvuPu3P7bgGaGoFubOGiiw9bVA,Address",
				"4085ff59-39bd-4cc3-8a55-c5b1c6785922,PbZ8hc4alo56RYc9/m+vECyVdjHqZRGMlxUGigh3/uE,Kirkcaldy United Kingdom",
			},
		},
		{
			name:       "test.csv skip header",
			file:       "./testdata/test.csv",
			template:   "{{.Field1}},{{hash .Field2}},{{.Field3}}",
			skipHeader: true,
			wants: []string{
				"ID,Name,Address",
				"4085ff59-39bd-4cc3-8a55-c5b1c6785922,PbZ8hc4alo56RYc9/m+vECyVdjHqZRGMlxUGigh3/uE,Kirkcaldy United Kingdom",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template, err := NewTemplate(tt.template)
			if err != nil {
				t.Fatal(err)
			}

			f, err := os.Open(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			r := NewReader(f, template)
			if tt.skipHeader {
				r.SkipHeader = true
			}

			for i := 0; ; i++ {
				got, err := r.Read()
				if err != nil {
					if err == io.EOF {
						break
					}
					t.Fatal(err)
				}
				if got != tt.wants[i] {
					t.Errorf("want read %v, but got %v", tt.wants[i], got)
				}
			}
		})
	}
}

func TestReadFuncs(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		template string
		want     string
	}{
		{"hash", "foo,bar", "{{hash .Field1}},{{.Field2}}", "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564,bar"},
		{"checksum", "foo,bar", "{{checksum .Field1}},{{.Field2}}", "8c736521,bar"},
		{"maskright", "foo,東京都港区芝公園4丁目2-8", `{{.Field1}},{{maskright "6" "x" .Field2}}`, "foo,東京都港区芝公園xxxxxx"},
		{"maskright pipeline", "foo,東京都港区芝公園4丁目2-8", `{{.Field1}},{{.Field2 | maskright "6" "x"}}`, "foo,東京都港区芝公園xxxxxx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template, err := NewTemplate(tt.template)
			if err != nil {
				t.Fatal(err)
			}
			r := NewReader(bytes.NewBufferString(tt.str), template)
			got, err := r.Read()
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("want read %v, but got %v", tt.want, got)
			}
		})
	}
}
