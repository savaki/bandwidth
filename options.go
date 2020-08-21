package bandwidth

import "os"

type Options struct {
	accountID string
	username  string
	password  string
}

type Option func(*Options)

func WithCredentials(accountID, username, password string) Option {
	return func(o *Options) {
		o.accountID = accountID
		o.username = username
		o.password = password
	}
}

func buildOptions(opts ...Option) Options {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	if options.accountID == "" && options.username == "" && options.password == "" {
		options.accountID = os.Getenv("BANDWIDTH_ACCOUNT_ID")
		options.username = os.Getenv("BANDWIDTH_USERNAME")
		options.password = os.Getenv("BANDWIDTH_PASSWORD")
	}

	return options
}
