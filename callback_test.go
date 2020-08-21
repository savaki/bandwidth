package bandwidth

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseEvent(t *testing.T) {
	testCases := map[string]struct {
		Filename string
		Want     interface{}
	}{
		"answer": {
			Filename: "testdata/answer.json",
			Want: &AnswerEvent{
				EventType:     "answer",
				AccountId:     "55555555",
				ApplicationId: "7fc9698a-b04a-468b-9e8f-91238c0d0086",
				To:            "+15553334444",
				From:          "+15551112222",
				Direction:     "outbound",
				CallId:        "c-95ac8d6e-1a31c52e-b38f-4198-93c1-51633ec68f8d",
				CallUrl:       "https://voice.bandwidth.com/api/v2/accounts/55555555/calls/c-95ac8d6e-1a31c52e-b38f-4198-93c1-51633ec68f8d",
				StartTime:     "2019-06-20T15:54:22.234Z",
				AnswerTime:    "2019-06-20T15:54:25.432Z",
				Tag:           "example-tag",
			},
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			v := parseEventFile(t, tc.Filename)
			fmt.Println(v)
		})
	}
}

func parseEventFile(t *testing.T, filename string) Event {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	event, err := ParseEvent(data)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	return event
}
