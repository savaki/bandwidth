package bandwidth

import (
	"context"
	"fmt"
	"path/filepath"
)

// Call - https://dev.bandwidth.com/voice/methods/calls/postCalls.html
type Call struct {
	AccountId            string  `json:"accountId,omitempty"`            // AccountId
	AnswerFallbackMethod string  `json:"answerFallbackMethod,omitempty"` // AnswerFallbackMethod - (optional) The HTTP method to use to deliver the answer callback to answerFallbackUrl. GET or POST. Default value is POST.
	AnswerFallbackUrl    string  `json:"answerFallbackUrl,omitempty"`    // AnswerFallbackUrl - (optional) A fallback url which, if provided, will be used to retry the answer callback delivery in case answerUrl fails to respond
	AnswerMethod         string  `json:"answerMethod,omitempty"`         // AnswerMethod - (optional) The HTTP method to use for the request to answerUrl. GET or POST. Default value is POST.
	AnswerTime           string  `json:"answerTime,omitempty"`           // AnswerTime
	AnswerURL            string  `json:"answerUrl,omitempty"`            // AnswerURL - The full URL to send the Answer event to when the called party answers. This endpoint should return the first BXML document to be executed in the call.
	ApplicationId        string  `json:"applicationId,omitempty"`        // ApplicationId
	CallbackTimeout      float32 `json:"callbackTimeout,omitempty"`      // CallbackTimeout - (optional) This is the timeout (in seconds) to use when delivering callbacks for the call. Can be any numeric value (including decimals) between 1 and 25. Default: 15
	CallId               string  `json:"callId,omitempty"`               // CallId
	CallTimeout          float32 `json:"callTimeout,omitempty"`          // CallTimeout - (optional) This is the timeout (in seconds) for the callee to answer the call. Can be any numeric value (including decimals) between 1 and 300. Default: 30
	CallUrl              string  `json:"callUrl,omitempty"`              // CallUrl
	DisconnectCause      string  `json:"disconnectCause,omitempty"`      // DisconnectCause
	DisconnectMethod     string  `json:"disconnectMethod,omitempty"`     // DisconnectMethod - (optional) The HTTP method to use for the request to disconnectUrl. GET or POST. Default value is POST.
	DisconnectURL        string  `json:"disconnectUrl,omitempty"`        // DisconnectURL - The full URL to send the Disconnect event to when the called party disconnects. This endpoint should return the first BXML document to be executed in the call.
	EndTime              string  `json:"endTime,omitempty"`              // EndTime
	FallbackPassword     string  `json:"fallbackPassword,omitempty"`     // FallbackPassword - (optional) The password to send in the HTTP request to answerFallbackUrl
	FallbackUsername     string  `json:"fallbackUsername,omitempty"`     // FallbackUsername - (optional) The username to send in the HTTP request to answerFallbackUrl
	From                 string  `json:"from,omitempty"`                 // From - A Bandwidth phone number on your account the call should come from (must be in E.164 format, like +15555551212).
	Password             string  `json:"password,omitempty"`             // Password - (optional) The password to send in the HTTP request to answerUrl and disconnectUrl.
	StartTime            string  `json:"startTime,omitempty"`            // StartTime
	Tag                  string  `json:"tag,omitempty"`                  // Tag - (optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	To                   string  `json:"to,omitempty"`                   // To - The number to call (must be an E.164 formatted number, like +15555551212
	Username             string  `json:"username,omitempty"`             // Username - (optional) The username to send in the HTTP request to answerUrl and disconnectUrl.
}

// https://dev.bandwidth.com/voice/methods/calls/postCalls.html
type CreateCallInput struct {
	AnswerFallbackMethod string `json:"answerFallbackMethod,omitempty"` // AnswerFallbackMethod - (optional) The HTTP method to use to deliver the answer callback to answerFallbackUrl. GET or POST. Default value is POST.
	AnswerFallbackUrl    string `json:"answerFallbackUrl,omitempty"`    // AnswerFallbackUrl - (optional) A fallback url which, if provided, will be used to retry the answer callback delivery in case answerUrl fails to respond
	AnswerMethod         string `json:"answerMethod,omitempty"`         // AnswerMethod - (optional) The HTTP method to use for the request to answerUrl. GET or POST. Default value is POST.
	AnswerURL            string `json:"answerUrl,omitempty"`            // AnswerURL - The full URL to send the Answer event to when the called party answers. This endpoint should return the first BXML document to be executed in the call.
	ApplicationId        string `json:"applicationId,omitempty"`        // ApplicationId	The id of the application to associate this call with, for billing purposes.
	CallbackTimeout      int    `json:"callbackTimeout,omitempty"`      // CallbackTimeout - (optional) This is the timeout (in seconds) to use when delivering callbacks for the call. Can be any numeric value (including decimals) between 1 and 25. Default: 15
	CallTimeout          int    `json:"callTimeout,omitempty"`          // CallTimeout - (optional) This is the timeout (in seconds) for the callee to answer the call. Can be any numeric value (including decimals) between 1 and 300. Default: 30
	DisconnectMethod     string `json:"disconnectMethod,omitempty"`     // DisconnectMethod - (optional) The HTTP method to use for the request to disconnectUrl. GET or POST. Default value is POST.
	DisconnectURL        string `json:"disconnectUrl,omitempty"`        // DisconnectURL - The full URL to send the Disconnect event to when the called party disconnects. This endpoint should return the first BXML document to be executed in the call.
	FallbackPassword     string `json:"fallbackPassword,omitempty"`     // FallbackPassword - (optional) The password to send in the HTTP request to answerFallbackUrl
	FallbackUsername     string `json:"fallbackUsername,omitempty"`     // FallbackUsername - (optional) The username to send in the HTTP request to answerFallbackUrl
	From                 string `json:"from,omitempty"`                 // From - A Bandwidth phone number on your account the call should come from (must be in E.164 format, like +15555551212).
	Password             string `json:"password,omitempty"`             // Password - (optional) The password to send in the HTTP request to answerUrl and disconnectUrl.
	Tag                  string `json:"tag,omitempty"`                  // Tag - (optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	To                   string `json:"to,omitempty"`                   // To - The number to call (must be an E.164 formatted number, like +15555551212
	Username             string `json:"username,omitempty"`             // Username - (optional) The username to send in the HTTP request to answerUrl and disconnectUrl.
}

// CreateCall - Creates a new outbound phone call.
// https://dev.bandwidth.com/voice/methods/calls/postCalls.html
func (v *Voice) CreateCall(ctx context.Context, input CreateCallInput) (call Call, err error) {
	if err := v.client.Post(ctx, "/calls", input, &call); err != nil {
		return Call{}, fmt.Errorf("failed to create call: %w", err)
	}
	return call, nil
}

func (v *Voice) FindCall(ctx context.Context, callId string) (call Call, err error) {
	path := filepath.Join("/calls", callId)
	if err := v.client.Get(ctx, path, &call); err != nil {
		return Call{}, fmt.Errorf("unable to find call, %v: %w", callId, err)
	}
	return Call{}, nil
}

type PauseRecordingInput struct {
	CallId string `json:"-"`               // CallId
	State  string `json:"state,omitempty"` // State - The recording state. Possible values: paused to pause an active recording OR recording to resume a paused recording
}

// PauseRecording - Pause or resume a recording on an active phone call.
// https://dev.bandwidth.com/voice/methods/recordings/putCallsCallIdRecording.html
func (v *Voice) PauseRecording(ctx context.Context, input PauseRecordingInput) error {
	path := filepath.Join("/calls", input.CallId, "recording")
	if err := v.client.Put(ctx, path, input, nil); err != nil {
		return fmt.Errorf("failed to update call, %v: %w", input.CallId, err)
	}
	return nil
}

// UpdateCallInput - https://dev.bandwidth.com/voice/methods/calls/postCallsCallId.html
type UpdateCallInput struct {
	CallId                 string `json:"-"`                                // CallId
	FallbackPassword       string `json:"fallbackPassword,omitempty"`       // FallbackPassword - (optional) The password to send in the HTTP request to redirectFallbackUrl
	FallbackUsername       string `json:"fallbackUsername,omitempty"`       // FallbackUsername - (optional) The username to send in the HTTP request to redirectFallbackUrl
	Password               string `json:"password,omitempty"`               // Password - (optional) The password to send in the HTTP request to answerUrl and disconnectUrl.
	RedirectFallbackURL    string `json:"redirectFallbackUrl,omitempty"`    // RedirectFallbackURL - (optional) A fallback url which, if provided, will be used to retry the redirect callback delivery in case redirectUrl fails to respond
	RedirectFallbackMethod string `json:"redirectFallbackMethod,omitempty"` // RedirectFallbackMethod - (optional) The HTTP method to use to deliver the redirect callback to redirectFallbackUrl. GET or POST. Default value is POST.
	RedirectMethod         string `json:"redirectMethod,omitempty"`         // RedirectMethod - (optional) The HTTP method to use for the request to redirectUrl. GET or POST. Default value is POST.
	RedirectURL            string `json:"redirectUrl,omitempty"`            // RedirectURL - The full URL to send the Redirect event to when the called party redirects. This endpoint should return the first BXML document to be executed in the call.
	State                  string `json:"state,omitempty"`                  // State - (optional) The call state. Possible values: active to redirect the call (default) OR completed to hangup the call
	Username               string `json:"username,omitempty"`               // Username - (optional) The username to send in the HTTP request to answerUrl and disconnectUrl.
	Tag                    string `json:"tag,omitempty"`                    // Tag - (optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
}

// UpdateCall - Update properties of an active phone call.
// https://dev.bandwidth.com/voice/methods/calls/postCallsCallId.html
func (v *Voice) UpdateCall(ctx context.Context, input UpdateCallInput) error {
	path := filepath.Join("/calls", input.CallId)
	if err := v.client.Post(ctx, path, input, nil); err != nil {
		return fmt.Errorf("failed to update call, %v: %w", input.CallId, err)
	}
	return nil
}
