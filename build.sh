#!/bin/bash

kubectl apply -f 00-namespace.yaml

kubectl apply -f 01-nfs.yaml

kubectl apply -f 02-docker.yaml

kubectl apply -f 03-utils.yaml

kubectl -n ledger wait --for=condition=complete --timeout=600s job/utils

kubectl apply -f 04-orderer.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/orderer

kubectl apply -f 05-ca.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/ca

kubectl apply -f 06-org1.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/org1peer1

kubectl apply -f 07-org2.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/org2peer1

kubectl apply -f 08-org3.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/org3peer1

kubectl apply -f 09-org4.yaml

kubectl -n ledger wait --for=condition=available --timeout=600s deployment/org4peer1

kubectl apply -f 10-create-channel.yaml

kubectl -n ledger wait --for=condition=complete --timeout=600s job/create-channel

kubectl apply -f 11-join-channel.yaml

kubectl -n ledger wait --for=condition=complete --timeout=600s job/join-channel

kubectl apply -f 12-install-chaincode.yaml

kubectl -n ledger wait --for=condition=complete --timeout=600s job/chaincode-install

kubectl apply -f 13-instantiate-chaincode.yaml

kubectl -n ledger wait --for=condition=complete --timeout=600s job/chaincode-instantiate

kubectl apply -f 14-application.yaml
