package main

import (
	"fmt"
	"sync"
)

func main() {
	chList := make([]chan bool, 100)
	for i := range chList {
		chList[i] = make(chan bool, 1)
	}
	chList[0] <- true
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		curChan := chList[i]

		go func(i int) {
			wg.Add(1)
			select {
			case <-curChan:
				fmt.Println(i)
				if i >= 99 {
					wg.Done()
					return
				}
				chList[i+1] <- true
				wg.Done()
				return
			}
		}(i)
	}
	wg.Wait()
}
