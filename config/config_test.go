package config

import (
	"os"
	"testing"
)

// TestGetConfigFile tests the GetConfig function with a valid config file
func TestGetConfigFile(t *testing.T) {
	os.Setenv("CONFIG_FILE_NAME", "../test.yaml")
	file, err := os.Create("../test.yaml")
	if err != nil {
		t.Fatalf("failed to create test config file: %v", err)
	}
	defer file.Close()
	file.WriteString(`todo:
  url: "https://example.com"
  max_goroutine: 10
  max_todo: 100
  todo_id_start: 1
  todo_id_type: "int"
`)
	// GetConfig check result
	config, err := GetConfig()
	if err != nil {
		t.Errorf("GetConfig returned an error: %v", err)
	}
	// match the expected
	if config.Url != "https://example.com" {
		t.Errorf("config.Url is %s, want https://example.com", config.Url)
	}
	if config.MaxGoroutine != 10 {
		t.Errorf("config.MaxGoroutine is %d, want 10", config.MaxGoroutine)
	}
	if config.MaxTodo != 100 {
		t.Errorf("config.MaxTodo is %d, want 100", config.MaxTodo)
	}
	if config.TodoStartID != 1 {
		t.Errorf("config.TodoStartID is %d, want 1", config.TodoStartID)
	}
	if config.TodoIDType != "int" {
		t.Errorf("config.TodoIDType is %s, want int", config.TodoIDType)
	}
	// Remove file
	os.Remove("test.yaml")
}

// TestGetConfigEnv tests the GetConfig env vars
func TestGetConfigEnv(t *testing.T) {
	// Set env varis
	os.Setenv("CONFIG_FILE_NAME", "")
	os.Setenv("URL", "https://example.com")
	os.Setenv("MAX_GOROUTINE", "10")
	os.Setenv("MAX_TODO", "100")
	os.Setenv("TODO_ID_START", "1")
	os.Setenv("TODO_ID_TYPE", "int")
	// GetConfig check result
	config, err := GetConfig()
	if err != nil {
		t.Errorf("GetConfig returned an error: %v", err)
	}
	// match the expected vals
	if config.Url != "https://example.com" {
		t.Errorf("config.Url is %s, want https://example.com", config.Url)
	}
	if config.MaxGoroutine != 10 {
		t.Errorf("config.MaxGoroutine is %d, want 10", config.MaxGoroutine)
	}
	if config.MaxTodo != 100 {
		t.Errorf("config.MaxTodo is %d, want 100", config.MaxTodo)
	}
	if config.TodoStartID != 1 {
		t.Errorf("config.TodoStartID is %d, want 1", config.TodoStartID)
	}
	if config.TodoIDType != "int" {
		t.Errorf("config.TodoIDType is %s, want int", config.TodoIDType)
	}
}

// tests the GetConfig invalid inputs
func TestGetConfigError(t *testing.T) {
	// Set the env vars
	os.Setenv("CONFIG_FILE_NAME", "")
	os.Setenv("URL", "")
	os.Setenv("MAX_GOROUTINE", "")
	os.Setenv("MAX_TODO", "")
	os.Setenv("TODO_ID_START", "")
	os.Setenv("TODO_ID_TYPE", "")

	config, err := GetConfig()
	if err == nil {
		t.Error("GetConfig did not return an error")
	}

	if config != nil {
		t.Error("config is not nil")
	}
}
