package main

import "sync"

import "time"

import "math/rand"

import "fmt"

import "log"

import "errors"

func main() {
	//創造一個channel,使用bufferchannel,異部狀態,顯示各個Job完成狀態
	outChan := make(chan string, 100)
	//創造一個channel使用bufferchannel,異部狀態,發生錯誤時跳出func
	errChan := make(chan error, 100)
	//創造一個channel使用unbufferchannel,同步狀態,通知全部job已經完成
	//struct不佔任何記憶體,finish只要丟一個空的channel給他就可以了
	finishChan := make(chan struct{})
	//使用sync包中的WaitGroup,創造一個名為wg的任務對列結構,來將任務加入隊伍中
	wg := sync.WaitGroup{}
	//將20個goroutine加入wg隊伍中
	wg.Add(20)
	//使用for迴圈來執行go func
	for i := 0; i < 20; i++ {
		//執行20次 func 並將val,wg 任務隊伍,out channel放進去
		go func(val int, wg *sync.WaitGroup, out chan string, err chan error) {
			//
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			
			// 把輸出的結果丟到out channel
			out <- fmt.Sprintf("finished job id: %d", val)

			if val == 15 {
				err <- errors.New("fail job in 15")
			}
			//任務完成,將wg隊伍中的數量-1,其實就是wg.Add(-1)
			wg.Done()
		}(i, &wg, outChan, errChan)
	}

	//上面的for迴圈執行完以後,接著執行這個go func
	go func() {
		//發生阻塞,等待所有wg隊伍中的任務執行完成後就會解除
		wg.Wait()
		//close會執行到下面的Loop
		close(finishChan)
	}()

Loop:
	//用一個for迴圈去讀整個channel
	for {
		//因為要接多個channel,所以使用select的方式
		select {
			//把outChan 這個channel讀出來,並列印出log
		case out := <-outChan:
			log.Println(out)
		case err := <-errChan:
			log.Println(err)
			break Loop
			//遇到finishChan,跳出迴圈
		case <-finishChan:
			//這邊單用break只會跳出select迴圈,所以外面再包一層loop,並跳出loop
			break Loop
		case <-time.After(100 * time.Millisecond):
			log.Println("timeout")
			break Loop
		}
	}

}


/* 2019/11/18 02:24:29 finished job id: 2
2019/11/18 02:24:29 finished job id: 9
2019/11/18 02:24:29 finished job id: 3
2019/11/18 02:24:29 finished job id: 12
2019/11/18 02:24:29 finished job id: 19
2019/11/18 02:24:29 finished job id: 11
2019/11/18 02:24:29 finished job id: 18
2019/11/18 02:24:29 finished job id: 17
2019/11/18 02:24:29 finished job id: 14
2019/11/18 02:24:29 finished job id: 8
2019/11/18 02:24:29 finished job id: 4
2019/11/18 02:24:29 timeout */

