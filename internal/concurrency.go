package internal

import (
	"fmt"
	"math/rand"
	"sync"
)

// SquaringNumbers Squaring numbers
func SquaringNumbers() {
	in := make([]int8, 10)

	var i int8
	for i = 0; i < 10; i++ {
		in[i] = i
	}

	out := make(chan int8, 10)

	go func() {
		for value := range in {
			out <- int8(value * 2)
		}
	}()

	go func() {
		for value := range out {
			fmt.Println(value)
		}
		close(out)
	}()
}

// SimpleWaitGroup Simple waitGroup
func SimpleWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(rand.Intn(100))
		}()
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

// FanInFanOut Fan in fan out
func FanInFanOut() {

}
