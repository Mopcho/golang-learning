package main

import (
	"fmt"
	"time"
)

func concurrentFib(n int) {
	fibonacciChannel := make(chan int)
	go fibonacci(n, fibonacciChannel)

	for number := range fibonacciChannel {
		fmt.Printf("%v\n", number)
	}
}

// don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
