```bash
export PROJECT_ID=online-game-demo-kl
export FOLDER_ID=<specify folder>
export BILLING_ID=<specify billing id>
export GCP_REGION=us-central1
export GCP_ZONE=us-central1-b
gcloud projects create ${PROJECT_ID} \
    --folder=${FOLDER_ID} \
    --enable-cloud-apis
gcloud beta billing projects link ${PROJECT_ID} \
    --billing-account ${BILLING_ID}

gcloud config set project ${PROJECT_ID}
# Enable Services
gcloud services enable container.googleapis.com compute.googleapis.com cloudbuild.googleapis.com
```