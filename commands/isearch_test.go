package commands

import (
	"testing"
)

// IFacer2 is a sample
type IFacer2 interface {
	Do2(n int, s []string) error
	While2(s string) (string, error)
}

type sampler2 interface {
	punch2() error
	kick2() string
}

type Cooler interface{}

type Heater interface{}

func TestSearchExportedInterfaces(t *testing.T) {
	got, _ := searchInterfaces(true)

	for _, v := range got {
		if v.name[0] >= 'a' && v.name[0] <= 'z' {
			t.Fatalf("searchInterfaces( exportedOnly ) = %q, should not return unexported", v.name)

		}
	}
}
