package internal

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func getRandomPath() string {
	path := os.Getenv("HOME") + "/.dost-" + strconv.Itoa(rand.Intn(500))
	fmt.Printf("--- [Test path for dost: %s]\n", path)
	return path
}

func cleanUpPath(path string) {
	os.RemoveAll(path)
}

func TestStorageInitReady(t *testing.T) {
	// random path for testing purpose
	path := getRandomPath()
	// clean up later
	defer cleanUpPath(path)

	storage := GetStorage(path)

	storage.Init()

	if !storage.IsReady() {
		t.Errorf("Expected directory: %s to be created on storage.Init()", path)
	}

}

func TestStorageAddShow(t *testing.T) {
	// random path for testing purpose
	path := getRandomPath()
	// clean up later
	defer cleanUpPath(path)

	storage := GetStorage(path)
	storage.Init()

	identifier := "email/sri@example.com"
	password := "someRandomPassword"

	addErr := storage.Add(password, identifier)

	if addErr != nil {
		t.Errorf("Did not expect an error when calling storage.Add: \n%v", addErr)
	}

	// check if password exists on the file
	passwordFromFile, showErr := storage.Show(identifier)
	if showErr != nil {
		t.Errorf("Did not expect an error when calling storage.Show: \n%v", showErr)
	}

	if passwordFromFile != password {
		t.Errorf("Password that was added did not match the one from the one that got saved\npassword: %s, passwordFromFile: %s", 
				password, passwordFromFile)
	}


}


// TestAddDuplicate verifies adding a new password with existing identifier should raise an error
func TestAddDuplicate(t *testing.T) {
	// random path for testing purpose
	path := getRandomPath()
	// clean up later
	defer cleanUpPath(path)

	storage := GetStorage(path)
	storage.Init()

	identifier := "email/sri@example.com"
	password := "someRandomPassword"

	addErr1 := storage.Add(password, identifier)
	
	if addErr1 != nil {
		t.Errorf("Did not expect an error when calling storage.Add: \n%v", addErr1)
	}

	// try saving the new password with same identifier
	addErr2 := storage.Add("someNewRandomPassword", identifier)

	if addErr2 == nil {
		t.Errorf("Error should have been raised when add new password to existing indetifier.")
	}
}

// TestShowNonExisting - should raise error
func TestShowNonExisting(t *testing.T) {
	// random path for testing purpose
	path := getRandomPath()
	// clean up later
	defer cleanUpPath(path)

	storage := GetStorage(path)
	storage.Init()

	_, err := storage.Show("someNonExisting/Identifier@email.com")

	if err == nil {
		t.Errorf("Error should have been raised when non-existing identifier is being requested to be shown")
	}

}

// TestEmptyIdentifier
func TestEmptyIdentifier(t *testing.T) {
	// random path for testing purpose
	path := getRandomPath()
	// clean up later
	defer cleanUpPath(path)

	storage := GetStorage(path)
	storage.Init()

	identifier := ""
	password := "someRandomPassword"

	err := storage.Add(password, identifier)

	if err == nil {
		t.Errorf("Error should have been raised when identifier is empty string.")
	}
}
