package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mitchellh/cli"
)

type mcTouchCommand struct {
	ui cli.Ui
}

func (cmd mcTouchCommand) Help() string {
	helpText := `
Usage: mc touch [options] <key> <TTL>

  Update the expiry of <key> to <TTL> seconds on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
`
	return strings.TrimSpace(helpText)
}

func (cmd mcTouchCommand) Run(args []string) int {
	var (
		serverHost string
		serverPort int
	)

	cmdFlags := flag.NewFlagSet("touch", flag.ExitOnError)
	cmdFlags.Usage = func() { cmd.ui.Output(cmd.Help()) }
	cmdFlags.StringVar(&serverHost, "serverHost", "127.0.0.1", "memcached server host")
	cmdFlags.IntVar(&serverPort, "serverPort", 11211, "memcached server port")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if cmdFlags.NArg() < 2 {
		cmd.ui.Error("Error: missing arguments\n")
		cmd.ui.Error(cmd.Help())
		return 1
	}

	key := cmdFlags.Arg(0)
	ttl := cmdFlags.Arg(1)

	ttlSeconds, err := strconv.ParseInt(ttl, 10, 32)
	if err != nil {
		cmd.ui.Error("Error: invalid TTL value")
		return 1
	}

	mc := memcache.New(fmt.Sprintf("%s:%d", serverHost, serverPort))

	if err := mc.Touch(key, int32(ttlSeconds)); err != nil {
		cmd.ui.Error(fmt.Sprintf("Error: unable to update key %q expiry: %s", key, err))
		return 1
	}

	cmd.ui.Output("OK")

	return 0
}

func (cmd mcTouchCommand) Synopsis() string {
	return "update the expiry of a key"
}
