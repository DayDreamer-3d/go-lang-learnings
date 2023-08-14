// https://go.dev/tour/concurrency/2
// A channel is a communication object using which goroutines can communicate with each other.
// Technically, a channel is a data transfer pipe where data can be passed into or read from.
// Hence one goroutine can send data into a channel,
// while other goroutines can read that data from the same channel.

package main

import (
	"fmt"
	"strconv"
)

func pmt(text string, times int, c chan string) {
	for i := 0; i < times; i++ {
		msg := fmt.Sprintln(strconv.Itoa(i) + ": " + text)
		// send message through the channel
		c <- msg
	}
	close(c)
}

func main() {
	// create a channel
	c := make(chan string)

	go pmt("test A goroutines", 3, c)
	go pmt("test B goroutines", 3, c)

	// receive messages through the channel
	// is done by data := <- c or by the below for loop
	// note, receiving through a channel is a blocking call.
	for msg := range c {
		fmt.Println(msg)
	}

	// this program is not printing all the values
	// because channel is getting closed before all values have printed
	// it is getting closed by the first completed goroutine

}
