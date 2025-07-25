package internal

import (
	"strings"
	"testing"
)


func reportError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGeneratePassword(t *testing.T){
	tests := []struct {
		name string
		length int
		expectError bool
		expectSpecial bool
		excludeSpecial bool
	} {
		{
			name: "Valid password of 10 chars with special characters",
			length: 10,
			expectError: false,
			expectSpecial: true,
			excludeSpecial: false,
		},
		{
			name: "Valid password of 20 chars without forcing special characters",
			length: 20,
			expectError: false,
			expectSpecial: false,
			excludeSpecial: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := generatePassword(tt.length, tt.excludeSpecial)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
				return
			}

			if len(password) != tt.length {
				t.Errorf("Expected password length %d, got %d", tt.length, len(password))
			}

			if tt.expectSpecial && !strings.ContainsAny(password, specialChars) {
				t.Errorf("Expected password to include a special character but got none: %s", password)
			}
		})
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
