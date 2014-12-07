package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mitchellh/cli"
)

type mcGetCommand struct {
	ui cli.Ui
}

func (cmd mcGetCommand) Help() string {
	helpText := `
Usage: mc get [options] <key>

  Get the value associated to <key> on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
`
	return strings.TrimSpace(helpText)
}

func (cmd mcGetCommand) Run(args []string) int {
	var (
		serverHost string
		serverPort int
	)

	cmdFlags := flag.NewFlagSet("get", flag.ExitOnError)
	cmdFlags.Usage = func() { cmd.ui.Output(cmd.Help()) }
	cmdFlags.StringVar(&serverHost, "serverHost", "127.0.0.1", "memcached server host")
	cmdFlags.IntVar(&serverPort, "serverPort", 11211, "memcached server port")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if cmdFlags.NArg() < 1 {
		cmd.ui.Error("Error: missing arguments\n")
		cmd.ui.Error(cmd.Help())
		return 1
	}

	key := cmdFlags.Arg(0)
	mc := memcache.New(fmt.Sprintf("%s:%d", serverHost, serverPort))

	item, err := mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			cmd.ui.Output("NOT FOUND")
			return 0
		}

		cmd.ui.Error(fmt.Sprintf("Error: unable to get value: %s", err))
		return 1
	}

	cmd.ui.Output(string(item.Value))

	return 0
}

func (cmd mcGetCommand) Synopsis() string {
	return "get the value of a key"
}
