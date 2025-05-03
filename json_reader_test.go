package config_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/OpenRunic/config"
)

// testing json config reader
func TestJsonReader(t *testing.T) {
	var data map[string]any
	_, err := ParseTestReader(
		config.ReaderJSON,
		config.NewJSONReader(),
		&data,
	)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[ReaderTestKey]
	if !ok || v != ReaderTestValue {
		t.Errorf("expected '%s' but got '%v'", ReaderTestValue, v)
	}
}

// testing json config reader io
func TestJsonReaderIO(t *testing.T) {
	byteValue, err := json.Marshal(map[string]any{
		ReaderTestKey: ReaderTestValue,
	})
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]any
	_, err = ParseTestReader(
		config.ReaderJSON,
		config.NewJSONReaderIO(bytes.NewReader(byteValue)),
		&data,
	)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[ReaderTestKey]
	if !ok || v != ReaderTestValue {
		t.Errorf("expected '%s' but got '%v'", ReaderTestValue, v)
	}
}
