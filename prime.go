package main

import (
	"fmt"
)

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, p int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%p != 0 {
				out <- n
			}
		}
	}()
	return out
}

func sieve() chan int {
	in := generate()
	out := make(chan int)
	go func() {
		for {
			prime := <-in
			out <- prime
			in = filter(in, prime)
		}
	}()

	return out
}

func main() {
	primes := sieve()
	for i := 0; i < 10; i++ {
		fmt.Println(<-primes)
	}
}
