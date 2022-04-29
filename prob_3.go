package main

import (
	"fmt"
	"sync"
)

//структура с мьютексом, показываю навык работы с конкурентной записью
//результаты мог суммировать в одной горутине через канал
type sum struct {
	int
	sync.Mutex
}

func prob3(){
	var wg sync.WaitGroup
	var res sum
	a:=[]int{2,4,6,8,10}
	fmt.Println("Начальные данные:",a)
	ch:=make(chan int)
	//горутина отправляет данные в канал
	go func(){
		for _,i:=range a{
			ch<-i*i
		}
		close(ch)
	}()
	wg.Add(2)
	//первый воркер конкурентно читает данные с канала и прибавляет к результату
	go func() {
		for i:=range ch{
			res.Lock()
			res.int+=i
			res.Unlock()
		}
		wg.Done()
	}()
	//второй воркер конкурентно читает данные с канала и прибавляет к результату
	go func() {
		for i:=range ch{
			res.Lock()
			res.int+=i
			res.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Сумма квадратов:",res.int)
}
