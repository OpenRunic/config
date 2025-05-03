package config_test

import "github.com/OpenRunic/config"

// sample key and value for test
const ReaderTestKey = "test_field"
const ReaderTestValue = "tvalue"

// configure test reader
func GetTestReader(name string, reader config.Reader) (*config.Config, error) {
	return config.New(
		config.NewOptions(
			config.WithPriority(name),
			config.WithFilename("samples/config"),
		),
		config.Register(reader),
		config.Add(ReaderTestKey, "", ""),
	)
}

// create test reader and parse data from it
func ParseTestReader(name string, reader config.Reader, data any) (*config.Config, error) {
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
