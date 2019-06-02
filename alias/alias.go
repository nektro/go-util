package alias

import (
	"errors"
	"fmt"
)

func F(base string, args ...interface{}) string {
	return fmt.Sprintf(base, args...)
}

func E(message string) error {
	return errors.New(message)
}
