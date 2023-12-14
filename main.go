package main

import (
	"fmt"
	"github.com/in-rajendra-pawar/demysttools/cmd"
)

func main() {

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
