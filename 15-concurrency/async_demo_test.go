package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestRun2TimesSync(t *testing.T) {
	// Lecturer, fold this function during the demonstration
	JustRegularFunction()
	JustRegularFunction()
}
func TestRun2TimesAsync(t *testing.T) {
	ready := make(chan struct{})
	// Lecturer, fold this function during the demonstration
	waiter := make(chan struct{}, 2)

	go func() {
		<-ready
		JustRegularFunction()
		waiter <- struct{}{}
	}()
	go func() {
		<-ready
		JustRegularFunction()
		waiter <- struct{}{}
	}()

	time.Sleep(1 * time.Millisecond) // allow goroutines to line up
	close(ready)                     // start the goroutines

	<-waiter // wait for the first goroutine to finish
	<-waiter // wait for the second goroutine to finish
}

func JustRegularFunction() {
	// This function is here to only appear in the stack during the debugging
	AnotherFunctionThatDoSomething()
}

var a int

func AnotherFunctionThatDoSomething() {

	fmt.Println("Doing something...") // breakpoint here

	a = 0
	b := 0
	for i := 0; i < 1000000; i++ {
		a++
		b++
	}

	fmt.Printf("And the result is: a=%d, b=%d", a, b)
	fmt.Println()

}
