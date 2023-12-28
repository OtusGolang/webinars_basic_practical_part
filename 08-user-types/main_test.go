package main

import (
	"fmt"
	"reflect"
	"testing"
	"webinars_basic_practical_part/08-user-types/inner"

	"github.com/stretchr/testify/assert"
)

func TestAnonTypesEquality(t *testing.T) {

	v1 := struct {
		a int
		b string
	}{1, "two"}
	v2 := struct {
		a int
		b string
	}{3, "four"}

	fmt.Println(v1, v2)
	assert.True(t, reflect.TypeOf(v1).AssignableTo(reflect.TypeOf(v2)))

	v1 = v2 // Works!

	v3 := inner.GetInstance()
	fmt.Println(v3)
	//v1 = v3 // Same type, but not works now!
}

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
