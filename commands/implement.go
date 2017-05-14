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
func (%s %s%s) %s {

}

`
	// output the signature for each method, while making the receiver type a pointer if applicable
	isPointer := false

	for _, v := range args[1:] {
		switch v {
		case "-p":
			isPointer = !isPointer
			continue
		case "-v":
			isPointer = !isPointer
			continue
		}

		if !isPointer {
			for i := range methods {
				fmt.Printf(funcTemplate, strings.ToLower(string(v[0])), "", v, methods[i])
			}
			continue
		}
		for i := range methods {
			fmt.Printf(funcTemplate, strings.ToLower(string(v[0])), "*", v, methods[i])
		}
	}

	return 0
}

// Help returns a description of the command and the options
func (i *Implement) Help() string {
	return `Usage: satisfy implement <interface-name>  [<option>] <type>, [<option> <type>...]
  Implement the methods of the types for the interface provided

Options:
  -p           Implement method signatures with pointer receivers
                 This flag will apply for all items in the argument list until
                 the "-v" flag is used in the argument list
  -v           Implement method signatures with value receivers
                 This flag will apply for all items in the argument list until
                 the "-p" flag is used in the argument list
Example:
  $ satisfy implement CoolInterface -p CoolStruct CoolerStruct -v AwesomeStruct
`
}

// Synopsis returns a short description of the these command
func (i *Implement) Synopsis() string {
	return "Implement interfaces in a package"
}
