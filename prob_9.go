package main

import (
	"fmt"
	"sync"
)

func prob9(){
	var wg sync.WaitGroup
	in:=make(chan int)
	out:=make(chan int)
	wg.Add(3)
	//пишем в канал данные из массива
	go func(in chan int) {
		x:=[]int{1,2,3,4,5,6,7,8,9,10}
		fmt.Printf("Входные данные: %v\n",x)
		fmt.Printf("Выходные данные: ")
		for _,n:=range x{
			in<-n
		}
		close(in)
		wg.Done()
	}(in)
	//промежуточная горутина, проводит вычисления и отправляет результаты в канал результатов
	go func(in chan int, out chan int) {
		for x:=range in{
			out<-x*2
		}
		close(out)
		wg.Done()
	}(in, out)
	//считываем данные из канала результатов и выводим их на экран
	go func(out chan int) {
		for x:=range out{
			fmt.Printf("%v ",x)
		}
		//после считывания результатов делаем перенос строки(для красоты)
		fmt.Println()
		wg.Done()
	}(out)
	wg.Wait()
}
