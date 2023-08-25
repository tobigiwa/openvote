package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// RegisterVoter registers voter to world state. voter is
// expected to be prepared by the backend.
func (s *SmartContract) ResgisterVoter(ctx contractapi.TransactionContextInterface, voter Voter) error {

	extisting, err := ctx.GetStub().GetState(Key(voter))
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)

	}
	if extisting != nil {
		return errors.New("voter already registered")
	}
	voterBytes, err := json.Marshal(voter)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(Key(voter), voterBytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}

	return nil
}


