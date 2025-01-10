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
	in := gen(4, 10, 45)

	c1 := square(in)
	c2 := square(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int, len(nums))

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
