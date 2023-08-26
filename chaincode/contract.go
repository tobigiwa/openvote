package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ITYPES enforces that only validated types are passed to smartcontract methods.
type ITYPES interface {
	Voter | Candidate | Election | PoliticalParty
}
type SmartContract struct {
	contractapi.Contract
}

// RegisterVoter registers voter to world state. voter is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterVoter(ctx contractapi.TransactionContextInterface, voter Voter) error {
	return registerFunc[Voter](ctx, Key(voter), voter, "voter")
}

// RegisterCandidate registers candidate to world state. candidate is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterCandidate(ctx contractapi.TransactionContextInterface, candidate Candidate) error {
	return registerFunc[Candidate](ctx, Key(candidate), candidate, "candidate")
}

// RegisterElection registers election to world state. election is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterElection(ctx contractapi.TransactionContextInterface, election Election) error {
	return registerFunc[Election](ctx, fmt.Sprint(election.ElectionYear), election, "election")
}

// RegisterPoliticalParty registers politicalparty to world state. politicalparty is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterPoliticalParty(ctx contractapi.TransactionContextInterface, politicalParty PoliticalParty) error {
	return registerFunc[PoliticalParty](ctx, politicalParty.PartyID, politicalParty, "politicalParty")
}

func registerFunc[T ITYPES](ctx contractapi.TransactionContextInterface, key string, body T, operation string) error {
	extisting, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	if extisting != nil {
		return fmt.Errorf("%s already registered", operation)
	}
	Bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(key, Bytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	return nil
}
