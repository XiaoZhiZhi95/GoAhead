package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond*50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result,")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func othertask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond*100)
	fmt.Println("Task is done")
}

func main() {
	// fmt.Println(service())
	retch := AsyncService()
	othertask()
	fmt.Println(<-retch)
	time.Sleep(time.Millisecond*50)
}
