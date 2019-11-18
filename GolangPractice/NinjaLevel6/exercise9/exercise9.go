/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
code: https://play.golang.org/p/0yGYPKh1y7
*/


/*

这是一个简单的回调例子，调用函数test时，调用真正的实现函数add

 package main

import "fmt"

type Callback func(x, y int) int

// 提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
    return callback(x, y)
}

// 回调函数的具体实现
func add(x, y int) int {
    return x + y
}

func main() {
    x, y := 1, 2
    fmt.Println(test(x, y, add))
}
*/

/*
这是一个将字符串转为Int的例子，在转换失败的情况下执行回调函数，输出错误信息

package main

import (
    "strconv"
    "fmt"
)

type Callback func(msg string)

//将字符串转换为int64，如果转换失败调用Callback
func stringToInt(s string, callback Callback) int64 {
    if value, err := strconv.ParseInt(s, 0, 0); err != nil {
        callback(err.Error())
        return 0
    } else {
        return value
    }
}

// 记录日志消息的具体实现
func errLog(msg string) {
    fmt.Println("Convert error: ", msg)
}

func main() {
    fmt.Println(stringToInt("18", errLog))
    fmt.Println(stringToInt("hh", errLog))
}

*/

/*
package main

import (
	"fmt"
	"strconv"
)

func foo(x,y int) int {
	return x+y
}

func foo2(x,y int,callback func(x,y int) int) string {

		sum := callback(x,y)
		
		return "hello" + strconv.Itoa(sum)
}

func main() {
	
	temp := foo2(5,4,foo(5,4))
	
	fmt.Println(temp)
}

# command-line-arguments
./main.go:21:22: cannot use foo(5, 4) (type int) as type func(int, int) int in argument to foo2

*/



package main

import (
	"fmt"
	"strconv"
)

func foo(x,y int) int {
	return x+y
}

func foo2(x,y int,callback func(x,y int) int) string {

		sum := callback(x,y)
		
		return "hello" + strconv.Itoa(sum)
}

func main() {
	
	temp := foo2(5,4,foo)
	
	fmt.Println(temp)
}
