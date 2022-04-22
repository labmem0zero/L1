package main

import "fmt"

func intSliceToSet(input []int)[]int{
	temp:=make(map[int]bool)
	for _,i:=range input{
		temp[i]=true
	}
	var res []int
	for k,_:=range temp{
		res=append(res,k)
	}
	return res
}

func Intersect(setOne[]int, setTwo[]int)[]int{
	tmpOne:=intSliceToSet(setOne)
	tmpTwo:=intSliceToSet(setTwo)
	temp:=make(map[int]int)
	for _,i:=range tmpOne{
		temp[i]=temp[i]-1
	}
	for _,i:=range tmpTwo{
		temp[i]=temp[i]+1
	}
	var res []int
	for k,v:=range temp{
		if v==0{
			res=append(res,k)
		}
	}
	return res
}

func prob11(){
	a:=[]int{1,2,3,4,5}
	b:=[]int{4,5,6,7,8}
	fmt.Println("Множество A =",a)
	fmt.Println("Множество B =",b)
	fmt.Println("Пересечение =",Intersect(a,b))
}
