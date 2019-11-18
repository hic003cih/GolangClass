/*
Hands-on exercise #1
● Create a value and assign it to a variable.
● Print the address of that value.
code: https://play.golang.org/p/57kW8xd0qT
*/

package main

import "fmt"

func main() {
	
	x := 6
	y :=&x
	
	//&取地址,*取地址的值
	fmt.Println(&x)
	fmt.Println(*y)
}