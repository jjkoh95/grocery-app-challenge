# setup golang libs
# go mod init
# go mod vendor

# deploy to GCP
gcloud functions deploy GetGrocery --runtime go111 --trigger-http --timeout=30s --memory=128MB --region=asia-east2 --env-vars-file .env.yaml
gcloud functions deploy UpsertGrocery --runtime go111 --trigger-http --timeout=30s --memory=128MB --region=asia-east2 --env-vars-file .env.yaml