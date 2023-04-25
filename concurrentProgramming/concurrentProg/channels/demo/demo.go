package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// jobs <-chan Job: read only channel
// result chan<- Result: write only channel
// control chan ControlMsg: read and write channel
func doubler(jobs <-chan Job, result chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("doubler: exit")
				control <- ExitOk
				return
			default:
				panic("doubler: unknown control message")
			}
		case job := <-jobs:
			result <- Result{result: job.data * 2, job: job}
		}
	}

}

func main() {
	// job
	jobs := make(chan Job, 50)
	// results
	results := make(chan Result, 50)
	// control
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 50; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case msg := <-results:
			fmt.Println("result", msg)

		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout")
			control <- DoExit
			// wait until exit signal is received
			<-control
			fmt.Println("Program exit")
			return
		}
	}
}
