package main

import (
	"fmt"
)

// var wg sync.WaitGroup

func main() {

	even := make(chan int)
	odd := make(chan int)

	// wg.Add(1)
	go printEven(even)
	// wg.Add(1)
	go printOdd(odd)
	until := 100

	for i := until; i > 0; i-- {
		if i%2 == 0 {
			even <- i
			<-even
		} else {
			odd <- i
			<-odd
		}
	}

	close(even)
	close(odd)
	fmt.Println("Hello World")
	// wg.Wait()
}

func printEven(num chan int) {

	for {
		even, ok := <-num
		if !ok {
			// wg.Done()
			return
		}
		fmt.Println(even)
		num <- 1
	}
}

func printOdd(num chan int) {

	for {
		odd, ok := <-num
		if !ok {
			// wg.Done()
			return
		}
		fmt.Println(odd)
		num <- 1
	}
}
