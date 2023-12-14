package cmd

import ( 	
	"fmt"
	"os" 
	"errors"
	"strings"
	"github.com/in-rajendra-pawar/demysttools/handler"
)

type Cmd struct {
	Name string
	DescriptionShort string
	DescriptionLong string
}

func (c *Cmd) Execute() error {
	err := handler.ToDoHandler()
	return err
}

type CliCommand interface {
	Execute() error
}

func ExecuteCommand(command CliCommand) {
	err := command.Execute() 
	if err != nil { 
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetCommand(cmdName string, cmdOption string) (*Cmd, error){
	switch {
		case strings.ToLower(cmdName) == "net" && strings.ToLower(cmdOption) == "todo":
			cmd := Cmd{
					"ToDo",
					"call rest api and fetch todo",
					`use
						net todo 
						
						net call api and print output on terminal
						todo fetch todo
		
						fetches the first 20 even numbered TODO's 
						in most performant way and output the title 
						and whether it is completed or not`,
			}
			return &cmd, nil
		default:
			return nil, errors.New("command not found")
	}
} 