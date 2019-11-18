/*
Hands-on exercise #10
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
code: https://play.golang.org/p/a56uWtoFSL
*/

  
  package main

	import "fmt"

  func main(){
    g :=loop()
    loop()
	  f :=loop()
    
	  fmt.Println(g())
	  fmt.Println(g())
	  fmt.Println(g())
	  
	  fmt.Println(f())
	  fmt.Println(f())
  }

func loop () func() int {
       x :=0
    
      return func() int {
      x++
      return x
    }
  }