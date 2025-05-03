package config

import (
	"os"
	"strings"
)

const ReaderEnv = "env"

// Configuration Reader for Environment Variables
type EnvReader struct {
	BaseReader
}

func (r EnvReader) Configurator() string {
	return ReaderEnv
}

func (r EnvReader) Get(opts *Options, field *Field) (any, bool) {
	for _, al := range field.Alias {
		key := strings.ToUpper(al)
		if len(opts.Prefix) > 0 {
			key = strings.ToUpper(opts.Prefix) + key
		}

		value, exists := os.LookupEnv(key)
		if exists {
			if field.List {
				return CastSliceValue(value, field.Kind)
			}

			return CastValue(value, field.Kind)
		}
	}

	return nil, false
}

func NewEnvReader() *EnvReader {
	return &EnvReader{}
}
