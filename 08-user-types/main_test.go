package main

import (
	"fmt"
	"testing"
	"webinars_basic_practical_part/08-user-types/inner"
)

func TestDemo(t *testing.T) {

	var user1 struct {
		Name     string
		LastName string
		Age      int
	}

	user1 = struct {
		Name     string
		LastName string
		Age      int
	}{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
	}

	fmt.Println(user1)

}

func TestDog(t *testing.T) {
	d1 := inner.Dog{}
	d1.Animal.Name = "Plutto"
	fmt.Println(d1.Name)
	fmt.Println(d1.TailLength)

	var a1 inner.Animal

	a1 = d1.Animal

	fmt.Println(a1)
}
