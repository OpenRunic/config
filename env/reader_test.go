package env_test

import (
	"os"
	"testing"

	"github.com/OpenRunic/config"
	"github.com/OpenRunic/config/env"
	"github.com/OpenRunic/config/options"
)

// sample key and value for test
const ReaderTestKey = "test_field"
const ReaderTestValue = "tvalue"

// create test reader
func CreateTestReader(data any) (*config.Config, error) {
	return config.Parse(
		options.New(),
		data,
		config.Register(env.New()),
		config.Add(ReaderTestKey, "", "config test field"),
	)
}

// testing environment config reader
func TestEnvReader(t *testing.T) {
	os.Setenv("APP_TEST_FIELD", ReaderTestValue)

	var data map[string]any
	_, err := CreateTestReader(&data)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[ReaderTestKey]
	if !ok || v != ReaderTestValue {
		t.Errorf("expected '%s' but got '%s'", ReaderTestValue, v)
	}
}
