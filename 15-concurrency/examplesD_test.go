package parallel

import (
	"fmt"
	"testing"
	"time"
)

// shows how buffered channel stores data
func TestBufferedSimple(t *testing.T) {
	tasks := make(chan string, 3)

	tasks <- "dog"
	tasks <- "cat"
	tasks <- "bird"
	//tasks <- "fish" // will block here

	fmt.Println(<-tasks)
	fmt.Println(<-tasks)
	fmt.Println(<-tasks)

	// highlights:
	// can send to channel even if no one is waiting for it

}

func TestBuffered(t *testing.T) {
	tasksChan := make(chan string, 6) // !

	fmt.Println("main: sending tasks")
	tasks := []string{"dog", "cat", "bird", "fish", "snake", "turtle"}
	for _, task := range tasks {
		tasksChan <- task
		fmt.Println("main: sent task", task)
	}

	fmt.Println("main: all tasks sent")

	fmt.Println("main: starting workers") // try move this after tasks sending
	go worker(1, tasksChan)
	go worker(2, tasksChan)

	// other highlights:
	// one channel can be used by multiple goroutines
	// useful pattern: one goroutine sends tasks, other goroutines process them
	time.Sleep(10 * time.Second)

}

func worker(workerId int, tasksChan chan string) {
	fmt.Printf("worker %d: started\n", workerId)

	for task := range tasksChan {
		time.Sleep(100 * time.Millisecond) // processing
		fmt.Printf("worker %d: task %s processed\n", workerId, task)
	}

	fmt.Printf("worker %d: finished\n", workerId)
}
