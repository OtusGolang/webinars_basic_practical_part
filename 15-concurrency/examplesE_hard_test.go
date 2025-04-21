package parallel

import (
	"fmt"
	"testing"
	"time"
)

func MakeGenerator(id string, delay time.Duration) chan struct{} {
	c := make(chan struct{})
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(delay)
			c <- struct{}{}
			fmt.Printf("generator %s: sent #%d\n", id, i)
		}
		close(c)
	}()
	return c
}

// demonstrates waiting data from multiple goroutines
func TestChanJoin(t *testing.T) {

	generatorOutputA := MakeGenerator("A", 149*time.Millisecond)
	generatorOutputB := MakeGenerator("B", 313*time.Millisecond)
	generatorOutputC := MakeGenerator("C", 541*time.Millisecond)

	//time.Sleep(1 * time.Second)

	for {
		select {
		case <-generatorOutputA:
			fmt.Println("main: got A")
		case <-generatorOutputB:
			fmt.Println("main: got B")
		case <-generatorOutputC:
			fmt.Println("main: got C")
		}
	}

	// how to stop waiting? ...close!
}

// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
func TestChanJoinWithCloseTry(t *testing.T) {
	// highlights problems with select and closed channels

	generatorOutputA := MakeGenerator("A", 149*time.Millisecond)
	generatorOutputB := MakeGenerator("B", 313*time.Millisecond)
	generatorOutputC := MakeGenerator("C", 541*time.Millisecond)

wait:
	for {
		select {
		case _, ok := <-generatorOutputA:
			if !ok {
				break wait
			}
			fmt.Println("main: got A")
		case _, ok := <-generatorOutputB:
			if !ok {
				break wait
			}
			fmt.Println("main: got B")
		case _, ok := <-generatorOutputC:
			if !ok {
				break wait
			}
			fmt.Println("main: got C")
		}
	}

	fmt.Println("main: all generators closed")

}

// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
func TestMergeChannels(t *testing.T) {
	// https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de

	mergeChannels := func(channels ...chan struct{}) chan struct{} {
		out := make(chan struct{})
		closedChannels := make(chan struct{}, len(channels))
		// Start a goroutine for each channel
		for _, ch := range channels {
			go func(ch chan struct{}) {
				for val := range ch {
					out <- val // Send values to the output channel
				}
				closedChannels <- struct{}{}
			}(ch)
		}

		go func() {
			for i := 0; i < len(channels); i++ {
				<-closedChannels
			}
			close(out)
		}()

		return out
	}

	generatorOutputA := MakeGenerator("A", 149*time.Millisecond)
	generatorOutputB := MakeGenerator("B", 313*time.Millisecond)
	generatorOutputC := MakeGenerator("C", 541*time.Millisecond)

	mergedChan := mergeChannels(generatorOutputA, generatorOutputB, generatorOutputC)
	cnt := 0
	for range mergedChan {
		cnt++
		fmt.Printf("got a message #%d\n", cnt)
	}

}

// other option: https://gist.github.com/montanaflynn/d9b358249939c5c541ec

// shows how to use select to send to channels
// ... also how 'default' works
func TestSelectToSend(t *testing.T) {

	consumer := func(c <-chan int) {
		for v := range c {
			fmt.Println("consumer: got", v)
		}
		fmt.Println("consumer: closed")
	}

	c1 := make(chan int)
	c2 := make(chan int)

	go consumer(c1)
	go consumer(c2)

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case c1 <- i:
		case c2 <- i:
			// default:
			// 	fmt.Println("main: no consumer is ready")
			// 	i--
			// 	time.Sleep(100 * time.Millisecond)
		}
	}
}
