package json_test

import (
	"bytes"
	eJSON "encoding/json"
	"testing"

	"github.com/OpenRunic/config"
	"github.com/OpenRunic/config/json"
	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
)

// sample key and value for test
const ReaderTestKey = "test_field"
const ReaderTestValue = "tvalue"

// create test reader
func CreateTestReader(r reader.Reader, data any) (*config.Config, error) {
	return config.Parse(
		options.New(
			options.WithFilename("../samples/config"),
		),
		data,
		config.Register(r),
		config.Add(ReaderTestKey, "", "config test field"),
	)
}

// testing json config reader
func TestJsonReader(t *testing.T) {
	var data map[string]any
	_, err := CreateTestReader(json.New().Strict(), &data)
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
	byteValue, err := eJSON.Marshal(map[string]any{
		ReaderTestKey: ReaderTestValue,
	})
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]any
	_, err = CreateTestReader(
		json.NewIO(bytes.NewReader(byteValue)),
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
