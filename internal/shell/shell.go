package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/groveshell/grove-shell/internal/cmd"
	"github.com/groveshell/grove-shell/internal/run"
)

func StartShell() {
    reader := bufio.NewReader(os.Stdin)

    handler := cmd.NewCommandHandler()
    handler.RegisterNew(cmd.CdCommand{})
    handler.RegisterNew(cmd.EchoCommand{})

    for {
        fmt.Print("# ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal("Failed to read input.")
        }

        hasQuit, err := run.RunCommand(handler, strings.TrimSpace(input))
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        if hasQuit {
            break
        }
    }
}
