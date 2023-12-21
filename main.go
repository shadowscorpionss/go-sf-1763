package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func intPrinter(wg *sync.WaitGroup) <-chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i + 1
			p := rand.Intn(100)
			time.Sleep(time.Duration(p) * time.Millisecond)

		}
	}()

	return c
}

func main() {
	wg := &sync.WaitGroup{}

	var ch <-chan int = intPrinter(wg)

	go func() {
		defer wg.Done()

		for v := range ch {
			fmt.Println(v)
		}

	}()

	wg.Wait()

}
