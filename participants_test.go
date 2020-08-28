package bandwidth

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestWebRTC_CreateParticipant(t *testing.T) {
	var (
		accountID = os.Getenv("BANDWIDTH_ACCOUNT_ID")
		username  = os.Getenv("BANDWIDTH_USERNAME")
		password  = os.Getenv("BANDWIDTH_PASSWORD")
	)

	if accountID == "" || username == "" || password == "" {
		t.SkipNow()
	}

	ctx := context.Background()
	webRTC := NewWebRTC(WithCredentials(accountID, username, password))

	create := CreateParticipantInput{
		CallbackUrl:        "",
		PublishPermissions: []string{"AUDIO"},
		Subscriptions:      Subscriptions{},
		Tag:                "",
	}
	output, err := webRTC.CreateParticipant(ctx, create)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
	fmt.Println(output)
}
