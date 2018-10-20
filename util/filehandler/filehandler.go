package filehandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/mortenterhart/trivial-tickets/structs"
)

/*
*
* Matrikelnummern
* 3040018
* 6694964
* 3478222
 */

// ReadUserFile takes a string as parameter for the location
// of the users.json file, reads the content and stores it inside
// of the hashmap for the users
func ReadUserFile(src string, users *map[string]structs.User) {

	// Read contents of users.json
	fileContent, errReadFile := ioutil.ReadFile(src)

	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	// Unmarshal into users hashmap
	errUnmarshal := json.Unmarshal(fileContent, users)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
}

// WriteUserFile writes the contents of the users map to the
// file system to persist any changes
func WriteUserFile(dest string, users *map[string]structs.User) error {

	// Create json from the hashmap
	usersMarshal, _ := json.MarshalIndent(users, "", "   ")

	// Write json to file
	return ioutil.WriteFile(dest, usersMarshal, 0644)
}

// CreateFile writes a given ticket to a given path in the json format
func CreateFile(path string, ticket *structs.Ticket) error {

	// TODO: If already exists, overwrite contents

	// Check if path exists. If not, create it
	if _, errExists := os.Stat(path); os.IsNotExist(errExists) {

		errCreateFolders := CreateFolders(path)
		if errCreateFolders != nil {
			log.Fatal(errCreateFolders)
		}
	}

	// Encode the struct with json
	marshalTicket, errMarshalTicket := json.Marshal(ticket)

	if errMarshalTicket != nil {
		log.Fatal(errMarshalTicket)
	}

	// Create the final output path
	finalPath := path + "/" + strconv.Itoa(int(ticket.Id)) + ".json"

	// If the file already exists, return an error. Otherwise write the file
	if _, errExistsFile := os.Stat(finalPath); os.IsNotExist(errExistsFile) {
		return ioutil.WriteFile(finalPath, marshalTicket, 0644)
	} else {
		return errors.New("File already exists on disk.\nPath: " + finalPath)
	}
}

// CreateFolders creates the folders specified in the parameter
func CreateFolders(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}