# App Nubank Routine Start Watch Email

App responsible to enable the e-mail to be watched and notify a pubsub queue.

Run Local
```shell
cd cmd && go run main.go
```

Deploy Cloud Functions
```shell
gcloud functions deploy ReadEmail --runtime go116 --trigger-topic "nubank-cc-account-detailed"
```