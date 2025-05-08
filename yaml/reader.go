package yaml

import (
	"io"

	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
	"gopkg.in/yaml.v3"
)

const ReaderYAML = "yaml"

// Configuration Reader for YAML Data/File
type Reader struct {
	reader.BaseReader
	ir     io.Reader
	strict bool
	path   string
	data   map[string]any
}

func (r *Reader) Configurator() string {
	return ReaderYAML
}

func (r *Reader) ConfigExtensions() []string {
	return []string{"yml", "yaml"}
}

func (r *Reader) Strict() *Reader {
	r.strict = true
	return r
}

func (r *Reader) Parse(opts *options.Options, _ []*reader.Field) error {
	fPaths := reader.ResolvePaths(r, opts, r.path)
	byteValue, err := reader.ResolveBytes(r, r.ir, fPaths, opts.Strict || r.strict)
	if err != nil {
		return err
	}

	if byteValue == nil {
		return nil
	}

	var data map[string]any
	err = yaml.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	r.data = reader.FlattenMap(data, "")
	return nil
}

func (r *Reader) Get(_ *options.Options, field *reader.Field) (any, bool) {
	if r.data == nil {
		return nil, false
	}

	value, exists := r.data[field.Name]
	if exists {
		if field.List {
			return reader.CastSliceValue(value, field.Kind)
		}
		return reader.CastValue(value, field.Kind)
	}

	return nil, false
}

// Create JSON configuration reader using file
func New() *Reader {
	return &Reader{}
}

// Create JSON configuration reader using io.Reader
func NewIO(reader io.Reader) *Reader {
	return &Reader{
		ir: reader,
	}
}
