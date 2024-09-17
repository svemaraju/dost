package internal

import (
	"strings"
	"testing"
)


func reportError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPasswordLength(t *testing.T) {
	password, err := generatePassword(10, true)
	reportError(err, t)
	if len(password) != 10 {
		t.Errorf("Expected 10 character password, got %d", len(password))
	}
}

func TestSpecialCharacter10K(t *testing.T) {
	splCharMissing := 0
	for i := 1; i <= 10000; i++ {
		password, err := generatePassword(10, false)
		reportError(err, t)
		if !strings.ContainsAny(password, specialChars) {
			splCharMissing++
		}
	}
	if splCharMissing > 0 {
		t.Errorf("Special character was missing in %d / 10000 instances.", splCharMissing)
	}
}


func TestSelectingRandomCharacter(t *testing.T) {
	characterSet := "abdexptw"
	randomCharacter:= selectRandomCharacter(characterSet)
	if !strings.Contains(characterSet, randomCharacter) {
		t.Errorf("Failed: Expected randomCharacter (%s) to be part of characterSet (%s):", 
				 randomCharacter, characterSet)
	}

}