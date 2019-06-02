package alias

import (
	"errors"
	"fmt"
	"time"
)

// F is an alias for fmt.Sprintf
func F(base string, args ...interface{}) string {
	return fmt.Sprintf(base, args...)
}

// E is an alias for errors.New
func E(message string) error {
	return errors.New(message)
}

// T is a shortcut to obtain the current UTC date-time in a 'yyyy-mm-dd hh:mm:ss' format
func T() string {
	return time.Now().UTC().String()[0:19]
}

// D is a shortcut to obtain the current UTC date in a 'yyyy-mm-dd' format
func D() string {
	return T()[0:10]
}
