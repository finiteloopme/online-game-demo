```bash
# Unit testing
go test

# Build the container & push to Container Registry
gcloud auth configue-docker
gcloud builds submit --tag gcr.io/${PROJECT_ID}/mbaas

# Run your container locally
docker run --rm -p 8080:8080 gcr.io/${PROJECT_ID}/mbaas

# Deploy to Cloud Run
gcloud run deploy --image gcr.io/${PROJECT_ID}/mbaas --platform managed

# Deploy to GKE
DEPLOY_APP=mbaas-app
kubectl create deployment ${DEPLOY_APP} --image=gcr.io/${PROJECT_ID}/mbaas
kubectl scale deployment ${DEPLOY_APP} --replicas=3
kubectl autoscale deployment ${DEPLOY_APP} --cpu-percent=80 --min=1 --max=10
kubectl get pods
kubectl expose deployment ${DEPLOY_APP} --name=${DEPLOY_APP}-service --type=LoadBalancer --port 80 --target-port 8080
kubectl get service
```