#!/bin/bash

kubectl delete -f 14-application.yaml

kubectl delete -f 13-instantiate-chaincode.yaml

kubectl delete -f 12-install-chaincode.yaml

kubectl delete -f 11-join-channel.yaml

kubectl delete -f 10-create-channel.yaml

kubectl delete -f 09-org4.yaml

kubectl delete -f 08-org3.yaml

kubectl delete -f 07-org2.yaml

kubectl delete -f 06-org1.yaml

kubectl delete -f 05-ca.yaml

kubectl delete -f 04-orderer.yaml

kubectl delete -f 03-utils.yaml

kubectl delete -f 02-docker.yaml

kubectl delete -f 01-nfs.yaml
