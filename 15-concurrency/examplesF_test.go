package parallel

import (
	"fmt"
	"testing"
)

// shows advantages of using directional channels
func TestDirected(t *testing.T) {

	c := make(chan int, 1)
	c <- 1
	<-c

	// cWrite := (chan<- int)(c)
	// cWrite <- 9

	// cRead := (<-chan int)(c)
	// fmt.Println(<-cRead)

	sender(c)
	reader(c)
}

func reader(ch <-chan int) {
	fmt.Println(<-ch)
	// ch <- 1 // compiler error
}
func sender(ch chan<- int) {
	// fmt.Println(<-ch) // compiler error
	ch <- 1
}

// note: mergeChannels can benefit from this!
