package internal

import "fmt"

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
