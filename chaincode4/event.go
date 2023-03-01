package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Event
type SmartContract struct {
	contractapi.Contract
}

// Event describes basic details of what makes up a simple event
//Insert struct field in alphabetic order => to achieve determinism accross languages
// golang keeps the order when marshal to json but doesn't order automatically
type Event struct {
	StudentID          				 string `json:"StudentID"`
	ID             						 string `json:"ID"`
	Title          						 string `json:"Title"`
	Location          			 string `json:"Location"`
	JoinDate          				 string `json:"JoinDate"`
	Orgnaizer          				 	 string `json:"Orgnaizer"`
	Description        				 string `json:"Description"`
}

// InitLedger adds a base set of events to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

// CreateEvent issues a new event to the world state with given details.
func (s *SmartContract) CreateEvent(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, issuedate string, DocHash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, yearsaverage int, description string) error {
	exists, err := s.EventExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the event %s already exists", id)
	}

	event := Event{
		ID:             						id,
		StudentID:          				studentid,
		Title:          						title,
		IssuePlace:                 issueplace,
		IssueDate:                  issuedate,
		DocHash:                    dochash,
		Semester1Average: 					semester1average,
		Semester2Average:           semester2average,
		Semester3Average:           semester3average,
		Semester4Average:           semester4average,
		Semester5Average:           semester5average,
		Semester6Average:           semester6average,
		YearsAverage:               yearsaverage,
		Description:                description,
	}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, eventJSON)
}

// ReadEvent returns the event stored in the world state with given id.
func (s *SmartContract) ReadEvent(ctx contractapi.TransactionContextInterface, id string) (*Event, error) {
	eventJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if eventJSON == nil {
		return nil, fmt.Errorf("the event %s does not exist", id)
	}

	var event Event
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// UpdateEvent updates an existing event in the world state with provided parameters.
func (s *SmartContract) UpdateEvent(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, issuedate string, DocHash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, yearsaverage int, description string) error {
	exists, err := s.EventExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the event %s does not exist", id)
	}

	// overwriting original event with new event
	event := Event{
		ID:             						id,
		StudentID:          				studentid,
		Title:          						title,
		IssuePlace:                 issueplace,
		IssueDate:                  issuedate,
		DocHash:                    dochash,
		Semester1Average: 					semester1average,
		Semester2Average:           semester2average,
		Semester3Average:           semester3average,
		Semester4Average:           semester4average,
		Semester5Average:           semester5average,
		Semester6Average:           semester6average,
		YearsAverage:               yearsaverage,
		Description:                description,
	}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, eventJSON)
}

// DeleteEvent deletes an given event from the world state.
func (s *SmartContract) DeleteEvent(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.EventExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the event %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// EventExists returns true when event with given ID exists in world state
func (s *SmartContract) EventExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	eventJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return eventJSON != nil, nil
}

// GetAllEvents returns all events found in world state
func (s *SmartContract) GetAllEvents(ctx contractapi.TransactionContextInterface) ([]*Event, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all events in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var events []*Event
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var event Event
		err = json.Unmarshal(queryResponse.Value, &event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}

func main() {
	eventChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating student chaincode: %v", err)
	}

	if err := eventChaincode.Start(); err != nil {
		log.Panicf("Error starting student chaincode: %v", err)
	}
}
