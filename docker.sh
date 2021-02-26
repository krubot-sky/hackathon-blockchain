#!/bin/bash

docker build -t eu.gcr.io/disco-mgmt-100/hackathon/app:latest application
docker build -t us.gcr.io/nbcu-disco-mgmt-003/hackathon/app:latest application

docker push eu.gcr.io/disco-mgmt-100/hackathon/app:latest
docker push us.gcr.io/nbcu-disco-mgmt-003/hackathon/app:latest
