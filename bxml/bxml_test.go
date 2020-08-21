package bxml

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	err := Write(buf,
		Forward{
			From: "+18005551212",
			To:   "+18885551212",
		},
		Pause{Duration: 12},
	)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
	want := `<?xml version="1.0" encoding="UTF-8"?><Response><Forward from="+18005551212" to="+18885551212"></Forward><Pause duration="12"></Pause></Response>`
	if got := buf.String(); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
