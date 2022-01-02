package appnubankroutinecheckemail

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type EmailNotified struct {
	EmailAddress string `json:"emailAddress"`
	HistoryID    uint64 `json:"historyId"`
}

func ReadEmail(ctx context.Context, m PubSubMessage) error {

	token := new(oauth2.Token)
	token.AccessToken = ""
	token.RefreshToken = ""
	token.Expiry = time.Time{}
	token.TokenType = "Bearer"
	config := &oauth2.Config{}
	gmailService, err := gmail.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	payload := string(m.Data)
	emailNotified := EmailNotified{}
	json.Unmarshal([]byte(payload), &emailNotified)

	if handleError("Create Context", err) {
		return err
	}

	resp, errHistoryList := gmailService.Users.History.List("").LabelId("Label_1782046973960748351").StartHistoryId(emailNotified.HistoryID).Do()

	if handleError("History List", errHistoryList) {
		return errHistoryList
	}

	obj, errMarshal := resp.MarshalJSON()

	if handleError("Marshall", errMarshal) {
		return errMarshal
	}

	log.Printf(">>> %s", obj)
	return nil
}

func handleError(phase string, err error) bool {
	if err != nil {
		log.Println("#######", phase, "Error:", err)
		return true
	}
	return false
}
