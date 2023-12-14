package cmd

import (
	"testing"
)

func TestGetCommand(t *testing.T) {
	// Test valid command name and option
	command, err := GetCommand("net", "todo")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if command.Name != "ToDo" {
		t.Errorf("Expected command name: ToDo, got: %v", command.Name)
	}

	// Test with invalid command name
	command, err = GetCommand("invalid", "option")
	if err == nil {
		t.Errorf("Expected error, got: %v", command)
	}

	// Test with invalid command option
	command, err = GetCommand("net", "invalid")
	if err == nil {
		t.Errorf("Expected error, got: %v", command)
	}

	// Test with no command name and option
	command, err = GetCommand("", "")
	if err == nil {
		t.Errorf("Expected error, got: %v", command)
	}
}
