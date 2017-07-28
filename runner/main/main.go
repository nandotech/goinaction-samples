/* This sample program demonstrates how to use a channel
to monitor amount of time the pgoram is runnning and terminate
the program if it runs too long. */
package main

import (
	"log"
	"os"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

//timeout is the number of seconds the program has to finish
const timeout = 3 * time.Second

//program entry point
func main() {
	log.Println("Starting Work")

	//Create a new timer for this run
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	//Run the tasks and handle the result
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")

}

//createTask returns an example task that sleeps for the specified
//number of seconds based on the id
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task# %d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
