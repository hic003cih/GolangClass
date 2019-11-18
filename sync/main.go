package main

import "fmt"

import "time"

import "sync"

<<<<<<< HEAD
//執行do func , 並將 i 和 w 任務隊伍傳入 
func do(i int, wg *sync.WaitGroup) {
	//使用%d 來顯示 int
	fmt.Printf("start job: %d\n", i)
	//延遲1秒來讓任務執行
	time.Sleep(1 * time.Second)
	//使用%d 來顯示 int
	fmt.Printf("end job: %d\n", i)
	//任務完成,將wg隊伍中的數量-1,其實就是wg.Add(-1)
=======
func do(i int, wg *sync.WaitGroup) {
	fmt.Printf("start job: %d\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("end job: %d\n", i)
>>>>>>> origin/master
	wg.Done()
}

func main() {
<<<<<<< HEAD
	//使用sync包中的WaitGroup,創造一個名為wg的任務對列結構,來將任務加入隊伍中
	wg := sync.WaitGroup{}
	//將3個gorountine加入wg隊伍中
=======
	wg := sync.WaitGroup{}
>>>>>>> origin/master
	wg.Add(3)
	go do(1, &wg)
	go do(2, &wg)
	go do(3, &wg)

	wg.Wait()
	fmt.Println("done")
	//time.Sleep("Done!")
}
