package cmd

import (
	"fmt"
	"github.com/chzyer/readline"
)

func GetUserInput() (prompt string) {

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer func() {
		_ = rl.Close()
	}()

	prompt, err = rl.Readline()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return prompt
}
