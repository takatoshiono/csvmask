package mask_test

import (
	"io"
	"os"
	"testing"

	"github.com/takatoshiono/csvmask/mask"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		template string
		wants    []string
	}{
		{
			name:     "valid",
			file:     "../testdata/test.csv",
			template: "{{.Field1}},{{hash .Field2}},{{.Field3}}",
			wants: []string{
				"ID,3NHVIj9zs6llwH4/9dvuPu3P7bgGaGoFubOGiiw9bVA,Address",
				"4085ff59-39bd-4cc3-8a55-c5b1c6785922,PbZ8hc4alo56RYc9/m+vECyVdjHqZRGMlxUGigh3/uE,Kirkcaldy United Kingdom",
			},
		},
		// TODO: add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template, err := mask.NewTemplate(tt.template)
			if err != nil {
				t.Fatal(err)
			}

			f, err := os.Open(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			r := mask.NewReader(f, template)
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
