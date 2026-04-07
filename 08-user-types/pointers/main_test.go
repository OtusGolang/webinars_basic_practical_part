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

func TestQ1(t *testing.T) {
	a := struct{ x, y int }{0, 0}
	b := a
	a.x = 1
	fmt.Println(b.x) // ?
}

func TestQ2(t *testing.T) {
	var a *struct{ x, y int }
	a = &struct{ x, y int }{}
	//a := &struct{ x, y int }{}
	b := a
	a.x = 1
	fmt.Println(b.x) // ?
	fmt.Printf("%p %p", a, b)
}

func TestQ3(t *testing.T) {
	a := struct{ x *int }{new(int)}
	b := a
	*(a.x) = 1
	fmt.Println(*b.x) // ?
}
