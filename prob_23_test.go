package main

import (
	"fmt"
	"testing"
)

type testpair struct {
	a []int
	n int
}
var in=testpair{[]int{1,2,3,4,5,6,7},4}

func TestDelSliceElementOne(t *testing.T) {
	for i:=1;i<1000;i++{
		a,n:=in.a,in.n
		a[n-1],a[len(a)-1]=a[len(a)-1],a[n-1]
		a=a[:len(a)-1]
	}

}

func TestDelSliceElementTwo(t2 *testing.T) {
	for i:=1;i<1000;i++ {
		a,n:=in.a,in.n
		fmt.Println(a)
		a=append(a[:n-1],a[n:]...)
	}

}
