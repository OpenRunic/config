package config

import (
	"flag"
	"reflect"
)

const READER_FLAG = "flag"

// Configuration Reader for Flags(Commandline) inputs
type FlagReader struct {
	BaseReader
	flags map[string]any
}

func (r *FlagReader) Configurator() string {
	return READER_FLAG
}

func (r *FlagReader) Parse(_ *Options, fields []*Field) error {
	if !flag.Parsed() {
		r.flags = make(map[string]any)

		for _, field := range fields {
			if field.List {
				if field.Kind == reflect.String {
					r.flags[field.Name] = flag.String(field.Name, field.ValueAsString(), field.Usage)
				}
			} else {
				switch field.Kind {
				case reflect.String:
					r.flags[field.Name] = flag.String(field.Name, field.ValueAsString(), field.Usage)
				case reflect.Int:
					r.flags[field.Name] = flag.Int(field.Name, field.ValueAsInt(), field.Usage)
				case reflect.Bool:
					r.flags[field.Name] = flag.Bool(field.Name, field.ValueAsBool(), field.Usage)
				}
			}
		}

		flag.Parse()
	}

	return nil
}

func (r *FlagReader) Get(opts *Options, field *Field) (any, bool) {
	val, ok := r.flags[field.Name]

	if ok && val != nil {
		if field.List {
			return CastSliceValue(*val.(*string), field.Kind)
		}

		switch field.Kind {
		case reflect.String:
			v := *val.(*string)
			if (field.Null && len(v) < 1) || field.Value == v {
				return "", false
			}
			return v, true
		case reflect.Int:
			v := *val.(*int)
			if (field.Null && v == 0) || field.Value == v {
				return 0, false
			}
			return v, true
		case reflect.Bool:
			v := *val.(*bool)
			if (field.Null && !v) || field.Value == v {
				return false, false
			}
			return v, true
		case reflect.Slice:
		}
	}

	return nil, false
}

func NewFlagReader() *FlagReader {
	return &FlagReader{}
}
