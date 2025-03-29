package config_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/OpenRunic/config"
)

// testing json config reader
func TestJsonReader(t *testing.T) {
	byteValue, err := json.Marshal(map[string]any{
		READER_TEST_KEY: READER_TEST_VALUE,
	})
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]any
	_, err = ParseTestReader(
		config.READER_JSON,
		config.NewJSONReaderIO(bytes.NewReader(byteValue)),
		&data,
	)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[READER_TEST_KEY]
	if !ok || v != READER_TEST_VALUE {
		t.Errorf("expected '%s' but got '%v'", READER_TEST_VALUE, v)
	}
}
