package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mitchellh/cli"
)

type mcReplaceCommand struct {
	ui cli.Ui
}

func (cmd mcReplaceCommand) Help() string {
	helpText := `
Usage: mc replace [options] <key> <value>

  Replace the current value of <key> with <value> on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
  -ttl=N            key expiration time in seconds (default: 0, i.e. no expiration)
`
	return strings.TrimSpace(helpText)
}

func (cmd mcReplaceCommand) Run(args []string) int {
	var (
		serverHost string
		serverPort int
		ttl        int
	)

	cmdFlags := flag.NewFlagSet("replace", flag.ExitOnError)
	cmdFlags.Usage = func() { cmd.ui.Output(cmd.Help()) }
	cmdFlags.StringVar(&serverHost, "serverHost", "127.0.0.1", "memcached server host")
	cmdFlags.IntVar(&serverPort, "serverPort", 11211, "memcached server port")
	cmdFlags.IntVar(&ttl, "ttl", 0, "key expiration time in seconds")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if cmdFlags.NArg() < 2 {
		cmd.ui.Error("Error: missing arguments\n")
		cmd.ui.Error(cmd.Help())
		return 1
	}

	key := cmdFlags.Arg(0)
	val := cmdFlags.Arg(1)

	mc := memcache.New(fmt.Sprintf("%s:%d", serverHost, serverPort))

	if err := mc.Replace(&memcache.Item{
		Key:        key,
		Value:      []byte(val),
		Expiration: int32(ttl),
	}); err != nil {
		cmd.ui.Error(fmt.Sprintf("Error: unable to replace key %q value: %s", key, err))
		return 1
	}

	cmd.ui.Output("OK")

	return 0
}

func (cmd mcReplaceCommand) Synopsis() string {
	return "replace the value to a key if it exists"
}
