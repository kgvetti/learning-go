// 
// filter primes, demo using channels and go-routines
// from: http://golang.org/doc/go_tutorial.html

package main

import "fmt"

// generates numbers and feeds them into a channel whcih is returned
func generate() chan int {
	ch := make(chan int);
	// inline fucntion as a goroutine
	go func() {
		for i:= 2; ; i++ {
			ch <- i
		}
	}();
	return ch;
}

func filter(in chan int, prime int) chan int {
	out := make(chan int);
	// inline function as a goroutine
	go func() {
		for {
			if i:= <-in; i%prime != 0 {
				out <- i;
			}
		}
	}();
	return out;
}

// sieve function, returns a channel of ints full of primes
// processes filters, creating a new one for each prime found
func sieve() chan int {
	out := make(chan int);
	// inline function as a goroutine
	go func() {
		ch := generate();
		for {
			prime := <-ch;
			out <- prime;
			ch = filter(ch, prime);
		}
	}();
	return out;
}

// entry point, runs forever, starts the sieving process and prints (consumes) the results
func main() {
	primes := sieve();
	for {
		fmt.Println(<-primes);
	}
}