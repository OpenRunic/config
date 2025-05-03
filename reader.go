package config

import (
	"fmt"
	"io"
	"os"
)

// configuration reader interface
type Reader interface {
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
func (r BaseReader) Parse(_ *Options, _ []*Field) error {
	return nil
}

// resolve file path for config file
func ResolvePath(r Reader, opts *Options, path string) string {
	if len(path) > 0 {
		return path
	}

	ext := r.ConfigExtension()
	if len(ext) > 0 {
		return fmt.Sprintf("%s.%s", opts.Filename, ext)
	}
	return opts.Filename
}

// resolve byte data from file or reader
func ResolveBytes(r Reader, opts *Options, path string, reader io.Reader) ([]byte, error) {
	if reader != nil {
		return io.ReadAll(reader)
	}

	fPath := ResolvePath(r, opts, path)
	oFile, err := os.Open(fPath)
	if err != nil {
		if opts.Strict {
			return nil, fmt.Errorf("load config failed: %s", fPath)
		}
		return nil, nil
	}
	defer oFile.Close()

	return io.ReadAll(oFile)
}
