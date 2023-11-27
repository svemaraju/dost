package internal

import (
	"testing"
)


func TestPasswordLength(t *testing.T) {
	password, err := GeneratePassword(10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else {
		if len(password) != 10 {
			t.Errorf("Expected 10 character password, got %d", len(password))
		}
	}
}