package internal

import (
	"strings"
	"testing"
)

func TestPasswordLength(t *testing.T) {
	password, err := generatePassword(10, true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else {
		if len(password) != 10 {
			t.Errorf("Expected 10 character password, got %d", len(password))
		}
	}
}

func TestSpecialCharacter10K(t *testing.T) {
	splCharMissing := 0
	for i := 1; i <= 10000; i++ {
		password, err := generatePassword(10, false)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		} else {
			if !strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:'\",.<>/?") {
				// t.Errorf("Does not contain special characters %s", password)
				splCharMissing++
			}
		}
	}
	if splCharMissing > 0 {
		t.Errorf("Special character was missing in %d / 10000 instances.", splCharMissing)
	}
}
