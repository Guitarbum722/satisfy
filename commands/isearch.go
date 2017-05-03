package commands

import (
	"fmt"
	"github.com/mitchellh/cli"
	_ "os"
)

type ISearch struct{}

func NewISearch() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &ISearch{}, nil
	}
}

func (c *ISearch) Run(_ []string) int {
	fmt.Println("TODO ***")
	return 0
}

func (c *ISearch) Help() string {
	return `Usage: satisfy iface <option> <arguments>
  Search for all interfaces in current and sub-directories

Options:
  filter	[{first letter}]      Filter results by provided letter
`
}

func (c *ISearch) Synopsis() string {
	return "Search interfaces in package"
}
