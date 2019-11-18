/* ##Hands-on exercise #3
##Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

 */
package main

import "fmt"

func main(){

	defer foo()
	
	fmt.Println("two")

}

func foo (){

	fmt.Println("foo")

}