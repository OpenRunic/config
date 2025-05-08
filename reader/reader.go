package reader

import (
	"fmt"
	"io"
	"os"

	"github.com/OpenRunic/config/options"
)

// configuration reader interface
type Reader interface {
	Configurator() string
	ConfigExtensions() []string
	Parse(opts *options.Options, fields []*Field) error
	Get(opts *options.Options, field *Field) (any, bool)
}

// base reader to extend with (optional)
type BaseReader struct{}

// file extension for related reader
func (r BaseReader) ConfigExtensions() []string {
	return nil
}

// placeholder parser func since not every reader needs to parse
func (r BaseReader) Parse(_ *options.Options, _ []*Field) error {
	return nil
}

// resolve file path for config file
func ResolvePaths(r Reader, opts *options.Options, path string) []string {
	if len(path) > 0 {
		return []string{path}
	}

	exts := r.ConfigExtensions()
	if len(exts) > 0 {
		paths := make([]string, len(exts))
		for i := range len(exts) {
			paths[i] = fmt.Sprintf("%s.%s", opts.Filename, exts[i])
		}
		return paths
	}

	return []string{opts.Filename}
}

// resolve byte data from file
func ResolveFileBytes(r Reader, paths []string, strict bool) ([]byte, error) {
	if len(paths) < 1 {
		return nil, nil
	}

	var err error
	var oFile *os.File

	for _, path := range paths {
		oFile, err = os.Open(path)
		if err != nil {
			oFile = nil
		} else {
			break
		}
	}

	if oFile != nil {
		defer oFile.Close()

		return io.ReadAll(oFile)
	}

	if strict {
		return nil, fmt.Errorf("[%s] load config failed: %s", r.Configurator(), paths[0])
	}

	return nil, nil
}

// resolve byte data from file or reader
func ResolveBytes(r Reader, reader io.Reader, paths []string, strict bool) ([]byte, error) {
	if reader != nil {
		return io.ReadAll(reader)
	}

	return ResolveFileBytes(r, paths, strict)
}
