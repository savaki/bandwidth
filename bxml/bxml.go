package bxml

import (
	"encoding/xml"
	"io"
)

// Write the responses
func Write(w io.Writer, vv ...interface{}) error {
	if _, err := io.WriteString(w, `<?xml version="1.0" encoding="UTF-8"?>`); err != nil {
		return err
	}
	if _, err := io.WriteString(w, `<Response>`); err != nil {
		return err
	}

	encoder := xml.NewEncoder(w)
	for _, v := range vv {
		if err := encoder.Encode(v); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, `</Response>`); err != nil {
		return err
	}

	return nil
}

// Bridge - https://dev.bandwidth.com/voice/bxml/verbs/bridge.html
type Bridge struct {
	BridgeCompleteFallbackMethod       string `xml:"bridgeCompleteFallbackMethod,attr,omitempty"`       // (optional) The HTTP method to use to deliver the Bridge Complete callback to bridgeCompleteFallbackUrl. GET or POST. Default value is POST.
	BridgeCompleteFallbackUrl          string `xml:"bridgeCompleteFallbackUrl,attr,omitempty"`          // (optional) A fallback url which, if provided, will be used to retry the Bridge Complete callback delivery in case bridgeCompleteUrl fails to respond.
	BridgeCompleteMethod               string `xml:"bridgeCompleteMethod,attr,omitempty"`               // (optional) The HTTP method to use for the request to bridgeCompleteUrl. GET or POST. Default value is POST.
	BridgeCompleteUrl                  string `xml:"bridgeCompleteUrl,attr,omitempty"`                  // (optional) URL to send the Bridge Complete event to and request new BXML.
	BridgeTargetCompleteFallbackMethod string `xml:"bridgeTargetCompleteFallbackMethod,attr,omitempty"` // (optional) The HTTP method to use to deliver the Bridge Target Complete callback to bridgeTargetCompleteFallbackUrl. GET or POST. Default value is POST.
	BridgeTargetCompleteFallbackUrl    string `xml:"bridgeTargetCompleteFallbackUrl,attr,omitempty"`    // (optional) A fallback url which, if provided, will be used to retry the Bridge Target Complete callback delivery in case bridgeTargetCompleteUrl fails to respond.
	BridgeTargetCompleteMethod         string `xml:"bridgeTargetCompleteMethod,attr,omitempty"`         // (optional) The HTTP method to use for the request to bridgeTargetCompleteUrl. GET or POST. Default value is POST.
	BridgeTargetCompleteUrl            string `xml:"bridgeTargetCompleteUrl,attr,omitempty"`            // (optional) URL to send the Bridge Target Complete event to and request new BXML.
	FallbackPassword                   string `xml:"fallbackPassword,attr,omitempty"`                   // (optional) The password to send in the HTTP request to bridgeCompleteFallbackUrl and to bridgeTargetCompleteFallbackUrl.
	FallbackUsername                   string `xml:"fallbackUsername,attr,omitempty"`                   // (optional) The username to send in the HTTP request to bridgeCompleteFallbackUrl and to bridgeTargetCompleteFallbackUrl.
	Password                           string `xml:"password,attr,omitempty"`                           // (optional) The password to send in the HTTP request to bridgeCompleteUrl and to bridgeTargetCompleteUrl.
	Tag                                string `xml:"tag,attr,omitempty"`                                // (optional) A custom string that will be sent with the bridgeComplete callback and all future callbacks of the call unless overwritten by a future tag attribute or cleared.
	Username                           string `xml:"username,attr,omitempty"`                           // (optional) The username to send in the HTTP request to bridgeCompleteUrl and to bridgeTargetCompleteUrl.
}

// Conference - https://dev.bandwidth.com/voice/bxml/verbs/conference.html
type Conference struct {
	CallbackTimeout               int    `xml:"callbackTimeout,attr,omitempty"`               // (optional) This is the timeout (in seconds) to use when delivering callbacks for the conference. If not set, it will inherit the callback timeout from the call that creates the conference. Can be any numeric value (including decimals) between 1 and 25.
	CallIdsToCoach                string `xml:"callIdsToCoach,attr,omitempty"`                // (optional) A comma-separated list of call ids to coach. When a call joins a conference with this attribute set, it will coach the listed calls. Those calls will be able to hear and be heard by the coach, but other calls in the conference will not hear the coach.
	ConferenceEventFallbackMethod string `xml:"conferenceEventFallbackMethod,attr,omitempty"` // (optional) The HTTP method to use to deliver the conference callbacks to conferenceEventFallbackUrl. GET or POST. Default value is POST.
	ConferenceEventFallbackUrl    string `xml:"conferenceEventFallbackUrl,attr,omitempty"`    // (optional) A fallback url which, if provided, will be used to retry the conference callback deliveries in case conferenceEventUrl fails to respond.
	ConferenceEventMethod         string `xml:"conferenceEventMethod,attr,omitempty"`         // (optional) The HTTP method to use for the request to conferenceEventUrl. GET or POST. Default value is POST.
	ConferenceEventUrl            string `xml:"conferenceEventUrl,attr,omitempty"`            // (optional) URL to send Conference events to. The URL, method, username, and password are set by the BXML document that creates the conference, and all events related to that conference will be delivered to that same endpoint. If more calls join afterwards and also have this property (or any other callback related properties like username and password), they will be ignored and the original callback information will be used. This URL may be a relative endpoint.
	FallbackPassword              string `xml:"fallbackPassword,attr,omitempty"`              // (optional) The password to send in the HTTP request to conferenceEventFallbackUrl.
	FallbackUsername              string `xml:"fallbackUsername,attr,omitempty"`              // (optional) The username to send in the HTTP request to conferenceEventFallbackUrl.
	Hold                          bool   `xml:"hold,attr,omitempty"`                          // (optional) A boolean value to indicate whether the member should be on hold in the conference. When on hold, a member cannot hear others, and they cannot be heard. Defaults to false.
	Mute                          bool   `xml:"mute,attr,omitempty"`                          // (optional) A boolean value to indicate whether the member should be on mute in the conference. When muted, a member can hear others speak, but others cannot hear them speak. Defaults to false.
	Password                      string `xml:"password,attr,omitempty"`                      // (optional) The password to send in the HTTP request to conferenceEventUrl.
	Tag                           string `xml:"tag,attr,omitempty"`                           // (optional) A custom string that will be sent with these and all future callbacks unless overwritten by a future tag attribute or cleared.
	Username                      string `xml:"username,attr,omitempty"`                      // (optional) The username to send in the HTTP request to conferenceEventUrl.
}

// https://dev.bandwidth.com/voice/bxml/verbs/forward.html
type Forward struct {
	CallTimeout        int    `xml:"callTimeout,attr,omitempty"`        // (optional) Number of seconds to wait for an answer before abandoning the call. Range: decimal values between 1 - 300. Default: 30
	DiversionReason    string `xml:"diversionReason,attr,omitempty"`    // (optional) Can be any of the following values:
	DiversionTreatment string `xml:"diversionTreatment,attr,omitempty"` // (optional) Can be any of the following:
	From               string `xml:"from,attr,omitempty"`               // (optional) Number to use for caller ID on the outgoing leg. Must be in E.164 format (e.g. +15555555555). If omitted, assumes the "to" number of the original leg.
	To                 string `xml:"to,attr,omitempty"`                 // Number to forward the call to. Must be in E.164 format (e.g. +15555555555)
}

// https://dev.bandwidth.com/voice/bxml/verbs/gather.html
type Gather struct {
	FallbackPassword     string `xml:"fallbackPassword,attr,omitempty"`     // 	(optional) The password to send in the HTTP request to gatherFallbackUrl.
	FallbackUsername     string `xml:"fallbackUsername,attr,omitempty"`     // 	(optional) The username to send in the HTTP request to gatherFallbackUrl.
	FirstDigitTimeout    int    `xml:"firstDigitTimeout,attr,omitempty"`    // 	(optional) Time (in seconds) to pause after any audio from nested <SpeakSentence> or <PlayAudio> verb is played (in seconds) before terminating the Gather. Default value is 5. Range: decimal values between 0 - 60.
	GatherFallbackMethod string `xml:"gatherFallbackMethod,attr,omitempty"` // 	(optional) The HTTP method to use to deliver the Gather event callback to gatherFallbackUrl. GET or POST. Default value is POST.
	GatherFallbackUrl    string `xml:"gatherFallbackUrl,attr,omitempty"`    // 	(optional) A fallback url which, if provided, will be used to retry the Gather event callback delivery in case gatherUrl fails to respond.
	GatherMethod         string `xml:"gatherMethod,attr,omitempty"`         // 	(optional) The HTTP method to use for the request to gatherUrl. GET or POST. Default value is POST.
	GatherUrl            string `xml:"gatherUrl,attr,omitempty"`            // 	(optional) URL to send Gather event to and request new BXML. May be a relative URL.
	InterDigitTimeout    int    `xml:"interDigitTimeout,attr,omitempty"`    // 	(optional) Time (in seconds) allowed between digit presses before automatically terminating the Gather. Default value is 5. Range: decimal values between 1 - 60.
	MaxDigits            int    `xml:"maxDigits,attr,omitempty"`            // 	(optional) Max number of digits to collect. Default value is 50. Range: decimal values between 1 - 50.
	Password             string `xml:"password,attr,omitempty"`             // 	(optional) The password to send in the HTTP request to gatherUrl.
	RepeatCount          int    `xml:"repeatCount,attr,omitempty"`          // 	(optional) The number of times the audio prompt should be played if no digits are pressed. For example, if this value is 3, the nested audio clip will be played a maximum of three times. The delay between repetitions will be equal to firstDigitTimeout. Default value is 1. repeatCount * number of verbs must not be greater than 20.
	Tag                  string `xml:"tag,attr,omitempty"`                  // 	(optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	TerminatingDigits    string `xml:"terminatingDigits,attr,omitempty"`    // 	(optional) When any of these digits are pressed, it will terminate the Gather. Default value is "", which disables this feature.
	Username             string `xml:"username,attr,omitempty"`             // 	(optional) The username to send in the HTTP request to gatherUrl.
}

// Hangup - https://dev.bandwidth.com/voice/bxml/verbs/hangup.html
type Hangup struct {
}

// Pause - https://dev.bandwidth.com/voice/bxml/verbs/pause.html
type Pause struct {
	Duration int `xml:"duration,attr,omitempty"` // (optional) The 'duration' attribute specifies how many seconds Bandwidth will wait silently before continuing on. Default value is 1. Range: decimal values between 0.1 - 86400.
}

// PauseRecording - https://dev.bandwidth.com/voice/bxml/verbs/pauseRecording.html
type PauseRecording struct {
}

// PlayAudio - https://dev.bandwidth.com/voice/bxml/verbs/playAudio.html
type PlayAudio struct {
	Username string `xml:"username,attr,omitempty"` // (optional) The username to send in the HTTP request to audioUri.
	Password string `xml:"password,attr,omitempty"` // (optional) The password to send in the HTTP request to audioUri.
}

// Record - https://dev.bandwidth.com/voice/bxml/verbs/record.html
type Record struct {
	FallbackPassword             string `xml:"fallbackPassword,attr,omitempty"`             //     (optional) The password to send in the HTTP request to recordCompleteFallbackUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
	FallbackUsername             string `xml:"fallbackUsername,attr,omitempty"`             //     (optional) The username to send in the HTTP request to recordCompleteFallbackUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
	FileFormat                   string `xml:"fileFormat,attr,omitempty"`                   //     (optional) The audio format that the recording will be saved as: mp3 or wav.Default value is wav.
	MaxDuration                  int    `xml:"maxDuration,attr,omitempty"`                  //     (optional) Maximum length of recording (in seconds).Max 10800 (3 hours).Default value is 60.
	Password                     string `xml:"password,attr,omitempty"`                     //     (optional) The password to send in the HTTP request to recordCompleteUrl, recordingAvailableUrl or transcriptionAvailableUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
	RecordCompleteFallbackMethod string `xml:"recordCompleteFallbackMethod,attr,omitempty"` //     (optional) The HTTP method to use to deliver the Record Complete callback to recordCompleteFallbackUrl.GET or POST.Default value is POST.
	RecordCompleteFallbackUrl    string `xml:"recordCompleteFallbackUrl,attr,omitempty"`    //     (optional) A fallback url which, if provided, will be used to retry the Record Complete callback delivery in case recordCompleteUrl fails to respond.
	RecordCompleteMethod         string `xml:"recordCompleteMethod,attr,omitempty"`         //     (optional) The HTTP method to use for the request to recordCompleteUrl.GET or POST.Default value is POST.
	RecordCompleteUrl            string `xml:"recordCompleteUrl,attr,omitempty"`            //  (optional) URL to send the Record Complete event to once it has ended.Accepts BXML, and may be a relative URL.
	RecordingAvailableMethod     string `xml:"recordingAvailableMethod,attr,omitempty"`     //     (optional) The HTTP method to use for the request to recordingAvailableUrl.GET or POST.Default value is POST.
	RecordingAvailableUrl        string `xml:"recordingAvailableUrl,attr,omitempty"`        //     (optional) URL to send the Recording Available event to once it has been processed.Does not accept BXML.May be a relative URL.
	SilenceTimeout               int    `xml:"silenceTimeout,attr,omitempty"`               //     (optional) Length of silence after which to end the recording (in seconds).Max is equivalent to the maximum maxDuration value.Default value is 0, which means no timeout.
	Tag                          string `xml:"tag,attr,omitempty"`                          //     (optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	TerminatingDigits            string `xml:"terminatingDigits,attr,omitempty"`            //     (optional) When pressed, this digit will terminate the recording.Default value is “#”.This feature can be disabed with "".
	Transcribe                   bool   `xml:"transcribe,attr,omitempty"`                   //     (optional) A boolean value to indicate that recording should be transcribed.Transcription can succeed only for recordings of length greater than 500 milliseconds and less than 4 hours.Default is false.
	TranscriptionAvailableMethod string `xml:"transcriptionAvailableMethod,attr,omitempty"` //     (optional) The HTTP method to use for the request to transcriptionAvailableUrl.GET or POST.Default value is POST.
	TranscriptionAvailableUrl    string `xml:"transcriptionAvailableUrl,attr,omitempty"`    //     (optional) URL to send the Transcription Available event to once it has been processed.Does not accept BXML.May be a relative URL.
	Username                     string `xml:"username,attr,omitempty"`                     //     (optional) The username to send in the HTTP request to recordCompleteUrl, recordingAvailableUrl or transcriptionAvailableUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
}

// Redirect - https://dev.bandwidth.com/voice/bxml/verbs/redirect.html
type Redirect struct {
	RedirectUrl            string `xml:"redirectUrl,attr,omitempty"`            // 	(required) URL to request new BXML from. A Redirect event will be sent to this endpoint. May be a relative URL.
	RedirectMethod         string `xml:"redirectMethod,attr,omitempty"`         // 	(optional) The HTTP method to use for the request to redirectUrl. GET or POST. Default Value is POST.
	RedirectFallbackUrl    string `xml:"redirectFallbackUrl,attr,omitempty"`    // 	(optional) A fallback url which, if provided, will be used to retry the Redirect callback delivery in case redirectUrl fails to respond.
	RedirectFallbackMethod string `xml:"redirectFallbackMethod,attr,omitempty"` // 	(optional) The HTTP method to use to deliver the Redirect callback to redirectFallbackUrl. GET or POST. Default value is POST.
	Username               string `xml:"username,attr,omitempty"`               // 	(optional) The username to send in the HTTP request to redirectUrl.
	Password               string `xml:"password,attr,omitempty"`               // 	(optional) The password to send in the HTTP request to redirectUrl.
	FallbackUsername       string `xml:"fallbackUsername,attr,omitempty"`       // 	(optional) The username to send in the HTTP request to redirectFallbackUrl.
	FallbackPassword       string `xml:"fallbackPassword,attr,omitempty"`       // 	(optional) The password to send in the HTTP request to redirectFallbackUrl.
	Tag                    string `xml:"tag,attr,omitempty"`                    // 	(optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
}

// ResumeRecording - https://dev.bandwidth.com/voice/bxml/verbs/resumeRecording.html
type ResumeRecording struct {
}

// Ring - https://dev.bandwidth.com/voice/bxml/verbs/ring.html
type Ring struct {
	Duration int `xml:"duration,attr,omitempty"` // 	(optional) How many seconds to play ringing on the call. Default value is 5. Range: decimal values between 0.1 - 86400.
}

// SendDtmf - https://dev.bandwidth.com/voice/bxml/verbs/sendDtmf.html
type SendDtmf struct {
	ToneDuration string `xml:"toneDuration,attr,omitempty"` // 	(optional) The length (in milliseconds) of each DTMF tone. Default value is 200. Range: decimal values between 50 - 5000.
	ToneInterval string `xml:"toneInterval,attr,omitempty"` // 	(optional) The duration of silence (in milliseconds) following each DTMF tone. Default value is 400. Range: decimal values between 50 - 5000.
}

// SpeakSentence - https://dev.bandwidth.com/voice/bxml/verbs/speakSentence.html
type SpeakSentence struct {
	Voice  string `xml:"voice,attr,omitempty"`  // 	Selects the voice of the speaker. Consult the voice column in the below table for valid values.
	Gender string `xml:"gender,attr,omitempty"` // 	Selects the gender of the speaker. Valid values are "male" or "female".
	Locale string `xml:"locale,attr,omitempty"` // 	Selects the locale of the speaker. Consult locale column in the below table for valid values.
}

// StartRecording - https://dev.bandwidth.com/voice/bxml/verbs/startRecording.html
type StartRecording struct {
	RecordingAvailableUrl        string `xml:"recordingAvailableUrl,attr,omitempty"`        //  (optional) URL to send the Recording Available event (or Conference Recording Available event if recording a conference) to once it has been processed.Does not accept BXML.May be a relative URL.
	RecordingAvailableMethod     string `xml:"recordingAvailableMethod,attr,omitempty"`     //     (optional) The HTTP method to use for the request to recordingAvailableUrl.GET or POST.Default value is POST.
	Transcribe                   string `xml:"transcribe,attr,omitempty"`                   //     (optional) A boolean value to indicate that recording should be transcribed.Transcription can succeed only for recordings of length greater than 500 milliseconds and less than 4 hours.Default is false.
	TranscriptionAvailableUrl    string `xml:"transcriptionAvailableUrl,attr,omitempty"`    //     (optional) URL to send the Transcription Available event to once it has been processed.Does not accept BXML.May be a relative URL.
	TranscriptionAvailableMethod string `xml:"transcriptionAvailableMethod,attr,omitempty"` //     (optional) The HTTP method to use for the request to transcriptionAvailableUrl.GET or POST.Default value is POST.
	Username                     string `xml:"username,attr,omitempty"`                     //     (optional) The username to send in the HTTP request to recordingAvailableUrl or transcriptionAvailableUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
	Password                     string `xml:"password,attr,omitempty"`                     //     (optional) The password to send in the HTTP request to recordingAvailableUrl or transcriptionAvailableUrl.If specified, the URLs must be TLS-encrypted (i.e., https).
	Tag                          string `xml:"tag,attr,omitempty"`                          //     (optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	FileFormat                   string `xml:"fileFormat,attr,omitempty"`                   //     (optional) The audio format that the recording will be saved as: mp3 or wav.Default value is wav.
	MultiChannel                 string `xml:"multiChannel,attr,omitempty"`                 //     (optional) A boolean value indicating whether or not the recording file should separate each side of the call into its own audio channel.Default value is false.
}

// StopRecording -
type StopRecording struct {
}

// Transfer - https://dev.bandwidth.com/voice/bxml/verbs/transfer.html
type Transfer struct {
	TransferCallerId               string `xml:"transferCallerId,attr,omitempty"`               // 	(optional) The caller ID to use when the call is transferred, if different. Must be in E.164 format (e.g. +15555555555).
	CallTimeout                    string `xml:"callTimeout,attr,omitempty"`                    // 	(optional) This is the timeout (in seconds) for the callee to answer the call. Range: decimal values between 1 - 300. Default value is 30 seconds.
	TransferCompleteUrl            string `xml:"transferCompleteUrl,attr,omitempty"`            // 	(optional) URL to send the Transfer Complete event to and request new BXML. Optional but recommended. See below for further details. May be a relative URL.
	TransferCompleteMethod         string `xml:"transferCompleteMethod,attr,omitempty"`         // 	(optional) The HTTP method to use for the request to transferCompleteUrl. GET or POST. Default value is POST.
	TransferCompleteFallbackUrl    string `xml:"transferCompleteFallbackUrl,attr,omitempty"`    // 	(optional) A fallback url which, if provided, will be used to retry the Transfer Complete callback delivery in case transferCompleteUrl fails to respond.
	TransferCompleteFallbackMethod string `xml:"transferCompleteFallbackMethod,attr,omitempty"` // 	(optional) The HTTP method to use to deliver the Transfer Complete callback to transferCompleteFallbackUrl. GET or POST. Default value is POST.
	Username                       string `xml:"username,attr,omitempty"`                       // 	(optional) The username to send in the HTTP request to transferCompleteUrl.
	Password                       string `xml:"password,attr,omitempty"`                       // 	(optional) The password to send in the HTTP request to transferCompleteUrl.
	FallbackUsername               string `xml:"fallbackUsername,attr,omitempty"`               // 	(optional) The username to send in the HTTP request to transferCompleteFallbackUrl.
	FallbackPassword               string `xml:"fallbackPassword,attr,omitempty"`               // 	(optional) The password to send in the HTTP request to transferCompleteFallbackUrl.
	Tag                            string `xml:"tag,attr,omitempty"`                            // 	(optional) A custom string that will be sent with this and all future callbacks unless overwritten by a future tag attribute or cleared.
	DiversionTreatment             string `xml:"diversionTreatment,attr,omitempty"`             // 	(optional) Can be any of the following:
	DiversionReason                string `xml:"diversionReason,attr,omitempty"`                // 	(optional) Can be any of the following values:
}
