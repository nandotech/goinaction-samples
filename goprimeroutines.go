package main

import (
	"fmt"
	"runtime"
	"sync"
)

//global variable wg toblock and allow program to finish
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

//printprime displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	//Schedule defered call to Done() -- code readability
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
}
