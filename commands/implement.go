package commands

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"strings"
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

	ifaces, err := searchInterfaces(false)
	if err != nil {
		log.Fatalln(err)
	}
	for i := range ifaces {
		ifaces[i].methods = append(ifaces[i].methods, tempMethods[ifaces[i].name]...)
	}
	methods, found := Contains(ifaces, args[0])

	if !found {
		fmt.Printf("This package does not contain the specified interface: %s\n", args[0])
		return 0
	}

	funcTemplate := `
func (%s %s) %s {

}

`
	// TODO(guitarbum722) provide a flag so that the receiver can be a pointer to the type 2017-05-10T18:30 4
	for _, arg := range args[1:] {
		for i := range methods {
			fmt.Printf(funcTemplate, strings.ToLower(string(arg[0])), arg, methods[i])
		}
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
