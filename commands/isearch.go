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

var regex = regexp.MustCompile(`^type(\s)+(?P<iface>[a-zA-Z1-9_]+)(\s)+interface`)

var tempMethods = make(map[string][]string)

// ISearch provides commands to output a list of interfaces in the current directory tree
type ISearch struct{}

// NewISearch returns a pointer to an ISearch
func NewISearch() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &ISearch{}, nil
	}
}

// Run performs the isearch command with options
func (c *ISearch) Run(args []string) int {
	if len(args) < 1 {
		allInterfaces(true)
		return 0
	}
	switch args[0] {
	case "filter":
		if len(args) == 2 {
			switch args[1] {
			case "-e":
				exportedInterfaces(false)
				return 0
			case "-v":
				allInterfaces(true)
				return 0
			}
		}
		allInterfaces(false)
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

func allInterfaces(verbose bool) {
	ifaces, err := searchInterfaces(false)

	if err != nil {
		log.Fatalln(err)
	}

	for i := range ifaces {
		ifaces[i].methods = append(ifaces[i].methods, tempMethods[ifaces[i].name]...)
	}

	display(ifaces, verbose)
}

func exportedInterfaces(verbose bool) {
	ifaces, err := searchInterfaces(true)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range ifaces {
		ifaces[i].methods = append(ifaces[i].methods, tempMethods[ifaces[i].name]...)
	}

	display(ifaces, verbose)
}

// TODO(guitarbum722) pretty the output so that it is at least aligned nicely 2017-05-19T18:00 2
func display(ifaces []IFace, verbose bool) {
	for _, v := range ifaces {
		fmt.Printf("Interface Name: %s - %s\n", v.name, v.containingFile)
		if verbose {
			for _, m := range v.methods {
				fmt.Printf("\t%s\n", m)
			}
		}
	}
}

func searchInterfaces(exported bool) ([]IFace, error) {
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

			// account for an interface with NO methods which might have the opening and closing brace on same line
			if ifaceFound && (line == "" || line[0] == '}') {
				ifaceFound = false
				tempName = ""
				continue
			}
			if ifaceFound {
				line = strings.TrimLeft(line, "\t")
				tempMethods[tempName] = append(tempMethods[tempName], line)
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

	return ifaces, nil
}
