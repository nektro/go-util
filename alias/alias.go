package alias

import (
	"fmt"
)

func F(base string, args ...interface{}) string {
	return fmt.Sprintf(base, args...)
}
