/*
Documentation
They're called goroutines because the existing terms—threads, coroutines, processes, and so
on—convey inaccurate connotations. A goroutine has a simple model: it is a function executing
concurrently with other goroutines in the same address space.
code:
● https://play.golang.org/p/lBKFKCwrue

*/

package main

import (
"fmt"
)

func doSomething(x int) int {
return x * 5
}

func main() {
ch := make(chan int)
go func() {
ch <- doSomething(5)
}()
fmt.Println(<-ch)
}