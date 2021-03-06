package main

import (
	"context"
	"fmt"
	"time"
)
type CancelFunc func()

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程退出，停止了...")
				return
			default:
				fmt.Println("协程运行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep( time.Second * 30)
	fmt.Println("两分钟时间到了，关闭子协程")
	cancel()
	time.Sleep( time.Second * 10)
	fmt.Println("结束")
}
