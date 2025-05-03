package config_test

import (
	"os"
	"testing"

	"github.com/OpenRunic/config"
)

// testing environment config reader
func TestEnvReader(t *testing.T) {
	os.Setenv("APP_TEST_FIELD", ReaderTestValue)

	var data map[string]any
	_, err := ParseTestReader(
		config.ReaderEnv,
		config.NewEnvReader(),
		&data,
	)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[ReaderTestKey]
	if !ok || v != ReaderTestValue {
		t.Errorf("expected '%s' but got '%s'", ReaderTestValue, v)
	}
}
