package main

// Asset registry chaincode — ports AssetRegistry.sol logic.
// Functions to implement: MintAsset, GetAsset, TransferAsset.
// Restricted-classification minting is gated by the endorsement policy,
// not by code in this chaincode.
// Owner: Avikrit

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type AssetContract struct {
	contractapi.Contract
}

// TODO: implement MintAsset(ctx, id, name, classification, ownerDID, metadataHash string) error
// TODO: implement GetAsset(ctx, id string) (string, error)
// TODO: implement TransferAsset(ctx, id, newOwnerDID string) error

func main() {
	chaincode, err := contractapi.NewChaincode(&AssetContract{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
