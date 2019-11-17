package main

import "sync"

import "fmt"

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	ch := make(chan int)

	go func(ch <-chan int) {
		/* i := <-ch
		fmt.Println(i)
		i = <-ch */
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		i := 100
		ch <- i
		ch <- 101
		ch <- 102
		ch <- 103
		ch <- 104
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
