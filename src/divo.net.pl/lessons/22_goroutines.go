// A _goroutine_ is a lightweight thread of execution.

package lessons

import(
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	
	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")
	
	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")
	
	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	// We must wait for some time for asynchronyous
	// Without this program will exit
	// and nothing print on screen
	time.Sleep(time.Duration(1)*time.Second)
}