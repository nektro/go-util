package vflag

import (
	"strconv"

	"github.com/spf13/pflag"
)

// Int registers an int flag
func Int(name string, value int, usage string) *int {
	var p int
	IntVar(&p, name, value, usage)
	return &p
}

// IntVar registers an int flag
func IntVar(p *int, name string, value int, usage string) {
	pflag.IntVar(p, name, value, usage)
	flagsI[p] = [...]string{name, strconv.Itoa(value)}
}
