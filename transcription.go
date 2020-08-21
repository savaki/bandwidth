package bandwidth

import (
	"context"
	"fmt"
	"path/filepath"
)

type Transcript struct {
	Text       string
	Confidence float64
}

type RequestTranscriptsInput struct {
	CallId         string `json:"-"`
	RecordingId    string `json:"-"`
	CallbackUrl    string `json:"callbackUrl,omitempty"`    // 	The URL to send the TranscriptionAvailable event to. You should not include sensitive or personally-identifiable information in the callbackUrl field! Always use the proper username and password fields for authorization.	No
	CallbackMethod string `json:"callbackMethod,omitempty"` // 	The HTTP method to use for the request to callbackUrl. GET or POST. Default value is POST.	No
	Username       string `json:"username,omitempty"`       // 	The username to send in the HTTP request to callbackUrl.	No
	Password       string `json:"password,omitempty"`       // 	The password to send in the HTTP request to callbackUrl.	No
	Tag            string `json:"tag,omitempty"`            // 	A custom string that will be sent with this callbacks.	No
}

// RequestTranscripts - Generate the transcription for a specific recording. Transcription can succeed only for recordings of length greater than 500 milliseconds and less than 4 hours.
// https://dev.bandwidth.com/voice/methods/recordings/postCallsCallIdRecordingsRecordingIdTranscription.html
func (v *Voice) RequestTranscripts(ctx context.Context, input RequestTranscriptsInput) error {
	path := filepath.Join("/calls", input.CallId, "recordings", input.RecordingId, "transcription")
	if err := v.client.Post(ctx, path, input, nil); err != nil {
		return fmt.Errorf("failed to request transcription for call, %v, and recording, %v: %w", input.CallId, input.RecordingId, err)
	}
	return nil
}

type DownloadTranscriptsInput struct {
	CallId      string `json:"-"`
	RecordingId string `json:"-"`
}

// DownloadTranscripts - Retrieve the specified recording's transcription file. ⚠️ Be sure to not expose your API Credentials to end-users
//
// If the transcribed recording was multi-channel, then there will be 2 transcripts.
// The caller/called party transcript will be the first item while <PlayAudio> and <SpeakSentence> transcript will be the second item.
// During a <Transfer> the A-leg transcript will be the first item while the B-leg transcript will be the second item.
//
// https://dev.bandwidth.com/voice/methods/recordings/getCallsCallIdRecordingsRecordingIdTranscription.html
func (v *Voice) DownloadTranscripts(ctx context.Context, input DownloadTranscriptsInput) ([]Transcript, error) {
	path := filepath.Join("/calls", input.CallId, "recordings", input.RecordingId, "transcription")

	var content struct{ Transcripts []Transcript }
	if err := v.client.Get(ctx, path, &content); err != nil {
		return nil, fmt.Errorf("unable to download")
	}

	return content.Transcripts, nil
}

type DeleteTranscriptsInput struct {
	CallId      string `json:"-"`
	RecordingId string `json:"-"`
}

// DeleteTranscripts - Delete the specified transcription.
// https://dev.bandwidth.com/voice/methods/recordings/deleteCallsCallIdRecordingsRecordingIdTranscription.html
func (v *Voice) DeleteTranscripts(ctx context.Context, input DeleteTranscriptsInput) (err error) {
	path := filepath.Join("/calls", input.CallId, "recordings", input.RecordingId, "transcription")
	if err := v.client.Delete(ctx, path, nil, nil); err != nil {
		return fmt.Errorf("unable to download")
	}
	return nil
}
