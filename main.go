package main

import (
	"fmt"
	"go-practice/langintro"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// This WaitGroup will wait for 2 Goroutines to be done or to finish
	wg.Add(2)
	go langintro.SayHello(&wg)
	go langintro.SayWorld(&wg)

	// It gives the number of threads our program runs upon by default
	fmt.Printf("Threads: #%v \n", runtime.GOMAXPROCS(-1))
	// We can increase the number of threads by passing in any number we want the program to execute upon
	// "runtime.GOMAXPROCS(100)" will execute our program on 100 threads

	wg.Wait()
}
