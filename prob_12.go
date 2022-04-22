package main

import "fmt"

func SliceToSet(input []string)[]string{
	temp:=make(map[string]bool)
	for _,i:=range input{
		temp[i]=true
	}
	var res []string
	for k,_:=range temp{
		res=append(res,k)
	}
	return res
}

func prob12(){
	a:=[]string{"cat","dog","cow","cat"}
	fmt.Println("Входные данные:        ",a)
	fmt.Println("Результирующие данные: ",SliceToSet(a))
}
