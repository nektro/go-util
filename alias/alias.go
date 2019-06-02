package alias

import (
	"errors"
	"fmt"
	"time"
)

func F(base string, args ...interface{}) string {
	return fmt.Sprintf(base, args...)
}

func E(message string) error {
	return errors.New(message)
}

func T() string {
	return time.Now().UTC().String()[0:19]
}

func D() string {
	return T()[0:10]
}
