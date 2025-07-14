package main

import "fmt"

type MyI interface {
	MyMethod(a string, b string) (string, error)
	OtherMethod()
}

type MyImpl struct {
	Name string
}

func (m *MyImpl) MyMethod(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
func (m *MyImpl) OtherMethod() {
	fmt.Println("ObjMethod called")
}

type MyImpl2 struct {
}

func (m *MyImpl2) MyMethod(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
func (m *MyImpl2) OtherMethod() {
	fmt.Println("MyImpl2 OtherMethod called")
}

func main() {

	var obj = &MyImpl{Name: "MyImpl Object"}
	obj.OtherMethod()

	var myInterface MyI
	myInterface = obj
	myInterface.OtherMethod()
	//fmt.Printf("i is of type %T and value %v\n", myInterface, myInterface)

	processValue(myInterface)

}

func processValue(i MyI) {
	switch v := i.(type) {
	case *MyImpl:
		fmt.Println("Processing MyImpl with Name:", v.Name)
	case *MyImpl2:
		fmt.Println("Processing MyImpl2", v)
	default:
		fmt.Println("Unknown type", v)
	}
}
