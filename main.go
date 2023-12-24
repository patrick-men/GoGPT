package main

import (
	API "GoGPT/Prompt"
	"GoGPT/cmd"
	"fmt"
)

func main() {

	var prompt string

	if cmd.CheckFlag() {
		prompt = cmd.ReadFlag()
	} else {
		prompt = cmd.GetUserInput()
	}

	answer := API.APIrequest(prompt)

	fmt.Println(answer)
}
