package main

import (
	"fmt"
)

func Sort(in []int, left int, right int) []int {
	l := left
	r := right
	center := in[(left+right)/2]
	for l <= r {
		for in[r] > center {
			r--
		}
		for in[l] < center {
			l++
		}
		if l <= r {
			in[r], in[l] = in[l], in[r]
			l++
			r--
		}
	}
	if r > left {
		Sort(in, left, r)
	}
	if l < right {
		Sort(in, l, right)
	}
	return in
}

func prob16(){
	a:=[]int{5,2,7,1,8,8}
	fmt.Println(a)
	a=Sort(a,0,len(a)-1)
	fmt.Println(a)
}
