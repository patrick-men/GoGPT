package main

import (
	API "GoGPT/Prompt"
	"GoGPT/cmd"
	"fmt"
)

func main() {

	prompt := cmd.GetUserInput()
	answer := API.APIrequest(prompt)

	fmt.Println(answer)
}
