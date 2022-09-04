/*
Example of concurrency

GO routines are function that can run at the same time as other functions by utilising multiple threads on a CPU

They are better described as green threads, which dynamically map function execution to operating system threads
as needed. Go routines might be executed asynchronously on the same operating system thread, which is especially
useful in the case that they perform IO operations

We can create many simultanious Go routines as Go does a very good job with concurrency,
but it won't make the program any faster as we are constrained by how many calls our CPU has.

It's often we will see a call for user input (fmt.Scanln()) at the end of a function with many Go routines. This
is to stop the function from terminating giving our Go routines time to execute.
A better solution for this is to use sync which will keep track of the go routines running via
counter and then will exit the program once complete.

A channel is a way for go routines to communicate with each other. We can accept a channel as an argument which is
like a pipe that allows you to send or recieve a message

The following five examples are found below:

- Example of basic concurrency
- Channel - Deadlock and channel closing
- Non blocking buffered channel (Sending to a channels is a blocking operation)
- Select Statement
- Worker pool pattern
*/

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*

Example of basic concurrency
---

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Increment by one because we have one Go routine to wait for.

	// Anon function that calls itself
	go func() {
		count("sheep")
		wg.Done()
	}()

	// if any Go routines are still running it will wait
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}

############################################################

Channel - Deadlock and channel closing
---

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		// Will keep recieving messages and putting then in msg until the channel is closed
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // An arrow pointing into the channel name will send a message
		time.Sleep(time.Millisecond * 500)
	}

	close(c)

}


############################################################


Non blocking buffered channel (Sending to a channels is a blocking operation)
---

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}


############################################################


Select Statement
---

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		//  Use select statement to recieve from whichever channel is ready
		// Otherwise the slower one will block teh faster one
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}


############################################################


Worker pool pattern
---
This is where you have a queue of work to be done and multiple concurent workers pulling items of fof the queue

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}


*/
