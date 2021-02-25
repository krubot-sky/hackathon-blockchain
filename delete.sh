#!/bin/bash

kubectl delete -f 15-application.yaml

kubectl delete -f 14-instantiate-chaincode.yaml

kubectl delete -f 13-install-chaincode.yaml

kubectl delete -f 12-join-channel.yaml

kubectl delete -f 11-create-channel.yaml

kubectl delete -f 10-org4.yaml

kubectl delete -f 09-org3.yaml

kubectl delete -f 08-org2.yaml

kubectl delete -f 07-org1.yaml

kubectl delete -f 06-ca.yaml

kubectl delete -f 05-orderer.yaml

kubectl delete -f 04-utils.yaml

kubectl delete -f 03-docker.yaml

kubectl delete -f 02-nfs.yaml

kubectl delete -f 01-permissions.yaml

kubectl -n ledger delete pvc docker-docker-dind-0 nfs-nfs-server-0
