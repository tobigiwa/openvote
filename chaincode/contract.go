package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ITYPES is a union set type constraint
// that enforces only allowable types are passed to smartcontract methods.
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
	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(key, bytes)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	return nil
}

// QueryVoter querys voter from world state.
func (s *SmartContract) QueryVoter(ctx contractapi.TransactionContextInterface, key string) (Voter, error) {
	return queryFunc[Voter](ctx, key)
}

// QueryCandidate querys candidate from world state.
func (s *SmartContract) QueryCandidate(ctx contractapi.TransactionContextInterface, key string) (Candidate, error) {
	return queryFunc[Candidate](ctx, key)
}

// QueryElection querys election from world state.
func (s *SmartContract) QueryElection(ctx contractapi.TransactionContextInterface, key string) (Election, error) {
	return queryFunc[Election](ctx, key)
}

// QueryPoliticalParty querys politicalparty from world state.
func (s *SmartContract) QueryPoliticalParty(ctx contractapi.TransactionContextInterface, key string) (PoliticalParty, error) {
	return queryFunc[PoliticalParty](ctx, key)
}

func queryFunc[T ITYPES](ctx contractapi.TransactionContextInterface, key string) (T, error) {

	var zeroValueOfITYPES T
	var result T
	extisting, err := ctx.GetStub().GetState(key)
	if err != nil {
		return zeroValueOfITYPES, fmt.Errorf("unable to interact with world state: %v", err)
	}
	if extisting == nil {
		return zeroValueOfITYPES, fmt.Errorf("Cannot read world state with key %s. Does not exist", key)
	}
	err = json.Unmarshal(extisting, result)
	if err != nil {
		return zeroValueOfITYPES, err
	}
	return result, nil
}

// updateVoter updates voter from world state. voter is
// expected to be prepared by the backend.
func (s *SmartContract) UpdateVoter(ctx contractapi.TransactionContextInterface, voter Voter) error {
	return updateFunc[Voter](ctx, Key(voter), voter)
}

// updateCandidate updates candidate from world state. candidate is
// expected to be prepared by the backend.
func (s *SmartContract) UpdateCandidate(ctx contractapi.TransactionContextInterface, key string, candidate Candidate) error {
	return updateFunc[Candidate](ctx, Key(candidate), candidate)
}

// updateElection updates election from world state. election is
// expected to be prepared by the backend.
func (s *SmartContract) UpdateElection(ctx contractapi.TransactionContextInterface, key string, election Election) error {
	return updateFunc[Election](ctx, fmt.Sprint(election.ElectionYear), election)
}

// updatePoliticalParty updates politicalparty from world state. politicalparty is
// expected to be prepared by the backend.
func (s *SmartContract) UpdatePoliticalParty(ctx contractapi.TransactionContextInterface, politicalParty PoliticalParty) error {
	return updateFunc[PoliticalParty](ctx, politicalParty.PartyID, politicalParty)
}

func updateFunc[T ITYPES](ctx contractapi.TransactionContextInterface, key string, update T) error {
	extisting, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	if extisting == nil {
		return fmt.Errorf("Cannot read world state with key %s. Does not exist", key)
	}
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(key, data)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	return nil
}

// DeleteVoter deletes voter from world state.
func (s *SmartContract) DeleteVoter(ctx contractapi.TransactionContextInterface, key string) error {
	return deleteFunc[Voter](ctx, key)
}

// DeleteCandidate deletes candidate from world state.
func (s *SmartContract) DeleteCandidate(ctx contractapi.TransactionContextInterface, key string) error {
	return deleteFunc[Candidate](ctx, key)
}

// DeleteElection deletes election from world state.
func (s *SmartContract) DeleteElection(ctx contractapi.TransactionContextInterface, key string) error {
	return deleteFunc[Election](ctx, key)
}

// DeletePoliticalParty deletes politicalparty from world state.
func (s *SmartContract) DeletePoliticalParty(ctx contractapi.TransactionContextInterface, key string) error {
	return deleteFunc[PoliticalParty](ctx, key)
}

func deleteFunc[T ITYPES](ctx contractapi.TransactionContextInterface, key string) error {

	extisting, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	if extisting == nil {
		return fmt.Errorf("Cannot read world state with key %s. Does not exist", key)
	}
	err = ctx.GetStub().DelState(key)
	if err != nil {
		return fmt.Errorf("unable to interact with world state: %v", err)
	}
	return nil
}

// QueryAllVoter queryAlls voter from world state.Any error
// encounter would return the function with a nil slice.
func (s *SmartContract) QueryAllVoter(ctx contractapi.TransactionContextInterface) ([]Voter, error) {
	return queryAllFunc[Voter](ctx)
}

// QueryAllCandidate queryAlls candidate from world state.Any error
// encounter would return the function with a nil slice.
func (s *SmartContract) QueryAllCandidate(ctx contractapi.TransactionContextInterface) ([]Candidate, error) {
	return queryAllFunc[Candidate](ctx)
}

// QueryAllElection queryAlls election from world state. Any error
// encounter would return the function with a nil slice.
func (s *SmartContract) QueryAllElection(ctx contractapi.TransactionContextInterface) ([]Election, error) {
	return queryAllFunc[Election](ctx)
}

// QueryAllPoliticalParty queryAlls politicalparty from world state. Any error
// encounter would return the function with a nil slice.
func (s *SmartContract) QueryAllPoliticalParty(ctx contractapi.TransactionContextInterface) ([]PoliticalParty, error) {
	return queryAllFunc[PoliticalParty](ctx)
}

func queryAllFunc[T ITYPES](ctx contractapi.TransactionContextInterface) ([]T, error) {
	startkey := ""
	endKey := ""
	container := []T{}
	resultIterator, err := ctx.GetStub().GetStateByRange(startkey, endKey)
	if err != nil {
		// it is unclear from the function what err here denotes, I'll assume no data was fetched, so return
		return nil, err
	}
	defer resultIterator.Close()
	for resultIterator.HasNext() {
		var result T
		queryResponse, err := resultIterator.Next()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(queryResponse.Value, result)
		if err != nil {
			return nil, err
		}
		container = append(container, result)
	}
	return container, nil
}
