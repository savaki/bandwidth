package bandwidth

import (
	"context"
	"fmt"
	"path/filepath"
)

type Session struct {
	ID           string   `json:"id,omitempty"`
	Participants []string `json:"participants,omitempty"`
}

type AddParticipantInput struct {
	SessionId     string `json:"sessionId,omitempty"`
	ParticipantId string `json:"-"`
}

// AddParticipant - Subscriptions can optionally be provided as part of this call
// https://dev.bandwidth.com/webrtc/methods/sessions/addParticipantToSession.html
func (w *WebRTC) AddParticipant(ctx context.Context, input AddParticipantInput) error {
	path := filepath.Join("/sessions", input.SessionId, "participants", input.ParticipantId)
	if err := w.client.Put(ctx, path, input, nil); err != nil {
		return fmt.Errorf("unable to add participant, %v, to session, %v: %w", input.ParticipantId, input.SessionId, err)
	}
	return nil
}

type CreateSessionInput struct {
	Tag string `json:"tag,omitempty"` // 	User defined tag to associate with the session
}

// CreateSession - Sessions are idempotent, so relevant parameters must be set in this function if desired
// https://dev.bandwidth.com/webrtc/methods/sessions/createSession.html
func (w *WebRTC) CreateSession(ctx context.Context, input CreateSessionInput) (s Session, err error) {
	if err := w.client.Post(ctx, "/sessions", input, &s); err != nil {
		return Session{}, fmt.Errorf("unable to create session: %w", err)
	}
	return s, nil
}
