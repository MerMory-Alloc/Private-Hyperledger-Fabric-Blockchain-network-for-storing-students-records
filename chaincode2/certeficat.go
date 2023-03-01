package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Certeficat
type SmartContract struct {
	contractapi.Contract
}

// Certeficat describes basic details of what makes up a simple certeficat
//Insert struct field in alphabetic order => to achieve determinism accross languages
// golang keeps the order when marshal to json but doesn't order automatically
type Certeficat struct {
	StudentID          				 string `json:"StudentID"`
	ID             						 string `json:"ID"`
	Title          						 string `json:"Title"`
	IssuePlace          			 string `json:"IssuePlace"`
	IssueDate          				 string `json:"IssueDate"`
	DocHash          				 	 string `json:"DocHash"`
	Semester1Average 					 int    `json:"Semester1Average"`
	Semester2Average           int    `json:"Semester2Average"`
	Semester3Average           int    `json:"Semester2Average"`
	Semester4Average           int    `json:"Semester2Average"`
	Semester5Average           int    `json:"Semester2Average"`
	Semester6Average           int    `json:"Semester2Average"`
	YearsAverage           		 int    `json:"YearsAverage"`
	Description        				 string `json:"Description"`
}

// InitLedger adds a base set of certeficats to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

// CreateCerteficat issues a new certeficat to the world state with given details.
func (s *SmartContract) CreateCerteficat(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, issuedate string, DocHash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, yearsaverage int, description string) error {
	exists, err := s.CerteficatExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the certeficat %s already exists", id)
	}

	certeficat := Certeficat{
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
	certeficatJSON, err := json.Marshal(certeficat)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, certeficatJSON)
}

// ReadCerteficat returns the certeficat stored in the world state with given id.
func (s *SmartContract) ReadCerteficat(ctx contractapi.TransactionContextInterface, id string) (*Certeficat, error) {
	certeficatJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if certeficatJSON == nil {
		return nil, fmt.Errorf("the certeficat %s does not exist", id)
	}

	var certeficat Certeficat
	err = json.Unmarshal(certeficatJSON, &certeficat)
	if err != nil {
		return nil, err
	}

	return &certeficat, nil
}

// UpdateCerteficat updates an existing certeficat in the world state with provided parameters.
func (s *SmartContract) UpdateCerteficat(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, issuedate string, DocHash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, yearsaverage int, description string) error {
	exists, err := s.CerteficatExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the certeficat %s does not exist", id)
	}

	// overwriting original certeficat with new certeficat
	certeficat := Certeficat{
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
	certeficatJSON, err := json.Marshal(certeficat)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, certeficatJSON)
}

// DeleteCerteficat deletes an given certeficat from the world state.
func (s *SmartContract) DeleteCerteficat(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.CerteficatExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the certeficat %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// CerteficatExists returns true when certeficat with given ID exists in world state
func (s *SmartContract) CerteficatExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	certeficatJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return certeficatJSON != nil, nil
}

// GetAllCerteficats returns all certeficats found in world state
func (s *SmartContract) GetAllCerteficats(ctx contractapi.TransactionContextInterface) ([]*Certeficat, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all certeficats in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var certeficats []*Certeficat
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var certeficat Certeficat
		err = json.Unmarshal(queryResponse.Value, &certeficat)
		if err != nil {
			return nil, err
		}
		certeficats = append(certeficats, &certeficat)
	}

	return certeficats, nil
}

func main() {
	certeficatChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating student chaincode: %v", err)
	}

	if err := certeficatChaincode.Start(); err != nil {
		log.Panicf("Error starting student chaincode: %v", err)
	}
}
