// https://talks.golang.org/2012/concurrency.slide
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("Denis!", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving.")
	quit <- true
}

func boring(msg string, quit chan bool) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				//
			case <-quit:
				return
			}
		}
	}()
	return c // Return the channel to the caller.
}
