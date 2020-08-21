package bandwidth

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Event interface {
}

// AnswerEvent - https://dev.bandwidth.com/voice/bxml/callbacks/answer.html
type AnswerEvent struct {
	EventType     string `json:"eventType,omitempty"`     // 	The event type, value is answer
	AccountId     string `json:"accountId,omitempty"`     // 	The user account associated with the call.
	ApplicationId string `json:"applicationId,omitempty"` // 	The id of the application associated with the call.
	To            string `json:"to,omitempty"`            // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From          string `json:"from,omitempty"`          // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction     string `json:"direction,omitempty"`     // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId        string `json:"callId,omitempty"`        // 	The call id associated with the event.
	CallUrl       string `json:"callUrl,omitempty"`       // 	The URL of the call associated with the event.
	StartTime     string `json:"startTime,omitempty"`     // 	Time the call was started, in ISO 8601 format.
	AnswerTime    string `json:"answerTime,omitempty"`    // 	Time the call was answered, in ISO 8601 format.
	Tag           string `json:"tag,omitempty"`           // 	(optional) The tag specified on call creation. If no tag was specified or it was previously cleared, null.
}

// BridgeCompleteEvent - https://dev.bandwidth.com/voice/bxml/callbacks/bridgeComplete.html
type BridgeCompleteEvent struct {
	EventType     string `json:"eventType,omitempty"`     // 	The event type, value is bridgeComplete.
	AccountId     string `json:"accountId,omitempty"`     // 	The user account associated with the call.
	ApplicationId string `json:"applicationId,omitempty"` // 	The id of the application associated with the call.
	From          string `json:"from,omitempty"`          // 	The phone number used in the from field of the original call, in E.164 format (e.g. +15555555555).
	To            string `json:"to,omitempty"`            // 	The phone number user in the to field of the original call, in E.164 format (e.g. +15555555555).
	Direction     string `json:"direction,omitempty"`     // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId        string `json:"callId,omitempty"`        // 	The call id associated with the event.
	CallUrl       string `json:"callUrl,omitempty"`       // 	The URL of the call associated with the event.
	StartTime     string `json:"startTime,omitempty"`     // 	Time the call was started, in ISO 8601 format.
	AnswerTime    string `json:"answerTime,omitempty"`    // 	Time the call was answered, in ISO 8601 format.
	Tag           string `json:"tag,omitempty"`           // 	The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	Cause         string `json:"cause,omitempty"`         // 	Reason the bridge failed - busy, rejected, or unknown.
	ErrorMessage  string `json:"errorMessage,omitempty"`  // 	Text explaining the reason that caused the bridge to fail in case of errors.
	ErrorId       string `json:"errorId,omitempty"`       // 	Bandwidth internal id that references the error event.
}

// BridgeTargetCompleteEvent - https://dev.bandwidth.com/voice/bxml/callbacks/bridgeTargetComplete.html
type BridgeTargetCompleteEvent struct {
	EventType     string `json:"eventType,omitempty"`     // 	The event type, value is bridgeTargetComplete.
	AccountId     string `json:"accountId,omitempty"`     // 	The user account associated with the call.
	ApplicationId string `json:"applicationId,omitempty"` // 	The id of the application associated with the call.
	From          string `json:"from,omitempty"`          // 	The phone number used in the from field of the original call, in E.164 format (e.g. +15555555555).
	To            string `json:"to,omitempty"`            // 	The phone number user in the to field of the original call, in E.164 format (e.g. +15555555555).
	Direction     string `json:"direction,omitempty"`     // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId        string `json:"callId,omitempty"`        // 	The bridge target call id.
	CallUrl       string `json:"callUrl,omitempty"`       // 	The URL of the call associated with the event.
	StartTime     string `json:"startTime,omitempty"`     // 	Time the call was started, in ISO 8601 format.
	AnswerTime    string `json:"answerTime,omitempty"`    // 	Time the call was answered, in ISO 8601 format.
	Tag           string `json:"tag,omitempty"`           // 	The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
}

// ConferenceCreatedEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceCreated.html
type ConferenceCreatedEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceCreated.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the new conference that was created.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
	Tag          string `json:"tag,omitempty"`          // 	(optional) The tag that was set at conference creation, if any.
}

// ConferenceMemberJoinEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceMemberJoin.html
type ConferenceMemberJoinEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceMemberJoin.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the new conference that was created.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
	CallId       string `json:"callId,omitempty"`       // 	The callId of the member that left the conference.
	From         string `json:"from,omitempty"`         // 	The from number of the call that left the conference.
	To           string `json:"to,omitempty"`           // 	The to number of the call that left the conference.
	Tag          string `json:"tag,omitempty"`          // 	(optional) The tag that was set at conference creation, if any.
}

// ConferenceMemberExitEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceMemberExit.html
type ConferenceMemberExitEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceMemberExit.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the new conference that was created.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
	CallId       string `json:"callId,omitempty"`       // 	The callId of the member that left the conference.
	From         string `json:"from,omitempty"`         // 	The from number of the call that left the conference.
	To           string `json:"to,omitempty"`           // 	The to number of the call that left the conference.
	Tag          string `json:"tag,omitempty"`          // 	(optional) The tag that was set at conference creation, if any.
}

// ConferenceCompletedEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceCompleted.html
type ConferenceCompletedEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceCompleted.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the new conference that was created.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
	Tag          string `json:"tag,omitempty"`          // 	(optional) The tag that was set at conference creation, if any.
}

// ConferenceRedirectEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceRedirect.html
type ConferenceRedirectEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceRedirect.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the new conference that was created.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
}

// ConferenceRecordingAvailableEvent - https://dev.bandwidth.com/voice/bxml/callbacks/conferenceRecordingAvailable.html
type ConferenceRecordingAvailableEvent struct {
	EventType    string `json:"eventType,omitempty"`    // 	The event type, value is conferenceRecordingAvailable.
	ConferenceId string `json:"conferenceId,omitempty"` // 	The ID of the conference that the recording was made on.
	Name         string `json:"name,omitempty"`         // 	The custom name used to reference this conference. This the name that you included inside the body of the <Conference> tag.
	AccountId    string `json:"accountId,omitempty"`    // 	The user account associated with the conference.
	RecordingId  string `json:"recordingId,omitempty"`  // 	The unique id for this recording.
	Channels     string `json:"channels,omitempty"`     // 	Number of channels in the recording (always 1 for conference recordings).
	StartTime    string `json:"startTime,omitempty"`    // 	The time that the recording started (in ISO8601 format).
	EndTime      string `json:"endTime,omitempty"`      // 	The time that the recording ended (in ISO8601 format).
	Duration     string `json:"duration,omitempty"`     // 	The duration of the recording (in ISO8601 format).
	FileFormat   string `json:"fileFormat,omitempty"`   // 	The audio format that the recording was saved as (wav or mp3).
	MediaUrl     string `json:"mediaUrl,omitempty"`     // 	The URL of the recording media.
	Tag          string `json:"tag,omitempty"`          // 	(optional) The tag that was set at conference creation, if any.
	Status       string `json:"status,omitempty"`       // 	The state of the recording. Can be complete, partial, or error. A partial status indicates that, although the recording is available to be downloaded, parts of the recording are missing.
}

// DisconnectEvent - https://dev.bandwidth.com/voice/bxml/callbacks/disconnect.html
type DisconnectEvent struct {
	EventType     string `json:"eventType,omitempty"`     // 	The event type, value is disconnect
	AccountId     string `json:"accountId,omitempty"`     // 	The user account associated with the call.
	ApplicationId string `json:"applicationId,omitempty"` // 	The id of the application associated with the call.
	To            string `json:"to,omitempty"`            // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From          string `json:"from,omitempty"`          // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction     string `json:"direction,omitempty"`     // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId        string `json:"callId,omitempty"`        // 	The call id associated with the event.
	CallUrl       string `json:"callUrl,omitempty"`       // 	The URL of the call associated with the event.
	StartTime     string `json:"startTime,omitempty"`     // 	Time the call was started, in ISO 8601 format.
	AnswerTime    string `json:"answerTime,omitempty"`    // 	(optional) Time the call was answered, in ISO 8601 format.
	EndTime       string `json:"endTime,omitempty"`       // 	Time the call ended, in ISO 8601 format.
	Cause         string `json:"cause,omitempty"`         // 	Reason the call ended
	ErrorMessage  string `json:"errorMessage,omitempty"`  // 	(optional) Text explaining the reason that caused the call to be ended in case of errors.
	ErrorId       string `json:"errorId,omitempty"`       // 	(optional) Bandwidth internal id that references the error event.
	Tag           string `json:"tag,omitempty"`           // 	(optional) The tag specified on call creation. If no tag was specified or it was previously cleared, null.
}

// https://dev.bandwidth.com/voice/bxml/callbacks/gather.html
type GatherEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is gather.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	To               string `json:"to,omitempty"`               // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From             string `json:"from,omitempty"`             // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	(optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, null.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the call was started, in ISO 8601 format.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	Time the call was answered, in ISO 8601 format.
	Tag              string `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	Digits           string `json:"digits,omitempty"`           // 	(optional) The digits collected from user. Null if a timeout occurred before any digits were pressed.
	TerminatingDigit string `json:"terminatingDigit,omitempty"` // 	(optional) The digit the user pressed to end the gather. Null if no terminating digit was pressed.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). Otherwise, null.
	TransferTo       string `json:"transferTo,omitempty"`       // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the to field of the B-leg call in E.164 format (e.g. +15555555555). Otherwise, null.
}

// InitiateEvent - https://dev.bandwidth.com/voice/bxml/callbacks/initiate.html
type InitiateEvent struct {
	EventType     string `json:"eventType,omitempty"`     // 	The event type, value is initiate.
	AccountId     string `json:"accountId,omitempty"`     // 	The user account associated with the call.
	ApplicationId string `json:"applicationId,omitempty"` // 	The id of the application associated with the call.
	To            string `json:"to,omitempty"`            // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From          string `json:"from,omitempty"`          // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction     string `json:"direction,omitempty"`     // 	The direction of the call; can only be inbound. The direction never changes.
	CallId        string `json:"callId,omitempty"`        // 	The call id associated with the event.
	CallUrl       string `json:"callUrl,omitempty"`       // 	The URL of the call associated with the event.
	StartTime     string `json:"startTime,omitempty"`     // 	Time the call was started, in ISO 8601 format.
	Diversion     string `json:"diversion,omitempty"`     // 	(optional) Information from the most recent Diversion header, if any. If present, the value will be a sub-object like "diversion": {"param1": "value1", "param2": "value2"}.
}

// RecordCompleteEvent - https://dev.bandwidth.com/voice/bxml/callbacks/recordComplete.html
type RecordCompleteEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is recordComplete.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	To               string `json:"to,omitempty"`               // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From             string `json:"from,omitempty"`             // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	(optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, null.
	RecordingId      string `json:"recordingId,omitempty"`      // 	The unique id for this recording.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	MediaUrl         string `json:"mediaUrl,omitempty"`         // 	URL to retrieve the contents of the recording.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	Time the call was answered, in ISO 8601 format.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the recording was started, in ISO 8601 format.
	EndTime          string `json:"endTime,omitempty"`          // 	Time the recording ended, in ISO 8601 format.
	Duration         string `json:"duration,omitempty"`         // 	Duration of the recording, in ISO 8601 format.
	Channels         string `json:"channels,omitempty"`         // 	Number of channels in the recording.
	FileFormat       string `json:"fileFormat,omitempty"`       // 	The audio format that the recording was saved as (wav or mp3).
	Tag              string `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). Otherwise, null.
	TransferTo       string `json:"transferTo,omitempty"`       // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the to field of the B-leg call in E.164 format (e.g. +15555555555). Otherwise, null.
}

// RecordingAvailableEvent - https://dev.bandwidth.com/voice/bxml/callbacks/recordingAvailable.html
type RecordingAvailableEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is recordingAvailable.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	To               string `json:"to,omitempty"`               // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From             string `json:"from,omitempty"`             // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	(optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, null.
	RecordingId      string `json:"recordingId,omitempty"`      // 	The unique id for this recording.
	Channels         string `json:"channels,omitempty"`         // 	Number of channels in the recording (1 or 2).
	StartTime        string `json:"startTime,omitempty"`        // 	The time that the recording started (in ISO8601 format).
	EndTime          string `json:"endTime,omitempty"`          // 	The time that the recording ended (in ISO8601 format).
	Duration         string `json:"duration,omitempty"`         // 	The duration of the recording (in ISO8601 format).
	FileFormat       string `json:"fileFormat,omitempty"`       // 	The audio format that the recording was saved as (wav or mp3).
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	MediaUrl         string `json:"mediaUrl,omitempty"`         // 	The URL of the recording media.
	Tag              string `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	Status           string `json:"status,omitempty"`           // 	The state of the recording. Can be complete, partial, or error. A partial status indicates that, although the recording is available to be downloaded, parts of the recording are missing.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). Otherwise, null.
	TransferTo       string `json:"transferTo,omitempty"`       // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the to field of the B-leg call in E.164 format (e.g. +15555555555). Otherwise, null.
}

// TranscriptionEvent - https://dev.bandwidth.com/voice/bxml/callbacks/transcriptionAvailable.html
type TranscriptionEvent struct {
	Id            string `json:"id,omitempty"`            // 	The unique id of the transcription.
	Url           string `json:"url,omitempty"`           // 	URL to retrieve the transcription output.
	Status        string `json:"status,omitempty"`        // 	The state of the transcription. Can be available, error, timeout, file-size-too-big, file-size-too-small.
	CompletedTime string `json:"completedTime,omitempty"` // 	Time the transcription was completed (in ISO8601 format).
}

// TranscriptionAvailableEvent - https://dev.bandwidth.com/voice/bxml/callbacks/transcriptionAvailable.html
type TranscriptionAvailableEvent struct {
	EventType        string             `json:"eventType,omitempty"`        // 	The event type, value is transcriptionAvailable.
	AccountId        string             `json:"accountId,omitempty"`        // 	The account id associated with the event.
	ApplicationId    string             `json:"applicationId,omitempty"`    // 	The application id associated with the event.
	CallId           string             `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string             `json:"parentCallId,omitempty"`     // 	(optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, null.
	RecordingId      string             `json:"recordingId,omitempty"`      // 	The unique id for this recording.
	To               string             `json:"to,omitempty"`               // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From             string             `json:"from,omitempty"`             // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction        string             `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	Tag              string             `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call.
	StartTime        string             `json:"startTime,omitempty"`        // 	Time the recording started (in ISO8601 format).
	EndTime          string             `json:"endTime,omitempty"`          // 	Time the recording ended (in ISO8601 format).
	Duration         string             `json:"duration,omitempty"`         // 	Length of the recording (in ISO8601 format).
	FileFormat       string             `json:"fileFormat,omitempty"`       // 	The format that the recording was saved in - mp3 or wav.
	CallUrl          string             `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	MediaUrl         string             `json:"mediaUrl,omitempty"`         // 	URL to retrieve the contents of the recording.
	Transcription    TranscriptionEvent `json:"transcription,omitempty"`    // 	Transcription information, see below.
	TransferCallerId string             `json:"transferCallerId,omitempty"` // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). Otherwise, null.
	TransferTo       string             `json:"transferTo,omitempty"`       // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the to field of the B-leg call in E.164 format (e.g. +15555555555). Otherwise, null.
}

// RedirectEvent - https://dev.bandwidth.com/voice/bxml/callbacks/redirect.html
type RedirectEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is redirect.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	To               string `json:"to,omitempty"`               // 	The phone number that received the call, in E.164 format (e.g. +15555555555).
	From             string `json:"from,omitempty"`             // 	The phone number that made the call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	(optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, null.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the call was started, in ISO 8601 format.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	(optional) Time the call was answered, in ISO 8601 format.
	Tag              string `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). Otherwise, null.
	TransferTo       string `json:"transferTo,omitempty"`       // 	(optional) If the event is related to the B leg of a <Transfer>, the phone number used as the to field of the B-leg call in E.164 format (e.g. +15555555555). Otherwise, null.
}

// TransferAnswerEvent - https://dev.bandwidth.com/voice/bxml/callbacks/transferAnswer.html
type TransferAnswerEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is transferAnswer.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	From             string `json:"from,omitempty"`             // 	The phone number used in the from field of the original call, in E.164 format (e.g. +15555555555).
	To               string `json:"to,omitempty"`               // 	The phone number used in the to field of the original call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Always outbound for this event.
	CallId           string `json:"callId,omitempty"`           // 	The call id of the newly-created B leg.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	The call id of the original call leg that executed the <Transfer> tag.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the call was started, in ISO 8601 format.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	Time the call was answered, in ISO 8601 format.
	Tag              string `json:"tag,omitempty"`              // 	(optional) The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferTo       string `json:"transferTo,omitempty"`       // 	The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555).
}

// TransferCompleteEvent - https://dev.bandwidth.com/voice/bxml/callbacks/transferComplete.html
type TransferCompleteEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is transferComplete.
	AccountId        string `json:"accountId,omitempty"`        // 	The user account associated with the call.
	ApplicationId    string `json:"applicationId,omitempty"`    // 	The id of the application associated with the call.
	From             string `json:"from,omitempty"`             // 	The phone number used in the from field of the original call, in E.164 format (e.g. +15555555555).
	To               string `json:"to,omitempty"`               // 	The phone number user in the to field of the original call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Either inbound or outbound. The direction of a call never changes.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the call was started, in ISO 8601 format.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	Time the call was answered, in ISO 8601 format.
	Tag              string `json:"tag,omitempty"`              // 	The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferTo       string `json:"transferTo,omitempty"`       // 	The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555).
	Cause            string `json:"cause,omitempty"`            // 	Reason the call ended - busy, timeout, hangup, cancel, rejected, callback-error, invalid-bxml, account-limit, node-capacity-exceeded, error, unknown or application-error. hangup indicates the call has ended normally.
	ErrorMessage     string `json:"errorMessage,omitempty"`     // 	Text explaining the reason that caused the call to be ended in case of errors.
	ErrorId          string `json:"errorId,omitempty"`          // 	Bandwidth internal id that references the error event.
}

// TransferDisconnectEvent - https://dev.bandwidth.com/voice/bxml/callbacks/transferDisconnect.html
type TransferDisconnectEvent struct {
	EventType        string `json:"eventType,omitempty"`        // 	The event type, value is transferDisconnect.
	From             string `json:"from,omitempty"`             // 	The phone number used in the from field of the original call, in E.164 format (e.g. +15555555555).
	To               string `json:"to,omitempty"`               // 	The phone number user in the to field of the original call, in E.164 format (e.g. +15555555555).
	Direction        string `json:"direction,omitempty"`        // 	The direction of the call. Always outbound for this event.
	CallId           string `json:"callId,omitempty"`           // 	The call id associated with the event.
	ParentCallId     string `json:"parentCallId,omitempty"`     // 	The call id of the original call leg that contained the <Transfer> tag.
	CallUrl          string `json:"callUrl,omitempty"`          // 	The URL of the call associated with the event.
	Tag              string `json:"tag,omitempty"`              // 	The tag specified earlier in the call. If no tag was specified or it was previously cleared, null.
	StartTime        string `json:"startTime,omitempty"`        // 	Time the transferred leg was started, in ISO 8601 format.
	AnswerTime       string `json:"answerTime,omitempty"`       // 	(optional) Time the transferred leg was answered, in ISO 8601 format.
	EndTime          string `json:"endTime,omitempty"`          // 	Time the transferred leg ended, in ISO 8601 format.
	TransferCallerId string `json:"transferCallerId,omitempty"` // 	The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferTo       string `json:"transferTo,omitempty"`       // 	The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555).
	Cause            string `json:"cause,omitempty"`            // 	Reason the transferred leg ended - busy, timeout, hangup, cancel, rejected, callback-error, invalid-bxml, account-limit, node-capacity-exceeded, error, unknown or application-error. hangup indicates the call has ended normally.
	ErrorMessage     string `json:"errorMessage,omitempty"`     // 	Text explaining the reason that caused the transferred leg to be ended in case of errors.
	ErrorId          string `json:"errorId,omitempty"`          // 	Bandwidth internal id that references the error event.
}

var reEventType = regexp.MustCompile(`"eventType":\s*"([^"]+)"`)

func ParseEvent(data []byte) (Event, error) {
	matches := reEventType.FindSubmatch(data)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no event type found")
	}

	var (
		eventType = string(matches[1])
		event     Event
	)
	switch eventType {
	case "answer":
		event = &AnswerEvent{}
	case "bridgeComplete":
		event = &BridgeCompleteEvent{}
	case "bridgeTargetComplete":
		event = &BridgeTargetCompleteEvent{}
	case "conferenceCreated":
		event = &ConferenceCreatedEvent{}
	case "conferenceMemberJoin":
		event = &ConferenceMemberJoinEvent{}
	case "conferenceMemberExit":
		event = &ConferenceMemberExitEvent{}
	case "conferenceCompleted":
		event = &ConferenceCompletedEvent{}
	case "conferenceRedirect":
		event = &ConferenceRedirectEvent{}
	case "conferenceRecordingAvailable":
		event = &ConferenceRecordingAvailableEvent{}
	case "disconnect":
		event = &DisconnectEvent{}
	case "gather":
		event = &GatherEvent{}
	case "initiate":
		event = &InitiateEvent{}
	case "recordComplete":
		event = &RecordCompleteEvent{}
	case "recordingAvailable":
		event = &RecordingAvailableEvent{}
	case "transcriptionAvailable":
		event = &TranscriptionAvailableEvent{}
	case "redirect":
		event = &RedirectEvent{}
	case "transferAnswer":
		event = &TransferAnswerEvent{}
	case "transferComplete":
		event = &TransferCompleteEvent{}
	case "transferDisconnect":
		event = &TransferDisconnectEvent{}
	default:
		return nil, fmt.Errorf("unhandled event type, %v", eventType)
	}

	if err := json.Unmarshal(data, event); err != nil {
		return nil, fmt.Errorf("unable to unmarshal event, %v: %w", eventType, err)
	}

	return event, nil
}
