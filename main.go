package testmod

import (
	"fmt"
)

//Hi returns friendly greatings
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
