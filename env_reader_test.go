package config_test

import (
	"os"
	"testing"

	"github.com/OpenRunic/config"
)

// testing environment config reader
func TestEnvReader(t *testing.T) {
	os.Setenv("APP_TEST_NAME", READER_TEST_VALUE)

	var data map[string]any
	_, err := ParseTestReader(
		config.READER_ENV,
		config.NewEnvReader(),
		&data,
	)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := data[READER_TEST_KEY]
	if !ok || v != READER_TEST_VALUE {
		t.Errorf("expected '%s' but got '%s'", READER_TEST_VALUE, v)
	}
}
