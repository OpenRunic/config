package env

import (
	"os"
	"strings"

	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
)

const ReaderEnv = "env"

// Configuration Reader for Environment Variables
type Reader struct {
	reader.BaseReader
}

func (r Reader) Configurator() string {
	return ReaderEnv
}

func (r Reader) Get(opts *options.Options, field *reader.Field) (any, bool) {
	for _, al := range field.Alias {
		key := strings.ToUpper(al)
		if len(opts.Prefix) > 0 {
			key = strings.ToUpper(opts.Prefix) + key
		}

		value, exists := os.LookupEnv(key)
		if exists {
			if field.List {
				return reader.CastSliceValue(value, field.Kind)
			}

			return reader.CastValue(value, field.Kind)
		}
	}

	return nil, false
}

func New() *Reader {
	return &Reader{}
}
