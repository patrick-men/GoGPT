package cmd

import (
	"fmt"

	"github.com/chzyer/readline"
)

func GetUserInput() (answer string) {

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer func() {
		_ = rl.Close()
	}()

	answer, err = rl.Readline()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return answer
}
