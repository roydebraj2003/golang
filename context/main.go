package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Stopped worker : %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] Working ...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	rootCtx := context.Background()
	ctxCan, can1 := context.WithCancel(rootCtx)
	defer can1()
	ctxDed, can2 := context.WithTimeout(rootCtx, 4*time.Second)
	defer can2()
	go worker(ctxCan, "Worker-1")
	go worker(ctxDed, "Worker-2")

	time.Sleep(10 * time.Second)
	can1()
	can2()
}
