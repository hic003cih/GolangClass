package main

import "sync"

import "time"

import "math/rand"

import "fmt"

import "log"

import "errors"

func main() {
	//創造一個channel,使用bufferchannel,異部狀態
	outChan := make(chan string, 100)
	//創造一個channel使用bufferchannel,異部狀態
	errChan := make(chan error, 100)
	//創造一個channel使用unbufferchannel,同步狀態
	finishChan := make(chan struct{})
	//使用sync包中的WaitGroup,創造一個名為wg的任務對列結構,來將任務加入隊伍中
	wg := sync.WaitGroup{}
	//將20個goroutine加入wg隊伍中
	wg.Add(20)
	//使用for迴圈來執行go func
	for i := 0; i < 20; i++ {
		//執行func 並將
		go func(val int, wg *sync.WaitGroup, out chan string, err chan error) {
			//
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			
			out <- fmt.Sprintf("finished job id: %d", val)

			if val == 15 {
				err <- errors.New("fail job in 15")
			}
			//任務完成,將wg隊伍中的數量-1,其實就是wg.Add(-1)
			wg.Done()
		}(i, &wg, outChan, errChan)
	}

	go func() {
		//發生阻塞,等待所有wg隊伍中的任務執行完成後就會解除
		wg.Wait()
		close(finishChan)
	}()

Loop:
	for {

		select {
		case out := <-outChan:
			log.Println(out)
		case err := <-errChan:
			log.Println(err)
			break Loop
		case <-finishChan:
			break Loop
		case <-time.After(100 * time.Millisecond):
			log.Println("timeout")
			break Loop
		}
	}

}


2019/11/18 02:24:29 finished job id: 2
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
2019/11/18 02:24:29 timeout

