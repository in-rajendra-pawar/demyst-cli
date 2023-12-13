package cmd

import ( 	
	"fmt"
	"os" 
)

type Cmd struct {
	Name string
	DescriptionShort string
	DescriptionLong string
}

type CliCommand interface {
	Execute() error
}

func Run() {
	fmt.Println("cmd.Run()")
	args := os.Args
	 
	if len(args) > 1 {
		cmdName := args[1]
		cmdOption := args[2]

		command, err := GetCommand(cmdName, cmdOption)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_ = command.Execute()
	} else {
		fmt.Println("not enough arguments supplied")
	}
}

func ExecuteCommand(command CliCommand) {
	err := command.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}