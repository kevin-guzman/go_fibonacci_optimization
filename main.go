package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d has finished the job %d, and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	start := time.Now()
	tasks := []int{2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 17, 10, 40, 6, 22, 2, 5, 2, 34, 36, 22, 41, 44}
	withGoRutinesAndWorkers, err := strconv.ParseBool(os.Getenv("WORKERS"))
	if err != nil {
		withGoRutinesAndWorkers = false
	}
	nWorkers, err := strconv.Atoi(os.Getenv("NWORKERS"))
	if err != nil {
		nWorkers = 12
	}
	fmt.Printf("Workers impmented %v, number of workers %d\n", withGoRutinesAndWorkers, nWorkers)
	if withGoRutinesAndWorkers {
		jobs := make(chan int, len(tasks))
		results := make(chan int, len(tasks))

		for i := 0; i <= nWorkers; i++ {
			go Worker(i, jobs, results)
		}

		for _, value := range tasks {
			jobs <- value
		}
		close(jobs)

		for i := 0; i < len(tasks); i++ {
			<-results
		}
	} else {

		for _, v := range tasks {
			val := Fibonacci(v)
			fmt.Printf("%d has the Fibonacci series %d\n", v, val)
		}

	}

	elapsed := time.Since(start)
	fmt.Printf("Program took %s", elapsed)

}
