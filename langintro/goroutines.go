package langintro

import (
	"fmt"
	"sync"
)

// SayHello prints "Hello" to the CLI
func SayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

// SayWorld prints "Hello" to the CLI
func SayWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("World")
}

// Notes:
// Concurrency is just the ability of the application to work on multiple things at the same time
// We need to wait for one or more of the concurrent calculations to complete, so synchronization is essential
// WaitGroups & Mutexes are important, they are used for synchronization
// Concurrency is not parallelism
// But we can introduce parallelism in our apps as well
// Parallelism is generally a program that runs on multiple cores
// Whenever we write "go" in front of a function call we actually split the code on to a separate thread which is much more lightweight than the OS threads, thus go routines are much more easier to create and destroy
// In Go apps, there are thousands or even tens of thosands of goroutines executing at the same time
// Just like a thread, each goroutine has got an independent call stack
// Goroutines can also be declared via anonymous functions, but they need to be immediately invoked
// Whenever passing data into your Goroutines use arguments to do that and don't rely on golabl or closure scoped variables since they can change their value anytime from anywhere, arguments are pass by value so it won't result in unpredictable results
// WaitGroups are a great way to wait for the results for an ongoing execution of a Goroutine. It's designed to synchronize multiple Goroutines together
// We generally declare WaitGroups in package scope variables, since we want to declare them as done as soon as the execution is finished across multiple functions
// When we want to synchronize WaitGroups across packages, we pass them as pointer arguments
// Natively the order in which the Goroutines execute are out of our control, we can add them to WaitGroups to make sure they are all executed but the order of execution is in the control of the execution environment itself
//  So mutexes are important, in which we want to control the order of execution
// A Mutex is basically just a lock that the system is just going to honor, it's basically a lock and unlock mechanism. It ensures that only one thing can accesss or manipulate data at any given instance. We have to lock and unlock the resources manually. Just like we need to mark the WaitGroup as "defer wg.Done()"
// Goroutines are powerful and it can easily get out of hand
// To detect race conditions we use "go build -race"
// By default the Go compiler will use CPU threads equal to the number of cores available, we can change that with "runtime.GOMAXPROCS" we can increase the number of threads the program runs on, but doing that doesn't ensure efficiency, since the program might run into synchronization problems, too many threads can slow the system down
