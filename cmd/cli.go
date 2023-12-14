package cmd

import(
	"os" 
	"errors"
	"fmt"
)

func Run() error {
	args := os.Args

	if len(args) > 1 {
		cmdName := args[1]
		cmdOption := args[2]

		command, err := GetCommand(cmdName, cmdOption)
		if err != nil {
			return err
		}

		ExecuteCommand(command)

	} else {
		fmt.Println(`uses
		demysttools net todo`)
		return errors.New("not enough arguments supplied") 
	}
	return nil
}