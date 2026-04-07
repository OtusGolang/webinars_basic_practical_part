package main

// Псевдонимы типов и пользовательские типы в Go

// Они могут быть такими:

type MyType int

type MyInterface interface {
	MyMethod()
}

type MyFunc func(count int) int

type MySlice []string
type MyMap map[string]int
type MyChan chan int

type MyStruct struct {
	Field1 string
	Field2 int
}
