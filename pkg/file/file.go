package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// FileMethods contains all the methods associated with manipulating OS files
type FileMethods interface {
	ReadUsersFile() map[string]interface{}
	WriteUsersFile()
	ReadDataFile() map[string]interface{}
	WriteDataFile()
}

func dataFilePathNode(nodePID int) (string, error) {
	id := strconv.Itoa(nodePID)
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dataFilePath := filepath.Join(cwd, "nodeStorage", "node"+id, "data.json")
	return dataFilePath, nil
}

// ReadUsersFile returns all the information from the global users.json file
// Code adapted from: https://tutorialedge.net/golang/parsing-json-with-golang/
func ReadUsersFile() (map[string]interface{}, error) {

	cwd, err := os.Getwd()
	fmt.Println(cwd)
	if err != nil {
		return nil, err
	}
	userFilePath := filepath.Join(cwd, "users.json")
	jsonFile, osErr := os.Open(userFilePath)

	if osErr != nil { // if we os.Open returns an error then handle it
		return nil, osErr
	}

	defer jsonFile.Close()

	byteValue, readAllError := ioutil.ReadAll(jsonFile)
	if readAllError != nil {
		return nil, readAllError
	}

	var fileContents map[string]interface{}

	// Unmarshal parses the byteValue array to a type defined by fileContents
	marshalError := json.Unmarshal([]byte(byteValue), &fileContents)
	if marshalError != nil { // if we os.Open returns an error then handle it
		return nil, marshalError
	}
	return fileContents, nil
}

// ReadDataFile returns all the information from the data.json of the respective node's local file
func ReadDataFile(pid int) (map[string]interface{}, error) {

	filePath, err := dataFilePathNode(pid)
	if err != nil {
		return nil, err
	}
	jsonFile, osErr := os.Open(filePath)

	if osErr != nil {
		return nil, osErr
	}

	defer jsonFile.Close()

	byteValue, readAllError := ioutil.ReadAll(jsonFile)
	if readAllError != nil {
		return nil, readAllError
	}

	var fileContents map[string]interface{}
	marshalError := json.Unmarshal([]byte(byteValue), &fileContents)
	if marshalError != nil { // if we os.Open returns an error then handle it
		return nil, marshalError
	}

	return fileContents, nil
}

// WriteUsersFile takes updates the user file with the new users provided
func WriteUsersFile(addUsers map[string]interface{}) error {

	originalFileContent, readError := ReadUsersFile()

	if readError != nil {
		return readError
	}

	// Update the values from the file
	for key, value := range addUsers {
		originalFileContent[key] = value
	}

	file, marshallError := json.MarshalIndent(originalFileContent, "", " ")
	if marshallError != nil {
		return marshallError

	}

	cwd, err := os.Getwd()
	fmt.Println(cwd)
	if err != nil {
		return err
	}
	userFilePath := filepath.Join(cwd, "users.json")
	var writeError = ioutil.WriteFile(userFilePath, file, 0644)
	if writeError != nil {
		return writeError
	}
	return nil
}

// WriteDataFile taks in the variable with map type then update the user file
func WriteDataFile(pid int, addData map[string]interface{}) error {

	filePath, err := dataFilePathNode(pid)
	if err != nil {
		return err
	}
	originalFileContent, readError := ReadDataFile(pid)

	if readError != nil {
		return readError
	}

	for key, value := range addData {
		originalFileContent[key] = value
	}

	file, marshallError := json.MarshalIndent(originalFileContent, "", " ")
	if marshallError != nil {
		return marshallError
	}

	var writeError = ioutil.WriteFile(filePath, file, 0644)
	if writeError != nil {
		return writeError
	}
	return nil
}

// OverwriteDatafromFile taks in the variable with map type then update the user file
func OverwriteDataFile(pid int, addData map[string]interface{}) error {

	filePath, err := dataFilePathNode(pid)

	if err != nil {
		return err
	}

	file, marshallError := json.MarshalIndent(addData, "", " ")
	if marshallError != nil {
		return marshallError
	}

	var writeError = ioutil.WriteFile(filePath, file, 0644)
	if writeError != nil {
		return writeError
	}
	return nil
}