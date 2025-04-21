package parallel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// shows difference between parallelism and concurrency
func TestGoMaxProc(t *testing.T) {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU(), runtime.GOMAXPROCS(0))

	for i := 0; i < 10; i++ {
		go func(num int) {
			// never use "i" here
			fmt.Println("start", num)
			// runtime.Gosched()
			for i := 0; i < 1000000; i++ {
			} // why sleep does not help to block here (runtime.Gosched())?
			//time.Sleep(1 * time.Millisecond)
			fmt.Println("stop", num)
		}(i) // pass a parameter
	}

	// try runtime.GOMAXPROCS(2), 3,4...
	time.Sleep(1 * time.Second) // why need this?
}
