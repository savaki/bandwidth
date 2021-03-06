package bandwidth

type Voice struct {
	client *client
}

func NewVoice(opts ...Option) *Voice {
	const codebase = "https://voice.bandwidth.com/api/v2/accounts/"

	options := buildOptions(opts...)
	client := newClient(codebase, options.accountID, options.username, options.password)
	return &Voice{
		client: client,
	}
}
