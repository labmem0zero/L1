package main

import (
	"fmt"
	"sync"
)

func prob2(){
	var wg sync.WaitGroup
	a:=[]int{2,4,6,8,10}
	fmt.Println("Начальные данные:",a)
	ch:=make(chan int)
	//горутина отправляет данные в канал
	go func(){
		for _,i:=range a{
			ch<-i
		}
		close(ch)
	}()
	wg.Add(2)
	//первый воркер конкурентно читает данные с канала
	go func() {
		for i:=range ch{
			fmt.Printf("sqr(%v)=%v\n",i,i*i)
		}
		wg.Done()
	}()
	//второй воркер конкурентно читает данные с канала
	go func() {
		for i:=range ch{
			fmt.Printf("sqr(%v)=%v\n",i,i*i)
		}
		wg.Done()
	}()
	wg.Wait()
}
