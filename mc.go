package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	mcCLI := cli.NewCLI("mc", "0.1.0")
	mcCLI.Args = os.Args[1:]
	mcCLI.Commands = map[string]cli.CommandFactory{
		"delete": func() (cli.Command, error) {
			return mcDeleteCommand{ui: ui}, nil
		},
		"flush": func() (cli.Command, error) {
			return mcFlushCommand{ui: ui}, nil
		},
		"get": func() (cli.Command, error) {
			return mcGetCommand{ui: ui}, nil
		},
		"replace": func() (cli.Command, error) {
			return mcReplaceCommand{ui: ui}, nil
		},
		"set": func() (cli.Command, error) {
			return mcSetCommand{ui: ui}, nil
		},
		"touch": func() (cli.Command, error) {
			return mcTouchCommand{ui: ui}, nil
		},
	}

	exitStatus, err := mcCLI.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
