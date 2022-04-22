package main

import (
	"fmt"
	"time"
)

func DelSliceElementOne(a []int,n int)[]int{
	//метод чуть быстрее
	delta:=time.Now()
	a[n-1],a[len(a)-1]=a[len(a)-1],a[n-1]
	a=a[:len(a)-1]
	fmt.Println(time.Since(delta).Nanoseconds())
	return a
}

func DelSliceElementTwo(a []int,n int)[]int{
	//метод чуть медленее
	delta:=time.Now()
	fmt.Println(a)
	a=append(a[:n-1],a[n:]...)
	fmt.Println(time.Since(delta).Nanoseconds())
	return a
}

func prob23(){
	a:=[]int{1,2,3,4,5,6,7}
	var b []int
	copy(b,a)
	n:=4
	fmt.Println("Входные данные: ",a)
	fmt.Println("Выходные данные:",DelSliceElementOne(a,n))
	fmt.Println("Выходные данные:",DelSliceElementTwo(a,n))
}
