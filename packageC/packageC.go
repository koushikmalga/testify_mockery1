package packageC

import (
	"fmt"
)

type CIfc interface {
	Sub(int, int) int
}
type C struct{}

func (c *C) Sub(ele1, ele2 int) int {
	fmt.Println("hoal")
	return ele1 - ele2
}
