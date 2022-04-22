package main

import (
	"fmt"
	"sync"
)

func prob2(){
	var wg sync.WaitGroup
	a:=[]int{2,4,6,8,10}
	sqr:= func(i int) {
		fmt.Printf("sqr(%v)=%v\n",i,i*i)
		wg.Done()
	}
	for _,i:=range a{
		wg.Add(1)
		go sqr(i)
	}
	wg.Wait()
}
