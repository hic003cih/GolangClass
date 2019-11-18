/*
WaitGroup
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to
set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when
finished. At the same time, Wait can be used to block until all goroutines have finished. Writing
concurrent code is super easy: all we do is put “go” in front of a function or method call.
● runtime.NumCPU()
● runtime.NumGoroutine()
● sync.WaitGroup
○ func (wg *WaitGroup) Add(delta int)
○ func (wg *WaitGroup) Done()
○ func (wg *WaitGroup) Wait()
● personal note about teaching on Greater Commons
race conditions picture:
● https://goo.gl/JSC7M2
code:
● starting code:
○ https://play.golang.org/p/bnI0akWF9f
● finished:
○ https://play.golang.org/p/VDioqpZ65

*/

//finished code

package main

import (
"fmt"
"runtime"
"sync"
)

var wg sync.WaitGroup

func main() {
fmt.Println("OS\t\t", runtime.GOOS)
fmt.Println("ARCH\t\t", runtime.GOARCH)
fmt.Println("CPUs\t\t", runtime.NumCPU())
fmt.Println("Goroutines\t", runtime.NumGoroutine())

wg.Add(1)
go foo()
bar()

fmt.Println("CPUs\t\t", runtime.NumCPU())
fmt.Println("Goroutines\t", runtime.NumGoroutine())
wg.Wait()
}

func foo() {
for i := 0; i < 10; i++ {
fmt.Println("foo:", i)
}
wg.Done()
}

func bar() {
for i := 0; i < 10; i++ {
fmt.Println("bar:", i)
}
}