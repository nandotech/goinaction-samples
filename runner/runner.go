package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner runs a set of tasks within a given timeout and can be
//shut down on an os interrupt
type Runner struct {
	//interrupt channel reports a signal from os
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

//ErrTimeout returned when value received on timeout
var ErrTimeout = errors.New("received timeout")

//ErrInterrupt is returned when an event from the OS is received
var ErrInterrupt = errors.New("received interrupt")

//New returns a new ready-to-use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add attaches tasks to the runner. Task is a function that takes int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start runs all tasks and monitors channel events
func (r *Runner) Start() error {
	//We want to receive all interrupt based signals
	signal.Notify(r.interrupt, os.Interrupt)

	//Run the different tasks on separate goroutines
	go func() {
		r.complete <- r.run()
	}()
	select {
	//signaled when processing is done
	case err := <-r.complete:
		return err
	//signaled wen out of time
	case <-r.timeout:
		return ErrTimeout
	}
}

//run executes each registered task
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//check for interrupt
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//Execute the registered task
		task(id)
	}
	return nil
}

//gotInterrupt verifies if the interrupt signal has been issued
func (r *Runner) gotInterrupt() bool {
	select {
	//signaled when interrupt event sent
	case <-r.interrupt:
		//stop receiving futher signa
		signal.Stop(r.interrupt)
		return true
		//Continue running as normal
	default:
		return false
	}
}
