# Hyperledger blockchain

Theres a lot of good information on how this is deployed found in the repo here:

https://github.com/IBM/blockchain-network-on-kubernetes

IBM are one of the big contributors to blockchain tech.

# Let's build it

To build just make sure your kubectl config is set to the correct GKE cluster and run the following:

```
./build.sh
```

If you want to remove after then you can also run:

```
./delete.sh
```

# Give it a go

After the deployment you can jump on a peer node and query it for a book in the list by running the following:

```
peer chaincode invoke -o orderer:31010 -C library -n bookstore -c '{"Args":["QueryBook","9781841499895"]}'
```

Response should look something like the following:

```
2021-02-24 17:39:56.152 GMT [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 001 Chaincode invoke successful. result: status:200 payload:"{\"title\":\"The Expanse: Leviathan Wakes\",\"author\":\"James S. A. Corey\",\"description\":\"As close as you'll get to a hollywood blockbuster in book form.\",\"isbn\":\"9781841499895\"}"
```

# Issues

Created the following issue on the ticket here, worth looking into:

https://jira.hyperledger.org/browse/FABG-1033
