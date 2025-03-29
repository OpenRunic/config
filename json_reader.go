package config

import (
	"encoding/json"
	"io"
)

const READER_JSON = "json"

// Configuration Reader for JSON Data/File
type JSONReader struct {
	BaseReader
	reader io.Reader
	path   string
	data   map[string]any
}

func (r *JSONReader) Configurator() string {
	return READER_JSON
}

func (r *JSONReader) ConfigExtension() string {
	return "json"
}

func (r *JSONReader) Parse(opts *Options, fields []*Field) error {
	byteValue, err := r.ResolveBytes(opts, r.path, r.reader)
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

func (r *JSONReader) Get(opts *Options, field *Field) (any, bool) {
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
