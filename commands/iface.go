package commands

// IFace represents a Go Interface which has a containing file
type IFace struct {
	name           string   // name of the interface
	containingFile string   // file that contains the interface
	methods        []string // the methods of the interface
}

// Contains Checks to see if the []IFace contains the iface
func Contains(ifaces []IFace, name string) ([]string, bool) {
	for _, v := range ifaces {
		if v.name == name {
			return v.methods, true
		}
	}
	return nil, false
}
