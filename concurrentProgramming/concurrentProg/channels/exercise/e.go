package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalc(i Job) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", i, duration)
	return int(i) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	//fmt.Println("len jobs =", len(jobs))
	for i := 0; i < 100; i++ {
		rann := rand.Intn(10000)
		fmt.Println("rann:", rann)
		jobs = append(jobs, Job(rann))
	}
	return jobs
}

func runJob(resultChan chan int, job Job) {
	resultChan <- longCalc(job)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := makeJobs()

	resultChan := make(chan int, 10)

	//fmt.Println("len jobs =", len(jobs))
	for i := 0; i < len(jobs); i++ {
		go runJob(resultChan, jobs[i])
	}

	resultCount := 0
	sum := 0
	for {
		result := <-resultChan
		sum += result
		resultCount++
		if resultCount == len(jobs) {
			break
		}
	}
	fmt.Println("Sum:", sum)
}
