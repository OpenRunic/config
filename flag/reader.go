package config

import (
	"flag"
	"reflect"

	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
)

const ReaderFlag = "flag"

// Configuration Reader for Flags(Commandline) inputs
type Reader struct {
	reader.BaseReader
	flags map[string]any
}

func (r *Reader) Configurator() string {
	return ReaderFlag
}

func (r *Reader) Parse(_ *options.Options, fields []*reader.Field) error {
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

func (r *Reader) Get(_ *options.Options, field *reader.Field) (any, bool) {
	val, ok := r.flags[field.Name]

	if ok && val != nil {
		if field.List {
			return reader.CastSliceValue(*val.(*string), field.Kind)
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

func New() *Reader {
	return &Reader{}
}
