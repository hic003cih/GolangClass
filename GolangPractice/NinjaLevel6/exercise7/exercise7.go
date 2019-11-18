/*
Hands-on exercise #7
● Assign a func to a variable, then call that func
code: https://play.golang.org/p/_Qu7ZAyFDH

*/


package main

import "fmt"

var foo = func () {
	for i := 0; i < 3; i++ {
					fmt.Println(i)
				}
}


func main(){

	foo()

	g := foo

	g()
}
