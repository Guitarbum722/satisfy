package commands

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var regex = regexp.MustCompile(`^type(\s)+(?P<iface>[_a-zA-Z]+)(\s)+interface`)
var ifaces = []IFace{}

type IFacer interface {
	Do()
	While()
}
type ISearch struct {
	ifaces []IFace
}

func NewISearch() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &ISearch{}, nil
	}
}

func (c *ISearch) Run(args []string) int {

	if len(args) < 1 {
		allInterfaces()
		return 0
	}
	switch args[0] {
	case "filter":
		if len(args) == 2 {
			// add a function to filter
			allInterfaces()
			return 0
		}
		allInterfaces()
		return 0
	default:
		fmt.Println("INVALID SUBCOMMAND for isearch")
		return 1
	}

	return 0
}

// Help returns a description of the command and the options
func (c *ISearch) Help() string {
	return `Usage: satisfy iface <option> <arguments>
  Search for all interfaces in current and sub-directories

Options:
  filter	[prefix]      Filter results by provided prefix
`
}

// Synopsis returns a short description of the isearch command
func (c *ISearch) Synopsis() string {
	return "Search interfaces in package"
}

func allInterfaces() {
	ifaces, err := searchInterfaces("")
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range ifaces {
		fmt.Printf("Interface Name: %s - %s\n", v.name, v.containingFile)
	}
}

func searchInterfaces(prefix string) ([]IFace, error) {
	files := []string{}
	var ifaces []IFace
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			matches := regex.FindStringSubmatch(line)

			if matches != nil {
				iface := IFace{}
				iface.containingFile = file

				for i, n := range regex.SubexpNames() {
					if i == 0 || n == "" {
						continue
					}

					switch n {
					case "iface":
						iface.name = matches[i]
					}
				}
				ifaces = append(ifaces, iface)
			}
		}
		f.Close()
	}
	// if prefix == "" {

	// 	return nil, nil
	// }
	// fmt.Println(is.ifaces)
	return ifaces, nil
}
