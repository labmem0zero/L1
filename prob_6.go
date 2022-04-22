package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func prob6(){
	var wg sync.WaitGroup
	stopChan:=make(chan int)
	wg.Add(1)
	go func(stop <-chan int){
		for{
			select {
			case <-stop:
				fmt.Println("Горутина остановлена сигнальным каналом")
				wg.Done()
				return
			default:
				fmt.Println(time.Now())
				time.Sleep(100*time.Millisecond)
			}
		}
	}(stopChan)
	time.Sleep(1*time.Second)
	close(stopChan)

	//Остановка при помощи cancel() функции контекста
	wg.Add(1)
	ctx,cancel:=context.WithCancel(context.Background())
	go func(ctx context.Context){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена cancel() функцией")
				wg.Done()
				return
			default:
				fmt.Println(time.Now())
				time.Sleep(100*time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(1*time.Second)
	cancel()

	//Остановка при помощи таймаут-конекста
	wg.Add(1)
	toCtx,_:=context.WithTimeout(context.Background(), 1*time.Second)
	go func(ctx context.Context){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена таймаутом контекста")
				wg.Done()
				return
			default:
				fmt.Println(time.Now())
				time.Sleep(100*time.Millisecond)
			}
		}
	}(toCtx)
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
