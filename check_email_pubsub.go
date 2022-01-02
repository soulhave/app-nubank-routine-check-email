package appnubankroutinecheckemail

import (
	"context"
	"encoding/json"
	"log"
	"os"

	".app-nubank-routine-check-email/gateways"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const _appNubankMailLabel = "APP_NUBANK_MAIL_LABEL"

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type EmailNotified struct {
	EmailAddress string `json:"emailAddress"`
	HistoryID    uint64 `json:"historyId"`
}

func ReadEmail(ctx context.Context, m PubSubMessage) error {
	gmailService, err := gmail.NewService(ctx, option.WithTokenSource(gateways.NewGoogleToken().GetTokenSource()))

	payload := string(m.Data)
	emailNotified := EmailNotified{}
	json.Unmarshal([]byte(payload), &emailNotified)

	if gateways.HandleError("Create Context", err) {
		return err
	}

	resp, errHistoryList := gmailService.Users.History.List("").LabelId(os.Getenv(_appNubankMailLabel)).StartHistoryId(emailNotified.HistoryID).Do()

	if gateways.HandleError("History List", errHistoryList) {
		return errHistoryList
	}

	obj, errMarshal := resp.MarshalJSON()

	if gateways.HandleError("Marshall", errMarshal) {
		return errMarshal
	}

	log.Printf(">>> %s", obj)
	return nil
}
