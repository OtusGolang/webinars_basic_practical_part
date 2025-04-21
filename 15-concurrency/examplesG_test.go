package parallel

import (
	"fmt"
	"testing"
	"time"
)

// how channels are compared?
func TestChanRef(t *testing.T) {
	c1 := make(chan int)
	fmt.Println("c1 == c1:", c1 == c1) // nolint: staticcheck // for demonstration

	c2 := make(chan int)
	fmt.Println("c1 == c2:", c1 == c2)

	c1a := c1
	fmt.Println("c1 == c1a:", c1 == c1a)
	fmt.Printf("c1: %p, c1a: %p\n", c1, c1a)
}

// shows that channels are references
func TestRW(t *testing.T) {
	cWrite := make(chan int, 1)
	cRead := cWrite

	cWrite <- 1
	fmt.Println(<-cRead) // got "1" from 'cRead' but read 'cWrite'
}

// how nil channels work?
func TestNil(t *testing.T) {
	var c chan int

	fmt.Println("PHASE 1")
	go func() {
		_, ok := <-c // blocks
		fmt.Println("read from a nil channel:", ok)
	}()
	fmt.Println("PHASE 2")

	go func() {
		c <- 1 // blocks
		fmt.Println("sent to a nil channel")
	}()

	time.Sleep(1 * time.Second)

	close(c) // panic
}

// how closed channels work?
func TestClosed(t *testing.T) {
	c := make(chan int)
	close(c)

	_, ok := <-c // no wait
	fmt.Println("read from a closed channel:", ok)

	c <- 1 // panic
	fmt.Println("sent to a closed channel")

}
