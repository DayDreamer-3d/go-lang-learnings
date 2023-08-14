// https://go.dev/tour/concurrency/1
// A goroutine is a lightweight thread managed by the Go runtime.
// used in golang for concurrency
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func pmt(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(strconv.Itoa(i), ": ", text)
	}
}

func main() {
	// WaitGroup is used to sync goroutines
	var wg sync.WaitGroup
	// Add one go routine to wait group
	wg.Add(1)
	// go defines goroutine
	// note, this function is anonymous
	go func() {
		pmt("test A goroutines", 3)
		wg.Done()
	}()
	pmt("test B goroutines", 3)
	// Blocking wait call until all goroutine finishes
	wg.Wait()

}
