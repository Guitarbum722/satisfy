package main

import (
	"fmt"
	"github.com/Guitarbum722/satisfy/commands"
	"github.com/mitchellh/cli"
	"os"
)

const (
	satisfyVersion = "0.0.1"
	satisfyName    = "satisfy"
)

func main() {
	if retval, err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(retval)
	}
}

func run() (int, error) {

	c := &cli.CLI{
		Name:     satisfyName,
		Version:  satisfyVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc(satisfyName),
		Commands: map[string]cli.CommandFactory{
			"isearch": commands.NewISearch(),
		},
	}

	retval, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing program: %s\n", err.Error())
		return 1, err
	}

	return retval, nil
}
