package yaml_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/OpenRunic/config"
	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/reader"
	"github.com/OpenRunic/config/yaml"
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

// testing yaml config reader
func TestYamlReader(t *testing.T) {
	var data map[string]any
	_, err := CreateTestReader(yaml.New().Strict(), &data)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[ReaderTestKey]
	if !ok || v != ReaderTestValue {
		t.Errorf("expected '%s' but got '%v'", ReaderTestValue, v)
	}
}

// testing yaml config reader io
func TestYamlReaderIO(t *testing.T) {
	ymlData := fmt.Sprintf(`
%s: %s
`, ReaderTestKey, ReaderTestValue)

	var data map[string]any
	_, err := CreateTestReader(
		yaml.NewIO(bytes.NewReader([]byte(ymlData))),
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
