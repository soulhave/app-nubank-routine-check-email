# App Nubank Routine Check Email

App responsible to check new nubank emails on the pubsub subscription.

Run Local
```shell
cd cmd && go run main.go
```

Deploy Cloud Functions
```shell
gcloud functions deploy ReadEmail --runtime go116 --trigger-topic "nubank-cc-account-detailed"
```