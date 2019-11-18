/*
Race condition
Here is a picture of the race condition we are going to create. Race conditions are not good
code. A race condition will give unpredictable results. We will see how to fix this race
condition in the next video.
code: https://play.golang.org/p/FYGoflKQej
video: 134


*/


package main

import (
"fmt"
"runtime"
"sync"
)

func main() {
fmt.Println("CPUs:", runtime.NumCPU())
fmt.Println("Goroutines:", runtime.NumGoroutine())

counter := 0

const gs = 100
var wg sync.WaitGroup
wg.Add(gs)

for i := 0; i < gs; i++ {
	go func() {
		v := counter
		// time.Sleep(time.Second)
		runtime.Gosched()
		v++
		counter = v
		wg.Done()
	}()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
wg.Wait()
fmt.Println("Goroutines:", runtime.NumGoroutine())
fmt.Println("count:", counter)
}