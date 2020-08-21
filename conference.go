package bandwidth

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
)

type Conference struct {
	ID                    string `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	CreatedTime           string `json:"createdTime,omitempty"`
	CompletedTime         string `json:"completedTime,omitempty"`
	ConferenceEventUrl    string `json:"conferenceEventUrl,omitempty"`
	ConferenceEventMethod string `json:"conferenceEventMethod,omitempty"`
	Tag                   string `json:"tag,omitempty"`
}

// ConferenceMember information
// https://dev.bandwidth.com/voice/methods/conferences/getConferenceMember.html
type ConferenceMember struct {
	CallId         string `json:"callId,omitempty"`         // 	The conference member id.
	ConferenceId   string `json:"conferenceId,omitempty"`   // 	The conference id from the conference this member belongs to.
	MemberUrl      string `json:"memberUrl,omitempty"`      // 	The URL to to interact with this member.
	Mute           string `json:"mute,omitempty"`           // 	If true, the member is on mute and cannot speak in the conference.
	Hold           string `json:"hold,omitempty"`           // 	If true, the member is on hold and cannot speak or hear anything in the conference.
	CallIdsToCoach string `json:"callIdsToCoach,omitempty"` // 	The list of call ids to coach.
}

type FindAllConferencesInput struct {
	PageSize       string `json:"pageSize,omitempty"`       // 	Specifies the max number of conferences that will be returned. Range: integer values between 1 - 1000. Default value is 1000.
	Name           string `json:"name,omitempty"`           // 	Filter results by the name field.
	MinCreatedTime string `json:"minCreatedTime,omitempty"` // 	Filter results to conferences which have a createdTime after or including minCreatedTime (in ISO8601 format).
	MaxCreatedTime string `json:"maxCreatedTime,omitempty"` // 	Filter results to conferences which have a createdTime before or including maxCreatedTime (in ISO8601 format).
}

// FindAllConferences - Returns a max of 1000 conferences, sorted by createdTime from oldest to newest.
// https://dev.bandwidth.com/voice/methods/conferences/getConferences.html
func (v *Voice) FindAllConferences(ctx context.Context, input FindAllConferencesInput) (conferences []Conference, err error) {
	form := url.Values{}
	path := "/conferences?" + form.Encode()
	if err := v.client.Get(ctx, path, &conferences); err != nil {
		return nil, fmt.Errorf("failed to retrieve conferences: %w", err)
	}
	return conferences, nil
}

// FindConference - Retrieve the current state of a specific conference.
// https://dev.bandwidth.com/voice/methods/conferences/getConferencesConferenceId.html
func (v *Voice) FindConference(ctx context.Context, conferenceId string) (conference Conference, err error) {
	path := filepath.Join("/conferences", conferenceId)
	if err := v.client.Get(ctx, path, &conference); err != nil {
		return Conference{}, fmt.Errorf("failed to get conference, %v: %w", conferenceId, err)
	}
	return conference, nil
}

type UpdateConferenceInput struct {
	ConferenceId   string `json:"-"`                        // ConferenceId
	Status         string `json:"status,omitempty"`         // 	(optional) Setting the conference status to completed ends the conference.	No
	RedirectUrl    string `json:"redirectUrl,omitempty"`    // 	(optional) The URL to send the conferenceRedirect event which will provide new BXML. Not allowed if state is completed, but required if state is active	No
	RedirectMethod string `json:"redirectMethod,omitempty"` // 	(optional) The HTTP method to use for the request to redirectUrl Not allowed if state is completed	No
	Username       string `json:"username,omitempty"`       // 	(optional) The username to send in the HTTP request to redirectUrl	No
	Password       string `json:"password,omitempty"`       // 	(optional) The password to send in the HTTP request to redirectUrl	No
}

// UpdateConference - Update an active conference.
// https://dev.bandwidth.com/voice/methods/conferences/postConferencesConferenceId.html
func (v *Voice) UpdateConference(ctx context.Context, input UpdateConferenceInput) error {
	path := filepath.Join("/conferences", input.ConferenceId)
	if err := v.client.Post(ctx, path, input, nil); err != nil {
		return fmt.Errorf("unable to update conference, %v: %w", input.ConferenceId, err)
	}
	return nil
}

func (v *Voice) FindConferenceMember(ctx context.Context, conferenceID, memberID string) (member ConferenceMember, err error) {
	path := filepath.Join("/conferences", conferenceID, "members", memberID)
	if err := v.client.Get(ctx, path, &member); err != nil {
		return ConferenceMember{}, fmt.Errorf("unable to find member, %v, in conference, %v: %w", memberID, conferenceID, err)
	}
	return member, nil
}

type UpdateConferenceMemberInput struct {
	ConferenceId   string `json:"-"`
	MemberId       string `json:"-"`
	Mute           string `json:"mute,omitempty"`           // 	(optional) If true, member can't speak in the conference. If omitted, the parameter will not be modified.	No
	Hold           string `json:"hold,omitempty"`           // 	(optional) If true, member can't speak or hear in the conference. If omitted, the parameter will not be modified.	No
	CallIdsToCoach string `json:"callIdsToCoach,omitempty"` // 	(optional) Updates the list of calls to be coached by this member.
}

func (v *Voice) UpdateConferenceMember(ctx context.Context, input UpdateConferenceMemberInput) error {
	path := filepath.Join("/conferences", input.ConferenceId, "members", input.MemberId)
	if err := v.client.Put(ctx, path, input, nil); err != nil {
		return fmt.Errorf("unable to update member, %v, of conference, %v: %w", input.MemberId, input.ConferenceId, err)
	}
	return nil
}
