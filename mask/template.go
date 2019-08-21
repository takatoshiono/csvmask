package mask

import (
	"fmt"
	"io"
	"text/template"
)

type (
	// Template wraps template.Template of text/template package.
	Template interface {
		Execute(wr io.Writer, data interface{}) error
	}
)

// NewTemplate returns a new template.
func NewTemplate(text string) (Template, error) {
	t, err := template.New("mask").Funcs(funcMap()).Parse(text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template text: %v", err)
	}
	return t, nil
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"hash":     hash,
		"checksum": checksum,
	}
}
