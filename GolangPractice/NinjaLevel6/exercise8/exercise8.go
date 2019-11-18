/*
Hands-on exercise #8
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
code: https://play.golang.org/p/c2HwqVE1Rd
*/


package main

import "fmt"

func foo() func() int {
	
	return func() int {
		return 42
	}

}


func main(){
	f := foo()
	fmt.Println(foo())
	//顯示0x48cfc0
	fmt.Println(f())
	//顯示 0 
	//

}



/*
 package main

	import "fmt"

	func foo() func() int {
		
		return func() int {
			return 42
		}

	}
	func foo2(s string) func() string {
		
		return func(s string) string {
			return s + "hello"
		}

	}

	func main(){
		f := foo()
		fmt.Println(foo())
		//顯示0x48cfc0
		fmt.Println(f())
		//顯示 0 
		//這裡用一個變數去接
		
		f2 := foo2()
		fmt.Println(foo2("llll"))
		//顯示0x48cfc0
		fmt.Println(f2("lllll"))

	}
	
	
	# command-line-arguments
	./main.go:14:10: cannot use func literal (type func(string) string) as type func() string in return argument
	./main.go:28:13: not enough arguments in call to foo2
		have ()
		want (string)
	./main.go:31:17: too many arguments in call to f2
		have (string)
		want ()
	
	*/