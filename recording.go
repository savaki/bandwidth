package bandwidth

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Recording struct {
	AccountID        string              `json:"accountId,omitempty"`
	CallID           string              `json:"callId,omitempty"`
	ParentCallID     string              `json:"parentCallId,omitempty"`
	RecordingID      string              `json:"recordingId,omitempty"`
	To               string              `json:"to,omitempty"`
	From             string              `json:"from,omitempty"`
	TransferCallerID string              `json:"transferCallId,omitempty"`
	TransferTo       string              `json:"transferTo,omitempty"`
	Duration         string              `json:"duration,omitempty"`
	Direction        string              `json:"direction,omitempty"`
	Channels         int                 `json:"channels,omitempty"`
	StartTime        string              `json:"startTime,omitempty"`
	EndTime          string              `json:"endTime,omitempty"`
	FileFormat       string              `json:"fileFormat,omitempty"`
	Status           string              `json:"status,omitempty"`
	MediaURL         string              `json:"mediaUrl,omitempty"`
	Transcription    *TranscriptionEvent `json:"transcription,omitempty"`
}

type FindAllRecordingsInput struct {
	From         string `url:"from,omitempty"`         // 	Filter results by the from field.	No
	To           string `url:"to,omitempty"`           // 	Filter results by the to field.	No
	MinStartTime string `url:"minStartTime,omitempty"` // 	Filter results to recordings which have a startTime after or including minStartTime (in ISO8601 format).	No
	MaxStartTime string `url:"maxStartTime,omitempty"` // 	Filter results to recordings which have a startTime before maxStartTime (in ISO8601 format).	No
}

func (v *Voice) FindAllRecordings(ctx context.Context, input FindAllRecordingsInput) (recordings []Recording, err error) {
	form, err := query.Values(input)
	if err != nil {
		return nil, fmt.Errorf("unable to find recordings: %w", err)
	}

	path := "/recordings?" + form.Encode()
	if err := v.client.Get(ctx, path, &recordings); err != nil {
		return nil, fmt.Errorf("failed to fetch recordings: %w", err)
	}

	return recordings, nil
}

func (v *Voice) FindAllCallRecordings(ctx context.Context) ([]Recording, error) {
	return nil, nil
}

func (v *Voice) FindRecording(ctx context.Context) ([]Recording, error) {
	return nil, nil
}

func (v *Voice) DeleteRecording(ctx context.Context) (err error) {
	return nil
}
