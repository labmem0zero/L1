package main

import (
	"fmt"
	"sync"
)

type MyIterator struct{
	sync.Mutex
	counter int
}

func (mi *MyIterator) Inc(){
	mi.Lock()
	mi.counter++
	mi.Unlock()
}

func Iter(count int){
	//count - число, до которого нужно досчитать
	var iterator MyIterator
	var wg sync.WaitGroup
	wg.Add(10)
	//запускаем 10 воркеров, каждый инкрементит свою часть
	for i:=0;i<10;i++{
		go func() {
			for i:=0;i<count/10;i++{
				iterator.Inc()
			}
			wg.Done()
		}()
	}
	//если у нас есть остаток задач, после распределения по врокерам, то запускаем еще одного на остаток
	if count%10>0{
		wg.Add(1)
		go func() {
			for i:=0;i<count%10;i++{
				iterator.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(iterator.counter)
}

func prob18(){
	Iter(9999)
}
