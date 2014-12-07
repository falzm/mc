package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mitchellh/cli"
)

type mcDeleteCommand struct {
	ui cli.Ui
}

func (cmd mcDeleteCommand) Help() string {
	helpText := `
Usage: mc delete [options] <key>

  Delete <key> on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
`
	return strings.TrimSpace(helpText)
}

func (cmd mcDeleteCommand) Run(args []string) int {
	var (
		serverHost string
		serverPort int
	)

	cmdFlags := flag.NewFlagSet("delete", flag.ExitOnError)
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

	if err := mc.Delete(key); err != nil {
		if err == memcache.ErrCacheMiss {
			return 0
		}

		cmd.ui.Error(fmt.Sprintf("Error: unable to delete key: %s", err))
		return 1
	}

	cmd.ui.Output("OK")

	return 0
}

func (cmd mcDeleteCommand) Synopsis() string {
	return "delete a key"
}
