package main

import (
	"fmt"
	"runtime"
	"time"
)

func getResult(i int) string {
	time.Sleep(10*time.Millisecond)
	return fmt.Sprintf("The result is from %d", i)
}

func main() {
	fmt.Println("before num = ", runtime.NumGoroutine())
	// ch := make(chan string)	// 会导致内存泄漏，下面只拿走了一次，其他的协程放不进去，就会一直等待，after = 10
	ch := make(chan string, 10)	// 避免内存泄漏，让所有的协程都可以将结果放到ch中，不会被阻塞等待，after = 1
	// 起了10个协程去拿东西，
	for i:=0;i<10;i++ {
		go func(i int) {
			ret := getResult(i)
			ch <- ret
		}(i)
	}
	result := <-ch
	fmt.Println(result)
	time.Sleep(100*time.Millisecond)
	fmt.Println("after num = ", runtime.NumGoroutine())
}
