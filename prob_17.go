package main

import "fmt"

func BinFind(in[]int,toFind int,delta int) (res int, ok bool){
	if len(in)<2{
		return 0,false
	}
	tmp:=in[len(in)/2-1]
	var tmpSlice []int
	if toFind==tmp{
		return delta+len(in)/2, true
	}
	if toFind<tmp{
		tmpSlice=in[:len(in)/2]
	}
	if toFind>tmp{
		tmpSlice=in[len(in)/2:]
		delta=delta+len(in)/2
	}
	return BinFind(tmpSlice,toFind,delta)
}

func prob17(){
	a:=[]int{1,8,15,21,27,36,39}
	toFind:=21
	fmt.Println(a)
	fmt.Println("toFind =",toFind)
	fmt.Println(BinFind(a,toFind,0))
}
