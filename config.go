package config

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
)

// Configuration management struct
type Config struct {
	Options *options.Options
	Values  map[string]any
	readers []reader.Reader
	fields  []*reader.Field
}

// Callback type for configuration setup
type WithConfigCallback func(*Config) error

// Add new field for config read
func (c *Config) Add(field *reader.Field) {
	c.fields = append(c.fields, field)
}

// Add new configuration reader instance
func (c *Config) AddReader(r reader.Reader) {
	c.readers = append(c.readers, r)
}

// Parse the configurations based of registered fields
func (c *Config) Parse(data any) error {

	// ensuring that pointer reference is provided
	tp := reflect.ValueOf(data)
	if data != nil && tp.Kind() != reflect.Ptr {
		return errors.New("invalid configuration target, expected writable pointer")
	}

	values := make(map[string]any)

	for idx, r := range c.readers {
		err := r.Parse(c.Options, c.fields)
		if err != nil {
			return err
		}

		for _, field := range c.fields {
			if idx == 0 && !field.Null {
				values[field.Name] = field.Value
			}

			val, ok := r.Get(c.Options, field)
			if ok {
				values[field.Name] = val
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
func New(opts *options.Options, cbs ...WithConfigCallback) (*Config, error) {
	cnf := &Config{
		Options: opts,
		Values:  make(map[string]any),
		readers: make([]reader.Reader, 0),
		fields:  make([]*reader.Field, 0),
	}

	for _, cb := range cbs {
		err := cb(cnf)
		if err != nil {
			return nil, err
		}
	}

	return cnf, nil
}

// Create instance and parse configs
func Parse(opts *options.Options, data any, cbs ...WithConfigCallback) (*Config, error) {
	cnf, err := New(opts, cbs...)
	if err != nil {
		return nil, err
	}

	err = cnf.Parse(data)
	if err != nil {
		return nil, err
	}

	return cnf, nil
}

// Callback to register new reader instance
func Register(r reader.Reader) WithConfigCallback {
	return func(c *Config) error {
		c.AddReader(r)
		return nil
	}
}

// Callback shortcut to register new field
func Add[T any](name string, value T, usage string) WithConfigCallback {
	return func(c *Config) error {
		field, err := reader.NewField[T](name, value, usage)
		if err != nil {
			return err
		}

		c.Add(field)
		return nil
	}
}
