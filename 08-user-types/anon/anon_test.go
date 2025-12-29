package anon

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
