package main

import (
	"io"
)

// Project struct is used to read the project config
type Project struct {
	Name       string
	PathPrefix string
	Resources  []Resource
}

// Resource reqpresents the individual resource for a project
type Resource struct {
	Methods      []string
	Subresources []Resource
}

func (p *Project) load(r io.Reader, deocder Decoder) error {
	return deocder.Decode(r, p)
}
