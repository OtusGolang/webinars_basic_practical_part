package parallel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {

	done := doSmth()

	timeout := time.After(2 * time.Second)
	fmt.Println("main: started")

	select {
	case <-done:
		fmt.Println("main: done")
	case <-timeout:
		fmt.Println("main: timeout")
	}

	// Possible temporary mem leak.
	// https://arangodb.com/2020/09/a-story-of-a-memory-leak-in-go-how-to-properly-use-time-after/

	fmt.Println("main: stopped")

}

// doSmth simulates some work 0 to 3 seconds
func doSmth() chan struct{} {
	done := make(chan struct{})
	go func() {
		ms := rand.Intn(3000)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		fmt.Println("doSmth: done after", ms)
		close(done)
	}()
	return done
}

func TestFeeder(t *testing.T) {
	consumer := func(input chan int) {
		for range input {
			ms := rand.Intn(1000)
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}
	}

	ch := make(chan int)
	go consumer(ch)

	for i := 0; i < 100; i++ {
		select {
		case ch <- i:
			fmt.Println("main: sent")
		default:
			fmt.Println("main: overflow")
			// TODO: secondary consumer?
		}
		ms := rand.Intn(1000)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

}
