package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("main start")
	ch := make(chan struct{})
	for i:=0; i<5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCanceled(ch) {
					break
				}
				time.Sleep(time.Millisecond*50)
			}
			fmt.Println(i, "canceled")
		}(i, ch)
	}
	// cancel_2(ch)
	cancel_1(ch)
	time.Sleep(1*time.Second)
	fmt.Println("main end")
}

func isCanceled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{})  {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{})  {
	close(ch)
}
