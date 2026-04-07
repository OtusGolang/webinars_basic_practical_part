package inner

import "testing"

func TestAnimal(t *testing.T) {
	a := Animal{}
	a.Name = "Leo"
	a.age = 5
	t.Log(a)
}

func TestMyFunc(t *testing.T) {

	MyAnimalFunc()
}
