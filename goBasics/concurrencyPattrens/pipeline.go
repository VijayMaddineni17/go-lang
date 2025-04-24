package main

import "fmt"

// sliceToChannel takes a slice of integers and returns a channel that emits those integers one by one
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int) // unbuffered channel

	go func() {
		// Send each number to the channel
		for _, n := range nums {
			out <- n // blocks until receiver is ready (in sq function)
		}
		close(out) // close channel after sending all values
	}()
	return out // return the channel to the caller
}

// sq takes a channel of integers, squares each number, and returns a channel with the squared values
func sq(in <-chan int) <-chan int {
	out := make(chan int) // unbuffered channel for squared results

	go func() {
		// Read from input channel until it's closed
		for n := range in {
			out <- n * n // compute square and send to output channel
		}
		close(out) // close output channel once all input is processed
	}()
	return out
}

func main() {
	nums := []int{1, 4, 2, 6, 7}

	// Stage 1: Convert slice to a channel of integers
	dataChannel := sliceToChannel(nums)

	// Stage 2: Square each number from dataChannel
	finalChannel := sq(dataChannel)

	// Stage 3: Print each squared value from finalChannel
	for n := range finalChannel {
		fmt.Println(n) // receive and print values until channel is closed
	}
}
