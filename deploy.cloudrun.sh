#!/bin/bash
# https://cloud.google.com/run/docs/quickstarts/build-and-deploy

gcloud config set project my-project

gcloud builds submit --tag gcr.io/my-project/cors-unblocker

gcloud run deploy --image gcr.io/my-project/cors-unblocker --platform managed
