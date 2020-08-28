package bandwidth

import (
	"context"
	"fmt"
)

type Participant struct {
	ID                 string        `json:"id,omitempty"`
	CallbackUrl        string        `json:"callbackUrl,omitempty"`
	PublishPermissions []string      `json:"publishPermissions,omitempty"` // 	Defines if this participant can publish audio or video
	Sessions           []string      `json:"sessions,omitempty"`
	Subscriptions      Subscriptions `json:"subscriptions,omitempty"` // 	Subscription information for this participant
}

type Subscriptions struct {
	SessionId    string        `json:"sessionId,omitempty"`
	Participants []Participant `json:"participants,omitempty"`
}

type SubscriptionParticipant struct {
	ParticipantId string `json:"participantId,omitempty"`
}

type CreateParticipantInput struct {
	CallbackUrl        string        `json:"callbackUrl,omitempty"`        // 	Full callback url to use for notifications about this participant
	PublishPermissions []string      `json:"publishPermissions,omitempty"` // 	Defines if this participant can publish audio or video
	Subscriptions      Subscriptions `json:"subscriptions,omitempty"`      // 	Subscription information for this participant
	Tag                string        `json:"tag,omitempty"`                // 	User defined tag to associate with the participant
}

type CreateParticipantOutput struct {
	Participant Participant
	Token       string
}

// CreateParticipant - Participants are idempotent, so relevant parameters must be set in this function if desired
// https://dev.bandwidth.com/webrtc/methods/participants/createParticipant.html
func (w *WebRTC) CreateParticipant(ctx context.Context, input CreateParticipantInput) (output CreateParticipantOutput, err error) {
	path := "/participants"
	if err := w.client.Post(ctx, path, input, &output); err != nil {
		return CreateParticipantOutput{}, fmt.Errorf("failed to create webrtc participant: %w", err)
	}
	return output, nil
}
