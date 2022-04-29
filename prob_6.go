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
	//остановка при помощи сигнального канала
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
	//по сути используется все тот же сигнальный канал
	//но у нас есть контекст, который может использоваться для различных целей
	//например для хранения данных о результатах выполнения функции
	//(например, хэш карта, реализованная в пакете 'gorilla/ctx')
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
