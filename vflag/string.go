package vflag

import (
	"os"
	"strings"

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

// StringArrayVar registers a string array flag. for the env var, the first char
// is the splitter and the rest is the items. ie `:hello:world` vs `,hello,world`
func StringArrayVar(p *[]string, name string, value []string, usage string) {
	pflag.StringArrayVar(p, name, value, usage)
	s := os.Getenv(strings.ReplaceAll(strings.ToUpper(name), "-", "_"))
	if len(s) <= 1 {
		return
	}
	for _, item := range strings.Split(s[1:], string(s[0])) {
		*p = append(*p, item)
	}
}
