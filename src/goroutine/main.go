package main

import (
	"fmt"
	"time"
)

func main() {
	/*for i := 0;i<10;i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}*/
	for i:=0;i<10;i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1*time.Second)
}
