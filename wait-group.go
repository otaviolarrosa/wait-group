package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//Creating a waitGroup
	var waitGroup sync.WaitGroup

	for i := 0; i <= 5; i++ {
		//Adding new goroutine to a waitGroup "bag"
		waitGroup.Add(1)
		//Calling Goroutine
		go func(i int) {
			//Setting Goroutine as Done, after the routine worker ended
			defer waitGroup.Done()
			//Do the work
			worker(i)
		}(i)
	}

	//Telling to waitGroup, to wait all goroutines to end
	waitGroup.Wait()
}
