package config_test

import "github.com/OpenRunic/config"

// sample key and value for test
const READER_TEST_KEY = "test_field"
const READER_TEST_VALUE = "tvalue"

// configure test reader
func GetTestReader(name string, reader config.ConfigReader) (*config.Config, error) {
	return config.New(
		config.NewOptions(
			config.WithPriority(name),
		),
		config.Register(reader),
		config.Add(READER_TEST_KEY, "", ""),
	)
}

// create test reader and parse data from it
func ParseTestReader(name string, reader config.ConfigReader, data any) (*config.Config, error) {
	cnf, err := GetTestReader(name, reader)
	if err != nil {
		return nil, err
	}

	err = cnf.Parse(data)
	if err != nil {
		return cnf, err
	}
	return cnf, nil
}
