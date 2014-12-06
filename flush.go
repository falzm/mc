package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mitchellh/cli"
)

type mcFlushCommand struct {
	ui cli.Ui
}

func (cmd mcFlushCommand) Help() string {
	helpText := `
Usage: mc flush [options]

  Flush all keys stored on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
`
	return strings.TrimSpace(helpText)
}

func (cmd mcFlushCommand) Run(args []string) int {
	var (
		serverHost string
		serverPort int
	)

	cmdFlags := flag.NewFlagSet("flush", flag.ExitOnError)
	cmdFlags.Usage = func() { cmd.ui.Output(cmd.Help()) }
	cmdFlags.StringVar(&serverHost, "serverHost", "127.0.0.1", "memcached server host")
	cmdFlags.IntVar(&serverPort, "serverPort", 11211, "memcached server port")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	mc := memcache.New(fmt.Sprintf("%s:%d", serverHost, serverPort))

	if err := mc.FlushAll(); err != nil {
		cmd.ui.Error(fmt.Sprintf("Error: unable to flush key: %s", err))
		return 1
	}

	cmd.ui.Output("OK")

	return 0
}

func (cmd mcFlushCommand) Synopsis() string {
	return "flush all keys"
}
