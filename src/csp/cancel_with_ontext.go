package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	fmt.Println("main start")
	rootCtx := context.Background()	// 当前main为根节点
	ctx, cancelFunc := context.WithCancel(rootCtx)	// ctx为子节点的context，传给下面的协程， cancelFunc执行时取消
	for i:=0; i<5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceledWithContext(ctx) {
					break
				}
				time.Sleep(time.Millisecond*50)
			}
			fmt.Println(i, "canceled")
		}(i, ctx)
	}
	// 根节点取消context
	cancelFunc()
	time.Sleep(1*time.Second)
	fmt.Println("main end")
}

func isCanceledWithContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():	// 通过ctx.Done()接收cancel的通知
		return true
	default:
		return false
	}
}