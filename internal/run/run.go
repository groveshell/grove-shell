package run

import (
	"os"
	"os/exec"

	"github.com/groveshell/grove-shell/internal/cmd"
	"github.com/groveshell/grove-shell/internal/lex"
)

func RunCommand(handler *cmd.CommandHandler, input string) error {
	tokens := lex.Lex(input)
	commandName := tokens[0]
	args := tokens[1:]

	exists, err := handler.RunCmd(commandName, args)
    if err != nil {
        return err
    }

    if !exists {

		command := exec.Command(commandName, args...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			return err
		}
	}

    return nil
}
