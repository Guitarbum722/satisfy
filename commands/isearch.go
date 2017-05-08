package commands

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`^type(\s)+(?P<iface>[a-zA-Z_]+)(\s)+interface`)
var ifaces = []IFace{}

// ISearch provides commands to output a list of interfaces in the current directory tree
type ISearch struct {
	ifaces []IFace
}

// NewISearch returns a pointer to an ISearch
func NewISearch() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &ISearch{}, nil
	}
}

// Run performs the isearch command with options
func (c *ISearch) Run(args []string) int {
	if len(args) < 1 {
		allInterfaces()
		return 0
	}
	switch args[0] {
	case "filter":
		if len(args) == 2 {
			switch args[1] {
			case "-e":
				exportedInterfaces()
				return 0
			}
		}
		allInterfaces()
	default:
		fmt.Println("INVALID SUBCOMMAND for isearch")
	}
	return 0
}

// Help returns a description of the command and the options
func (c *ISearch) Help() string {
	return `Usage: satisfy iface <option> <arguments>
  Search for all interfaces in current and sub-directories

Options:
  filter	     [-e exported] [-v verbose]                 Filters the output
    -e only displays exported interfaces
    -v verbose displays the interface's methods along with the name
`
}

// Synopsis returns a short description of the isearch command
func (c *ISearch) Synopsis() string {
	return "Search interfaces in package"
}

func allInterfaces() {
	ifaces, err := searchInterfaces(false)
	if err != nil {
		log.Fatalln(err)
	}
	display(ifaces)
}

func exportedInterfaces() {
	ifaces, err := searchInterfaces(true)
	if err != nil {
		log.Fatalln(err)
	}
	display(ifaces)
}

func display(ifaces []IFace) {
	for _, v := range ifaces {
		fmt.Printf("Interface Name: %s - %s\n", v.name, v.containingFile)
	}
}

func searchInterfaces(exported bool) ([]IFace, error) {
	files := []string{}
	var ifaces []IFace
	var tempMethods = make(map[string][]string)

	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		// only necessary to scan .go files
		if !strings.HasSuffix(file, ".go") {
			continue
		}

		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(f)
		var ifaceFound bool
		var tempName string
		for scanner.Scan() {
			line := scanner.Text()

			if ifaceFound && line[0] == '}' {
				ifaceFound = false
				continue
			}
			if ifaceFound {
				tempMethods[tempName] = append(tempMethods[tempName], line)
				tempName = ""
			}

			matched := regex.FindStringSubmatch(line)

			if matched != nil {
				ifaceFound = true
				iface := IFace{}
				iface.containingFile = file

				for i, n := range regex.SubexpNames() {
					if i == 0 || n == "" {
						continue
					}

					switch n {
					case "iface":
						iface.name = matched[i]
						tempName = iface.name
					}
				}
				switch exported {
				case true:
					if iface.name[0] >= 'A' && iface.name[0] <= 'Z' {
						ifaces = append(ifaces, iface)
					}
				default:
					ifaces = append(ifaces, iface)

				}
			}
		}
		f.Close()
	}
	for k := range tempMethods {
		fmt.Printf("%v :::\n", k)
		for _, v := range tempMethods[k] {
			fmt.Println(v)
		}
	}
	return ifaces, nil
}
