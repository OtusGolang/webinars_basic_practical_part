package parallel

import (
	"fmt"
	"testing"
	"time"
)

// demonstrates simple channel usage to pass data between goroutines
func TestChannelsTransfer(t *testing.T) {
	testToA := make(chan string)
	aToB := make(chan string)
	bToTest := make(chan string)

	go func() {
		input := <-testToA
		aToB <- input + ", processed by A"
	}()

	go func() {
		input := <-aToB
		bToTest <- input + ", processed by B"
	}()

	testToA <- "process this"

	precessed := <-bToTest

	fmt.Println(precessed)

	// testToA <- "process also that"

	// why `testToA <- "process also that"` doesn't work?
}

// shows how to wait for goroutines to finish
func TestWaitForEnd(t *testing.T) {
	results := make(chan int)

	go func() {
		cur := 1
		prev := 0
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			prev, cur = cur, cur+prev
		}

		results <- cur
	}()

	fmt.Println(<-results)
}
