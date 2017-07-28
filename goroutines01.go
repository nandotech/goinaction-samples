package main

import (
	"fmt"
	"runtime"
	"sync"
)

//main function
func main() {
	//Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(2)

	//wg is used to wait for the program to finish.
	//Add a count of two, one for each goroutine.

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start  Goroutines")

	//declare anon function and create a goroutine
	go func() {
		//schedule the call to Done to tell main we are done
		defer wg.Done()

		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			fmt.Printf("\n")
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//declare anon function and create a goroutine
	go func() {
		//schedule the call to Done to tell main we are done
		defer wg.Done()

		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
