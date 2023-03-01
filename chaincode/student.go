package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Student
type SmartContract struct {
	contractapi.Contract
}


type Student struct {
	EnrollementYear 	 int    			`json:"EnrollementYear"`
	FullName           string 			`json:"FullName"`
	ID                 string 			`json:"ID"`
	Speciality         string 			`json:"Speciality"`
	Major          		 string 			`json:"Major"`
	CurrentYear    		 int    			`json:"CurrentYear"`
	Degree         		 string       `json:"Degree"`
	University     		 string 			`json:"University"`
	Faculty            string 			`json:"Faculty"`
	DateOfBirth        string 			`json:"DateOfBirth"`
	PlaceOfBirth       string 			`json:"PlaceOfBirth"`
	Address            string 			`json:"Address"`
}

// InitLedger adds a base set of students to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	students := []Student{
		{ID: "student1", FullName: "Mohamed", DateOfBirth: "12/06/1997", PlaceOfBirth: "Chelef", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2018, CurrentYear: 2, Speciality: "Math", Major: "Probabilite", Degree: "Master"},
		{ID: "student2", FullName: "Said", DateOfBirth: "12/12/2000", PlaceOfBirth: "Oran", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2020, CurrentYear: 2, Speciality: "Info", Major: "", Degree: "Licence"},
		{ID: "student3", FullName: "Israa", DateOfBirth: "12/07/1999", PlaceOfBirth: "Ain Temouchent", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2019, CurrentYear: 3, Speciality: "Math", Major: "", Degree: "Licence"},
		{ID: "student4", FullName: "Ali", DateOfBirth: "12/08/2001", PlaceOfBirth: "Oran", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2019, CurrentYear: 3, Speciality: "Info", Major: "SI", Degree: "Licence"},
		{ID: "student5", FullName: "Jenice", DateOfBirth: "12/02/2000", PlaceOfBirth: "Adrar", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2021, CurrentYear: 1, Speciality: "Math/Info", Major: "", Degree: "Licence"},
		{ID: "student6", FullName: "Kalthoum", DateOfBirth: "12/12/1995", PlaceOfBirth: "Bechar", Address: "Sidi Bel abbes,Tassala", University: "Djillali Liabes", Faculty: "Sciences Exactes", EnrollementYear: 2018, CurrentYear: 2, Speciality: "Info", Major: "RSSI", Degree: "Master"},
	}

	for _, student := range students {
		studentJSON, err := json.Marshal(student)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(student.ID, studentJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	return nil
}

// CreateStudent issues a new student to the world state with given details.
func (s *SmartContract) CreateStudent(ctx contractapi.TransactionContextInterface, id string, fullname string, currentyear int, speciality string, enrollementyear int, major string, degree string, university string, faculty string, dateofbirth string,placeofbirth string, adress string) error {
	exists, err := s.StudentExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the student %s already exists", id)
	}

	student := Student{
		ID:                    id,
		FullName:              fullname,
		CurrentYear:           currentyear,
		Speciality:            speciality,
		EnrollementYear:       enrollementyear,
		Major:								 major,
		Degree:								 degree,
		University:						 university,
		Faculty:							 faculty,
		DateOfBirth:					 dateofbirth,
		PlaceOfBirth:					 placeofbirth,
		Address:							 adress,
	}
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, studentJSON)
}

// ReadStudent returns the student stored in the world state with given id.
func (s *SmartContract) ReadStudent(ctx contractapi.TransactionContextInterface, id string) (*Student, error) {
	studentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if studentJSON == nil {
		return nil, fmt.Errorf("the student %s does not exist", id)
	}

	var student Student
	err = json.Unmarshal(studentJSON, &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// UpdateStudent updates an existing student in the world state with provided parameters.
func (s *SmartContract) UpdateStudent(ctx contractapi.TransactionContextInterface, id string, fullname string, currentyear int, speciality string, enrollementyear int, major string, degree string, university string, faculty string,dateofbirth string,placeofbirth string, adress string) error {
	exists, err := s.StudentExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the student %s does not exist", id)
	}

	// overwriting original student with new student
	student := Student{
		ID:                    id,
		FullName:              fullname,
		CurrentYear:           currentyear,
		Speciality:            speciality,
		EnrollementYear:       enrollementyear,
		Major:								 major,
		Degree:								 degree,
		University:						 university,
		Faculty:							 faculty,
		DateOfBirth:					 dateofbirth,
		PlaceOfBirth:					 placeofbirth,
		Address:							 adress,
	}
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, studentJSON)
}

// DeleteStudent deletes an given student from the world state.
func (s *SmartContract) DeleteStudent(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.StudentExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the student %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// StudentExists returns true when student with given ID exists in world state
func (s *SmartContract) StudentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	studentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return studentJSON != nil, nil
}

// TransferStudent updates the University field of student with given id in world state, and returns the old University.
func (s *SmartContract) TransferStudent(ctx contractapi.TransactionContextInterface, id string, newUniversity string) (string, error) {
	student, err := s.ReadStudent(ctx, id)
	if err != nil {
		return "", err
	}

	oldUniversity := student.University
	student.University = newUniversity

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, studentJSON)
	if err != nil {
		return "", err
	}

	return oldUniversity, nil
}

//ChangeMajor updates the Major field of student with given id in world state, and returns the old Major
func (s *SmartContract) ChangeMajor(ctx contractapi.TransactionContextInterface, id string, newMajor string) (string, error) {
	student, err := s.ReadStudent(ctx, id)
	if err != nil {
		return "", err
	}

	oldMajor := student.Major
	student.Major = newMajor

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, studentJSON)
	if err != nil {
		return "", err
	}

	return oldMajor, nil
}

//ChangeDegree updates the Degree field of student with given id in world state, and returns the old Degree
func (s *SmartContract) ChangeDegree(ctx contractapi.TransactionContextInterface, id string, newDegree string) (string, error) {
	student, err := s.ReadStudent(ctx, id)
	if err != nil {
		return "", err
	}

	oldDegree := student.Degree
	student.Degree = newDegree

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, studentJSON)
	if err != nil {
		return "", err
	}

	return oldDegree, nil
}


//ChangeCurrentYear updates the CurrentYear field of student with given id in world state, and returns the old CurrentYear
func (s *SmartContract) ChangeCurrentYear(ctx contractapi.TransactionContextInterface, id string, newCurrentYear int) (int, error) {
	student, err := s.ReadStudent(ctx, id)
	if err != nil {
		return -1, err
	}

	oldCurrentYear := student.CurrentYear
	student.CurrentYear = newCurrentYear

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return -1, err
	}

	err = ctx.GetStub().PutState(id, studentJSON)
	if err != nil {
		return -1, err
	}

	return oldCurrentYear, nil
}

// GetAllStudents returns all students found in world state
func (s *SmartContract) GetAllStudents(ctx contractapi.TransactionContextInterface) ([]*Student, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all students in the chaincode fullnamespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var students []*Student
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var student Student
		err = json.Unmarshal(queryResponse.Value, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func main() {
	studentChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating student chaincode: %v", err)
	}

	if err := studentChaincode.Start(); err != nil {
		log.Panicf("Error starting student chaincode: %v", err)
	}
}
