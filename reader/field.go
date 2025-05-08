package reader

import (
	"fmt"
	"reflect"
	"strings"
)

// Field definition struct for loading config
type Field struct {
	Name  string
	Usage string
	Value any
	Kind  reflect.Kind
	Alias []string
	List  bool
	Null  bool
}

// Convert value to string
func (f *Field) ValueAsString() string {
	if !f.Null && f.Kind == reflect.String {
		if f.List {
			return strings.Join(f.Value.([]string), ",")
		}

		return f.Value.(string)
	}

	return ""
}

// Convert value to int
func (f *Field) ValueAsInt() int {
	if !f.Null && f.Kind == reflect.Int {
		return f.Value.(int)
	}

	return 0
}

// Convert value to bool
func (f *Field) ValueAsBool() bool {
	if !f.Null && f.Kind == reflect.Bool {
		return f.Value.(bool)
	}

	return false
}

// Add field alias (e.g. env keys)
func (f *Field) AddAlias(a string) *Field {
	f.Alias = append(f.Alias, a)
	return f
}

// Create new instance of field and resolve configs
func NewField[T any](name string, value T, usage string) (*Field, error) {
	isNull := false
	isList := false

	var kind reflect.Kind
	tp := reflect.ValueOf(value)
	if tp.Kind() == reflect.Ptr {
		if !tp.Elem().IsValid() {
			isNull = true
		}
		kind = tp.Type().Elem().Kind()
	} else {
		kind = tp.Kind()
	}

	if kind == reflect.Slice {
		isList = true
		isNull = tp.IsNil()
		kind = tp.Type().Elem().Kind()

		if kind != reflect.String {
			return nil, fmt.Errorf("slice of [%s] isn't supported field value", kind)
		}
	}

	return &Field{
		Name:  name,
		Usage: usage,
		Value: value,
		Kind:  kind,
		List:  isList,
		Null:  isNull,
		Alias: []string{ToSnakeCase(name)},
	}, nil
}
