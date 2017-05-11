package commands

import (
	_ "bufio"
	"fmt"
	"github.com/mitchellh/cli"
	_ "log"
	_ "os"
	_ "path/filepath"
	_ "regexp"
	_ "strings"
)

// Implement provides commands to provide the skeletal implementation of the specified interface
type Implement struct{}

// NewImplement returns a pointer to a Implement
func NewImplement() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Implement{}, nil
	}
}

// Run performs the IFace command with options
func (i *Implement) Run(args []string) int {
	if len(args) < 2 {
		fmt.Printf("NOT ENOUGH ARGUMENTS TO COMMAND implement\n\n")
		fmt.Println(i.Help())
		return 0
	}
	return 0
}

// Help returns a description of the command and the options
func (i *Implement) Help() string {
	return `Usage: satisfy implement <interface-name> <type>, [<type>...]
  Implement the methods of the types for the interface provided

Options:
`
}

// Synopsis returns a short description of the these command
func (i *Implement) Synopsis() string {
	return "Implement interfaces in a package"
}
