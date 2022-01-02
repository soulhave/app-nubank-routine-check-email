# App Nubank Routine Check Email

App responsible to check new nubank emails on the pubsub subscription.

Run Local
```shell
cd cmd && go run main.go
```

Deploy Cloud Functions
```shell
gcloud functions deploy ReadEmail --runtime go116 --trigger-topic "nubank-cc-account-detailed" --env-vars-file .env.yaml
```

> Disclaimer: I encourage you to take a look on the google functions documention to understand how to deploy the application and the var environments.

You need to set some env vars to be used on your machine or deploy the function.

1. `APP_NUBANK_MAIL_LABEL:`
Define your label on the gmail that will be flaged the nunbak e-mails.
2. `APP_NUBANK_BUCKET:`
The bucket on GCP where you store all secret information about your account.
3. `APP_NUBANK_SECRET_FILE:`
The secret file with the data about your email account.