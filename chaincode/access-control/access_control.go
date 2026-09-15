package main

// Access-control chaincode — ports AccessControl.sol logic.
// Functions to implement: LogDecision, GetLogCount, GetLog.
// Must be append-only: no update/delete function should exist.
// Owner: Jeswin

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type AccessControlContract struct {
	contractapi.Contract
}

// TODO: implement LogDecision(ctx, assetID, requesterDID string, riskScore int, allowed bool) error
// TODO: implement GetLogCount(ctx) (int, error)
// TODO: implement GetLog(ctx, index int) (string, error)

func main() {
	chaincode, err := contractapi.NewChaincode(&AccessControlContract{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
