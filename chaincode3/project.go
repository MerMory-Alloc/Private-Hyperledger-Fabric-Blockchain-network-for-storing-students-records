package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Project
type SmartContract struct {
	contractapi.Contract
}

// Project describes basic details of what makes up a simple project
//Insert struct field in alphabetic order => to achieve determinism accross languages
// golang keeps the order when marshal to json but doesn't order automatically
type Project struct {
	StudentID          string 	 `json:"StudentID"`
	ID             		 string 	 `json:"ID"`
	Title          		 string 	 `json:"Title"`
	Advisor          	 string 	 `json:"Advisor"`
	Course          	 string 	 `json:"Course"`
	Type          	 	 string 	 `json:"Type"`
	DeposeDate         string 	 `json:"DeposeDate"`
	RepoHash           string 	 `json:"RepoHash"`
	UrlPath            string 	 `json:"UrlPath"`
	Partner1					 string    `json:"Partner1"`
	Partner2           string    `json:"Partner2"`
	Partner3           string    `json:"Partner3"`
	Partner4           string    `json:"Partner4"`
	partner5           string    `json:"partner5"`
	Partner6           string    `json:"Partner6"`
	Grade           	 int    	 `json:"Grade"`
	Description        string 	 `json:"Description"`
}

// InitLedger adds a base set of projects to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

// CreateProject issues a new project to the world state with given details.
func (s *SmartContract) CreateProject(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, deposedate string, repohash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, grade int, description string) error {
	exists, err := s.ProjectExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the project %s already exists", id)
	}

	project := Project{
		ID:             						id,
		StudentID:          				studentid,
		Title:          						title,
		IssuePlace:                 issueplace,
		DeposeDate:                  deposedate,
		RepoHash:                    repohash,
		Semester1Average: 					semester1average,
		Semester2Average:           semester2average,
		Semester3Average:           semester3average,
		Semester4Average:           semester4average,
		Semester5Average:           semester5average,
		Semester6Average:           semester6average,
		Grade:               grade,
		Description:                description,
	}
	projectJSON, err := json.Marshal(project)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, projectJSON)
}

// ReadProject returns the project stored in the world state with given id.
func (s *SmartContract) ReadProject(ctx contractapi.TransactionContextInterface, id string) (*Project, error) {
	projectJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if projectJSON == nil {
		return nil, fmt.Errorf("the project %s does not exist", id)
	}

	var project Project
	err = json.Unmarshal(projectJSON, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject updates an existing project in the world state with provided parameters.
func (s *SmartContract) UpdateProject(ctx contractapi.TransactionContextInterface, id string, studentid string,
	  title string,  issueplace string, deposedate string, repohash string, semester1average int, semester2average int,
		  semester3average int,  semester4average int, semester5average int,  semester6average int, grade int, description string) error {
	exists, err := s.ProjectExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the project %s does not exist", id)
	}

	// overwriting original project with new project
	project := Project{
		ID:             						id,
		StudentID:          				studentid,
		Title:          						title,
		IssuePlace:                 issueplace,
		DeposeDate:                  deposedate,
		RepoHash:                    repohash,
		Semester1Average: 					semester1average,
		Semester2Average:           semester2average,
		Semester3Average:           semester3average,
		Semester4Average:           semester4average,
		Semester5Average:           semester5average,
		Semester6Average:           semester6average,
		Grade:               grade,
		Description:                description,
	}
	projectJSON, err := json.Marshal(project)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, projectJSON)
}

// DeleteProject deletes an given project from the world state.
func (s *SmartContract) DeleteProject(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.ProjectExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the project %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// ProjectExists returns true when project with given ID exists in world state
func (s *SmartContract) ProjectExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	projectJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return projectJSON != nil, nil
}

// GetAllProjects returns all projects found in world state
func (s *SmartContract) GetAllProjects(ctx contractapi.TransactionContextInterface) ([]*Project, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all projects in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var projects []*Project
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var project Project
		err = json.Unmarshal(queryResponse.Value, &project)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}

func main() {
	projectChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating student chaincode: %v", err)
	}

	if err := projectChaincode.Start(); err != nil {
		log.Panicf("Error starting student chaincode: %v", err)
	}
}
