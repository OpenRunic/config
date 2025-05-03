package config

import (
	"encoding/json"
	"io"
)

const ReaderJSON = "json"

// Configuration Reader for JSON Data/File
type JSONReader struct {
	BaseReader
	reader io.Reader
	path   string
	data   map[string]any
}

func (r *JSONReader) Configurator() string {
	return ReaderJSON
}

func (r *JSONReader) ConfigExtension() string {
	return "json"
}

func (r *JSONReader) Parse(opts *Options, _ []*Field) error {
	byteValue, err := ResolveBytes(r, opts, r.path, r.reader)
	if err != nil {
		return err
	}

	if byteValue == nil {
		return nil
	}

	var data map[string]any
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	r.data = FlattenMap(data, "")
	return nil
}

func (r *JSONReader) Get(_ *Options, field *Field) (any, bool) {
	if r.data == nil {
		return nil, false
	}

	value, exists := r.data[field.Name]
	if exists {
		if field.List {
			return CastSliceValue(value, field.Kind)
		}
		return CastValue(value, field.Kind)
	}

	return nil, false
}

// Create JSON configuration reader using file
func NewJSONReader() *JSONReader {
	return &JSONReader{}
}

// Create JSON configuration reader using io.Reader
func NewJSONReaderIO(reader io.Reader) *JSONReader {
	return &JSONReader{
		reader: reader,
	}
}
