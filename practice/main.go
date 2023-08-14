package main

import (
	"fmt"
)

func say(ch chan string) {
	fmt.Println(<-ch)

}

func main() {
	ch := make(chan string, 3)

	ch <- "first"
	ch <- "second"
	close(ch)
	go say(ch)
	for n := range ch {
		fmt.Println(n)
	}
	fmt.Scan()

}
