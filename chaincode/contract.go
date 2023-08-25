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
func (s *SmartContract) RegisterVoter(ctx contractapi.TransactionContextInterface, voter Voter) error {

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

// RegisterCandidate registers candidate to world state. candidate is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterCandidate(ctx contractapi.TransactionContextInterface, candidate Candidate) error {
	extisting, err := ctx.GetStub().GetState(Key(candidate))
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)

	}
	if extisting != nil {
		return errors.New("candidate already registered")
	}
	candidateBytes, err := json.Marshal(candidate)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(Key(candidate), candidateBytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}

	return nil
}

// RegisterElection registers election to world state. election is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterElection(ctx contractapi.TransactionContextInterface, election Election) error {
	extisting, err := ctx.GetStub().GetState(fmt.Sprint(election.ElectionYear))
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)

	}
	if extisting != nil {
		return errors.New("election already registered")
	}
	electionBytes, err := json.Marshal(election)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(fmt.Sprint(election.ElectionYear), electionBytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}

	return nil
}

// RegisterPoliticalParty registers politicalparty to world state. politicalparty is
// expected to be prepared by the backend.
func (s *SmartContract) RegisterPoliticalParty(ctx contractapi.TransactionContextInterface, politicalParty PoliticalParty) error {
	extisting, err := ctx.GetStub().GetState(politicalParty.PartyID)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)

	}
	if extisting != nil {
		return errors.New("politicalParty already registered")
	}
	politicalPartyBytes, err := json.Marshal(politicalParty)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(politicalParty.PartyID, politicalPartyBytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}

	return nil
}
