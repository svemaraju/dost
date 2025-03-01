package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Storage struct {
	path string
}

func GetStorage(path string) *Storage {
	s := Storage{path: path}
	return &s
}

func (s *Storage) Init() {
	if err := os.Mkdir(s.path, 0755); err != nil {
		panic(err)
	}
}

func (s *Storage) Add(password, identifier string) error {
	directory := filepath.Dir(identifier)
    dirPath := filepath.Join(s.path, directory)
    // check if the directory specified for the password is present, 
    // if not then create it
	if _, err1 := os.Stat(dirPath); os.IsNotExist(err1) {
        log.Printf("%s does not exist, creating.", dirPath)
		if err3 := os.Mkdir(dirPath, 0755); err3 != nil {
            return err3
        } 
	}
    // write the password to the file
	absPath := filepath.Join(s.path, identifier)
	_, err2 := os.Stat(absPath)
	if os.IsNotExist(err2) {
		if err2 := os.WriteFile(absPath, []byte(password), 0600); err2 != nil {
			log.Println("Error writing file:", err2)
			return err2
		}
	} else {
		return fmt.Errorf("Storage.Add: identifier %s already in use, please use delete before add or use update", 
		identifier)
	}
	
	return nil

}

func (s *Storage) Show(identifier string) (string, error) {
    absPath := filepath.Join(s.path, identifier)
    password, err := os.ReadFile(absPath)
    if err != nil {
        return "", err
    }
    return string(password), nil
}

func (s *Storage) IsReady() bool {
	_, err := os.Stat(s.path)
	return !os.IsNotExist(err)
}
