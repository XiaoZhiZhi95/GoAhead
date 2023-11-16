package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var mut sync.Mutex
	wg := sync.WaitGroup{}
	for i:=0;i<5000;i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	// time.Sleep(1*time.Second)
	wg.Wait()
	fmt.Println("count = ", count)
}
