package bandwidth

type Voice struct {
	client *client
}

func NewVoice(opts ...Option) *Voice {
	options := buildOptions(opts...)
	client := newClient(options.accountID, options.username, options.password)
	return &Voice{
		client: client,
	}
}
