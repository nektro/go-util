package main

import (
	"log"

	"github.com/nektro/go-util/vflag"
)

func main() {
	arr := []string{}
	vflag.StringArrayVar(&arr, "test", []string{}, "")
	vflag.Parse()

	log.Println(len(arr), arr)
}
