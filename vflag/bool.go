package vflag

import (
	"strconv"

	"github.com/spf13/pflag"
)

// Bool registers a bool flag
func Bool(name string, value bool, usage string) *bool {
	var p bool
	BoolVar(&p, name, value, usage)
	return &p
}

// BoolVar registers a bool flag
func BoolVar(p *bool, name string, value bool, usage string) {
	pflag.BoolVar(p, name, value, usage)
	flagsB[p] = [...]string{name, strconv.FormatBool(value)}
}
