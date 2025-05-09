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
const ReaderTestKey2 = "test_int"
const ReaderTestValue = "tvalue"
const ReaderTestValue2 = 15000

type ReaderTestData struct {
	Field  string `json:"test_field"`
	Number int64  `json:"test_int"`
}

// create test reader
func CreateTestReader(r reader.Reader, data any) (*config.Config, error) {
	return config.Parse(
		options.New(
			options.WithFilename("../samples/config"),
		),
		data,
		config.Register(r),
		config.Add(ReaderTestKey, "", "config test field"),
		config.Add[int64](ReaderTestKey2, 15, "config test int field"),
	)
}

// testing yaml config reader
func TestYamlReader(t *testing.T) {
	var data ReaderTestData
	_, err := CreateTestReader(yaml.New().Strict(), &data)
	if err != nil {
		t.Fatal(err)
	}

	if data.Field != ReaderTestValue {
		t.Errorf("expected '%s' but got '%v'", ReaderTestValue, data.Field)
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
