package ansi

import (
	"strconv"
	"strings"
)

func mapItoS(n []int) []string {
	res := make([]string, len(n))
	for i := 0; i < len(n); i++ {
		res[i] = strconv.Itoa(n[i])
	}
	return res
}

func MakeCSISeq(c string, x ...int) string {
	return CSI + strings.Join(mapItoS(x), ";") + c
}
