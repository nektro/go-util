package util

import (
	"fmt"
	"time"
)

func Log(message ...interface{}) {
	fmt.Print("[" + GetIsoDateTime() + "] ")
	fmt.Println(message...)
}

func GetIsoDateTime() string {
	vil := time.Now().UTC().String()
	return vil[0:19]
}
