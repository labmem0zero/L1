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
	var iterator MyIterator
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			for i:=0;i<count/10;i++{
				iterator.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(iterator.counter)
}

func prob18(){
	Iter(1000000)
}
