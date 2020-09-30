package vflag

import (
	"os"
	"strconv"
	"strings"

	"github.com/nektro/go-util/arrays/stringsu"
	"github.com/spf13/pflag"
)

var (
	flagsS = map[*string][2]string{}
	flagsI = map[*int][2]string{}
	flagsB = map[*bool][2]string{}

	flagaS = []*[]string{}
)

// Parse parses all flags and also checks for environment variable values.
// So --foo-bar could also be read in as FOO_BAR
func Parse() {
	pflag.Parse()

	for k, v := range flagsS {
		nk := strings.ToUpper(strings.ReplaceAll(v[0], "-", "_"))
		ev := os.Getenv(nk)
		if len(ev) > 0 {
			if *k != ev {
				*k = ev
			}
		}
	}
	for k, v := range flagsI {
		nk := strings.ToUpper(strings.ReplaceAll(v[0], "-", "_"))
		ev := os.Getenv(nk)
		if len(ev) > 0 {
			i, err := strconv.ParseInt(ev, 10, 32)
			j := int(i)
			if *k != j && err == nil {
				*k = j
			}
		}
	}
	for k, v := range flagsB {
		nk := strings.ToUpper(strings.ReplaceAll(v[0], "-", "_"))
		ev := os.Getenv(nk)
		if len(ev) > 0 {
			b, err := strconv.ParseBool(ev)
			if *k != b && err == nil {
				*k = b
			}
		}
	}
	for _, item := range flagaS {
		*item = stringsu.Depupe(*item)
	}
}
