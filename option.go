package config

// Options struct for configurator
type Options struct {
	Strict   bool
	Prefix   string
	Filename string
	Priority []string
}

// Callback type for options setup
type WithOptionCallback func(*Options)

// Make new options instance using callbacks
func NewOptions(cbs ...WithOptionCallback) *Options {
	opts := &Options{
		Strict:   false,
		Prefix:   "app_",
		Filename: "config",
	}

	for _, cb := range cbs {
		cb(opts)
	}

	return opts
}

// Set prefix for keys
func WithPrefix(prefix string) WithOptionCallback {
	return func(opts *Options) {
		opts.Prefix = prefix
	}
}

// Set priority of config readers
func WithPriority(priorities ...string) WithOptionCallback {
	return func(opts *Options) {
		opts.Priority = priorities
	}
}

// Set readers to strict mode (experimental)
func UseStrict(s bool) WithOptionCallback {
	return func(opts *Options) {
		opts.Strict = s
	}
}

// Set filename for config file
func WithFilename(name string) WithOptionCallback {
	return func(opts *Options) {
		opts.Filename = name
	}
}
