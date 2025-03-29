package config

import (
	"encoding/json"
	"errors"
	"reflect"
)

// Full list of default configuration readers available
var AVAILABLE_READERS = []string{
	READER_FLAG,
	READER_JSON,
	READER_ENV,
}

// Configuration management struct
type Config struct {
	Options *Options
	Values  map[string]any
	readers map[string]ConfigReader
	fields  []*Field
}

// Callback type for configuration setup
type WithConfigCallback func(*Config) error

// Add new field for config read
func (c *Config) Add(field *Field) {
	c.fields = append(c.fields, field)
}

// Add new configuration reader instance
func (c *Config) AddReader(reader ConfigReader) {
	c.readers[reader.Configurator()] = reader
}

// Read the stored configuration reader
func (c *Config) Get(name string) ConfigReader {
	r, ok := c.readers[name]
	if ok {
		return r
	}

	return nil
}

// Parse the configurations based of registered fields
func (c *Config) Parse(data any) error {

	// ensuring that pointer reference is provided
	tp := reflect.ValueOf(data)
	if data != nil && tp.Kind() != reflect.Ptr {
		return errors.New("invalid configuration target, expected writable pointer")
	}

	values := make(map[string]any)

	for idx, priority := range c.Options.Priority {
		reader, ok := c.readers[priority]

		if ok {
			err := reader.Parse(c.Options, c.fields)
			if err != nil {
				return err
			}

			for _, field := range c.fields {
				if idx == 0 && !field.Null {
					values[field.Name] = field.Value
				}

				val, ok := reader.Get(c.Options, field)
				if ok {
					values[field.Name] = val
				}
			}
		}
	}

	c.Values = values

	// decode parsed values to provided struct
	if data != nil {
		bytes, err := json.Marshal(values)
		if err != nil {
			return err
		}
		return json.Unmarshal(bytes, data)
	}

	return nil
}

// Create new instance of configurations with support for setup callbacks
func New(opts *Options, cbs ...WithConfigCallback) (*Config, error) {
	config := &Config{
		Options: opts,
		Values:  make(map[string]any),
		readers: make(map[string]ConfigReader),
		fields:  make([]*Field, 0),
	}

	for _, cb := range cbs {
		err := cb(config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

// Create instance of configurations using default settings
func Default(data any, cbs ...WithConfigCallback) (*Config, error) {
	final_cbs := []WithConfigCallback{
		Register(NewFlagReader()),
		Register(NewEnvReader()),
		Register(NewJSONReader()),
	}
	if len(cbs) > 0 {
		final_cbs = append(final_cbs, cbs...)
	}

	conf, err := New(
		NewOptions(
			WithPriority(AVAILABLE_READERS...),
		),
		final_cbs...,
	)
	if err != nil {
		return nil, err
	}

	return conf, conf.Parse(&data)
}

// Callback to register new reader instance
func Register(reader ConfigReader) WithConfigCallback {
	return func(c *Config) error {
		c.AddReader(reader)
		return nil
	}
}

// Callback shortcut to register new field
func Add(name string, value any, usage string) WithConfigCallback {
	return func(c *Config) error {
		field, err := NewField(name, value, usage)
		if err != nil {
			return err
		}

		c.Add(field)
		return nil
	}
}
