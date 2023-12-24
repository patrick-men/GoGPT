package cmd

import (
	"os"
)

func CheckFlag() (exists bool){
	
	exists = false
	
	if len(os.Args) > 1 {
		// A flag is provided
		exists = true
	} 
	return
}

func ReadFlag() (prompt string) {
	prompt = os.Args[1]
	return
}