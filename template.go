package main

import (
	"fmt"
	"io"
	"text/template"
)

type (
	// Template wraps template.Template of text/template package.
	Template interface {
		Execute(wr io.Writer, data interface{}) error
		CloneWithEcho() (Template, error)
	}

	templateImpl struct {
		template *template.Template
	}
)

var (
	funcMap = template.FuncMap{
		"hash":     hash,
		"checksum": checksum,
		"right":    right,
		"left":     left,
	}

	echoFuncMap = template.FuncMap{
		"hash":     echo,
		"checksum": echo,
		"right":    echo,
		"left":     echo,
	}
)

// NewTemplate returns a new template.
func NewTemplate(text string) (Template, error) {
	tmpl, err := template.New("mask").Funcs(funcMap).Parse(text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template text: %v", err)
	}
	return &templateImpl{template: tmpl}, nil
}

// Execute just calls Execute() of text/template.
func (t *templateImpl) Execute(wr io.Writer, data interface{}) error {
	return t.template.Execute(wr, data)
}

// CloneWithEcho clones template and replace all funcs with echo.
func (t *templateImpl) CloneWithEcho() (Template, error) {
	tmpl, err := t.template.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone template: %v", err)
	}
	return &templateImpl{template: tmpl.Funcs(echoFuncMap)}, nil
}

func echo(s string) string {
	return s
}
