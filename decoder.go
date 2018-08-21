package main

import (
	"errors"
	"io"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Decoder interface {
	Decode(reader io.Reader, output interface{}) error
}

func InitDecoder(filename string) (Decoder, error) {
	switch filepath.Ext(filename) {
	case ".toml":
		return new(tomlDecoder), nil
	default:
		return nil, errors.New("Input Filetype not supported")
	}
}

type tomlDecoder struct{}

func (t *tomlDecoder) Decode(reader io.Reader, output interface{}) error {
	_, err := toml.DecodeReader(reader, output)
	if err != nil {
		return err
	}
	return nil
}
