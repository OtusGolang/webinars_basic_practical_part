package pointers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	a := 4

	b := &a

	c := *b

	fmt.Println(a, b, c)
}
