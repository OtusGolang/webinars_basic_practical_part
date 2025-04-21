package parallel

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	// Test caching can ruin the demonstration. Set use "-count=1" argument to avoid caching.

	log.SetFlags(log.Lmicroseconds)

	done := doSmth()

	timeout := time.After(2 * time.Second)
	log.Println("main: started")

	select {
	case <-done:
		log.Print("main: done")
	case <-timeout:
		log.Print("main: timeout")
	}

	// Possible temporary mem leak.
	// https://arangodb.com/2020/09/a-story-of-a-memory-leak-in-go-how-to-properly-use-time-after/

	log.Print("main: stopped")

}

// doSmth simulates some work 0 to 3 seconds
func doSmth() chan struct{} {
	done := make(chan struct{})
	go func() {
		ms := rand.Intn(3000)
		log.Printf("doSmth: it will take %d ms", ms)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		log.Print("doSmth: done")
		close(done)
	}()
	return done
}

// TestFeeder demonstrates example of rejecting data if overflow
func TestFeeder(t *testing.T) {
	consumer1PerSec := func(input chan int) {
		for range input {
			ms := rand.Intn(1000)
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}
	}

	ch := make(chan int)
	go consumer1PerSec(ch)

	for i := 0; i < 100; i++ {
		select {
		case ch <- i:
			log.Println("main: sent")
		default:
			log.Println("main: overflow")
			// TODO: secondary consumer?
		}
		ms := rand.Intn(1000)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

}
