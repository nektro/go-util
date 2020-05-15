package vflag

import (
	"github.com/spf13/pflag"
)

// String registers a string flag
func String(name string, value string, usage string) *string {
	var p string
	StringVar(&p, name, value, usage)
	return &p
}

// StringVar registers a string flag
func StringVar(p *string, name string, value string, usage string) {
	pflag.StringVar(p, name, value, usage)
	flagsS[p] = [...]string{name, value}
}

// StringArrayVar is filler and simply calls pflag
func StringArrayVar(p *[]string, name string, value []string, usage string) {
	pflag.StringArrayVar(p, name, value, usage)
}
