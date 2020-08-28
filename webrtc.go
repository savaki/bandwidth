package bandwidth

type WebRTC struct {
	client *client
}

func NewWebRTC(opts ...Option) *WebRTC {
	const codebase = "https://api.webrtc.bandwidth.com/v1/accounts"

	options := buildOptions(opts...)
	client := newClient(codebase, options.accountID, options.username, options.password)
	return &WebRTC{
		client: client,
	}
}
