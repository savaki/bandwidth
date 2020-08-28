package bandwidth

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestCall(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/call.json")
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	var call Call
	err = json.Unmarshal(data, &call)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
	if got, want := call.CallId, "c-d45a41e5-5eba9664-174a-4a1f-86ab-c947e393e4ee"; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
