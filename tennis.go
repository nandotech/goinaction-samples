package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

//main program entry point
func main() {
	//unbuffered chan
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	//start the set
	court <- 1

	//wait for game to finish
	wg.Wait()
}

//player simulates person playing tennis
func player(name string, court chan int) {
	defer wg.Done()

	for {
		//Wait for the ball to be hit back to us
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}

		//pick random number and see if we miss
		n := rand.Intn(100)
		if n%999 == 0 {
			fmt.Printf("Player %s Missed!\n", name)
			//Close channel to signal loss/endgame
			close(court)
			return
		}
		//Display and then increment the hit count by one
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		// Hit the ball back to opponent
		court <- ball
	}
}
