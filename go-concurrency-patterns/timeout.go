// https://talks.golang.org/2012/concurrency.slide
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Ann"), boring("Joe")) // Function returning a channel.
	// Create the timer once, outside the loop, to time out the entire conversation.
	// (In the previous program, we had a timeout for each message.)
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Printf("You say: %q\n", s)
		case <-timeout:
			fmt.Println("You talk too much!")
			return
		}
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
