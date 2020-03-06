/**
 * Play with chennels.
 */

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// From https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/
func TestChennel() {
	fmt.Println("Test chennels.")
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` goroutines.
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data
	in <- 1
	in <- 2
	in <- 3

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	youFirst()
	meFirst()
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in
	result := num * 2
	out <- result
}

// From "The Behavior Of Channels" Listing 10
// URL: https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
func youFirst() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "You"
	}()

	select {
	case p := <-ch:
		fmt.Println("You first, work complete.", p)

	case <-ctx.Done():
		fmt.Println("You frist, moving on.")
	}
}

func meFirst() {
	duration := 201 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "Me"
	}()

	select {
	case p := <-ch:
		fmt.Println("Me first, work complete.", p)

	case <-ctx.Done():
		fmt.Println("Me first, moving on.")
	}
}
