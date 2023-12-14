package config

import (
	"fmt" 
	"os" 
	"strings"
	"strconv"
	"gopkg.in/yaml.v2"  
	"errors"
)
 
type ToDo struct { 
	Url string `yaml:"url"`
	MaxGoroutine int `yaml:"max_goroutine"`
	MaxTodo int `yaml:"max_todo"`
	TodoStartID int `yaml:"todo_id_start"`
	TodoIDType string `yaml:"todo_id_type"`
} 

type Config struct {
	ToDo `yaml:"todo"`
}

func GetConfig() (*Config, error) {
	var config Config
	configFile := os.Getenv("CONFIG_FILE_NAME")
	configFile = strings.TrimSpace(configFile)
	if configFile != "" {
		_, err := os.Stat(configFile) 
		if err != nil {
			return &config, err
		}
		
		config, err = LoadConfigFile(configFile) 
		if err != nil {
			return nil, fmt.Errorf("failed to load config from file: %v", err)
		}
	} else {
		conf, err := LoadConfigEnv() 
		if err != nil {
			return nil, fmt.Errorf("failed to load config from file: %v", err)
		}
		config = conf
	} 
	return &config, nil
}

func LoadConfigFile(file string) (Config, error) {
	var config Config
 
	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	} 

	return config, nil
}

func LoadConfigEnv() (Config, error) {
	var config Config
	 
	config.Url = os.Getenv("URL")
	if config.Url == "" {
		return config, errors.New("URL is not set")
	}

	maxGoroutine := os.Getenv("MAX_GOROUTINE")
	if maxGoroutine == "" {
		return config, errors.New("MAX_GOROUTINE is not set")
	}
	maxGR, _ := strconv.Atoi(maxGoroutine)
	config.MaxGoroutine = maxGR

	maxTodostr := os.Getenv("MAX_TODO") 
	if maxTodostr == "" {
		return config, errors.New("MAX_TODO is not set")
	}
	maxTodo, _ := strconv.Atoi(maxTodostr)
	config.MaxTodo = maxTodo

	todoStartIDStr := os.Getenv("TODO_ID_START")
	if todoStartIDStr == "" {
		return config, errors.New("TODO_ID_START is not set")
	}
	todoStartID, _ := strconv.Atoi(todoStartIDStr)
	config.TodoStartID = todoStartID

	todoIDType := os.Getenv("TODO_ID_TYPE")
	if todoIDType == "" {
		return config, errors.New("TODO_ID_TYPE is not set")
	}
	config.TodoIDType = todoIDType

	return config, nil
}