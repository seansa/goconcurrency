package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
	pool        int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	startTime := time.Now()
	go allocate(100) //N Jobs
	done := make(chan bool)
	go result(done)
	createWorkerPool(10) //N Workers
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup, pool int) {
	for job := range jobs {
		output := Result{job, digits(job.randomno), pool}
		results <- output
	}
	wg.Done()
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(10 * time.Millisecond)
	return sum
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d on pool %v, input random no %d, sum of digits %d\n", result.job.id, result.pool, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
