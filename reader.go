package config

import (
	"fmt"
	"io"
	"os"
)

// configuration reader interface
type ConfigReader interface {
	Configurator() string
	ConfigExtension() string
	Parse(opts *Options, fields []*Field) error
	Get(opts *Options, field *Field) (any, bool)
}

// base reader to extend with (optional)
type BaseReader struct{}

// file extension for related reader
func (r BaseReader) ConfigExtension() string {
	return ""
}

// placeholder parser func since not every reader needs to parse
func (r BaseReader) Parse(opts *Options, fields []*Field) error {
	return nil
}

// resolve file path for config file
func (r BaseReader) ResolvePath(opts *Options, path string) string {
	if len(path) > 0 {
		return path
	}

	ext := r.ConfigExtension()
	if len(ext) > 0 {
		return fmt.Sprintf("%s.%s", opts.Filename, r.ConfigExtension())
	}
	return opts.Filename
}

// resolve byte data from file or reader
func (r BaseReader) ResolveBytes(opts *Options, path string, reader io.Reader) ([]byte, error) {
	if reader != nil {
		return io.ReadAll(reader)
	}

	fPath := r.ResolvePath(opts, path)
	jsonFile, err := os.Open(fPath)
	if err != nil {
		if opts.Strict {
			return nil, fmt.Errorf("unable to load config file: %s", fPath)
		}
		return nil, nil
	}
	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}
