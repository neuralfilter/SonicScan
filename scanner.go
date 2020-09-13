package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func scan(targetIP string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		ipPort := fmt.Sprintf("%s:%d", targetIP, j)
		conn, err := net.DialTimeout("tcp", ipPort, 500*time.Millisecond)
		if err == nil {
			conn.Close()
			fmt.Println("Port " + strconv.Itoa(j) + " is open.")
		}
		results <- j
	}
}

func main() {
	start := time.Now()

	setWorkers := flag.Int("n_threads", 4, "Enter number of threads, this should be your CPU.")
	setTarget := flag.String("target", "8.8.8.8", "Enter target")
	flag.Parse()

	const numJobs = 9000               // set number of jobs to do
	jobs := make(chan int, numJobs)    // make channel
	results := make(chan int, numJobs) // make receiver

	for numWorkers := 1; numWorkers <= *setWorkers; numWorkers++ {
		go scan(*setTarget, jobs, results) // initialize 4 workers
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j // send numJobs amount of jobs

	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	elapsed := time.Since(start)
	log.Printf("Scan took %s", elapsed)
}
