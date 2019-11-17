package ansi

import (
	"github.com/nektro/go-util/ascii"
)

// See: https://en.wikipedia.org/wiki/ANSI_escape_code#Escape_sequences

const (
	SS2 = ascii.ESC + "N"
	SS3 = ascii.ESC + "O"
	DCS = ascii.ESC + "P"
	CSI = ascii.ESC + "["
	ST  = ascii.ESC + "\\"
	OSC = ascii.ESC + "]"
	SOS = ascii.ESC + "X"
	PM  = ascii.ESC + "^"
	APC = ascii.ESC + "_"
	RIS = ascii.ESC + "c"
)
