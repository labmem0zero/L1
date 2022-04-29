package main

import (
	"fmt"
	"time"
)

func DelSliceElementOne(a []int,n int)[]int{
	//метод с перестановкой последнего элемента с n-ым элементом
	//ломает упорядоченность, но у нас остается доступ к стертому элементу
	//он находится в конце массива, с которого сделан срез.
	//Что бы достать удаленный элемент можно явно использовать массив среза:
	//arr:=(*[n]int)(a), где "n" - размер массива, "a" - срез, "arr" - массив, на котором основан срез
	delta:=time.Now()
	a[n-1],a[len(a)-1]=a[len(a)-1],a[n-1]
	a=a[:len(a)-1]
	fmt.Println(time.Since(delta).Nanoseconds())
	return a
}

func DelSliceElementTwo(a []int,n int)[]int{
	//метод без перестановок, порядок остается таки, как был
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
