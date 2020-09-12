```bash
export PROJECT_ID=online-game-demo-kl
export FOLDER_ID=<specify folder>
export BILLING_ID=<specify billing id>
gcloud projects create ${PROJECT_ID} \
    --folder=${FOLDER_ID} \
    --enable-cloud-apis
gcloud beta billing projects link ${PROJECT_ID} \
    --billing-account ${BILLING_ID}

# Enable Services
gcloud services enable container.googleapis.com 
```