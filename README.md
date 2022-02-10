# Fibonacci optimization

This program shows how go can do concurrency tasks on a easy way to improve the performance of this tasks.

## How the program do that?

### Worker
First we have to do a Worker function that have the responsability exec the Fibonacci function and insert the value into a results channel.

### Main
In the main function there are tasks that is an slice of numbers, it represents tasks because are values that will be used to calculate the Fibonacci series of that number, then there are two channels of integer numbers, one is for the jobs and the other is for the results of the jobs (both have the same size). 
Then, we do an iteration over the number of workers(beause we can choose how many workers we want to do the tasks) and do a gorutine that  call the function Work with te id of worker, the jobs and results. Next there is a iteration over the task to assing the value of the task to the jobs channel, then the jobs channel is closed. Finally we do an iteration over the length of the task to wait or obtain the values of the results channel.

## Commands
### Run without go rutines and workers implementation
```
$ go run main.go
```

### Run with go rutines and workers implementation (default 12 workers)
```
$ WORKERS=true go run main.go
```

### Run with go rutines, workers implementation and a custom number of workers
```
$ WORKERS=true NWORKERS=<Number of workers> go run main.go
```