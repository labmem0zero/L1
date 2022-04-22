package main

import (
	"fmt"
)

func prob3(){
	a:=[]int{2,4,6,8,10}
	resChan:=make(chan int)
	sqr:=func(i int, c chan int) {
		c<-i*i
	}
	for _,i:=range a{
		go sqr(i,resChan)
	}
	sum:=0
	for i:=0;i<5;i++{
		val:=<-resChan
		sum+=val
	}
	fmt.Printf("Сумма вадратов=%v\n",sum)
}
