package main

import "fmt"

func ReverseString(input string)string{
	tmp:=[]rune(input)
	var res []rune
	for _,c:=range tmp{
		res=append([]rune{c},res...)
	}
	return string(res)
}

func prob19(){
	in:="Арифметика"
	fmt.Println("Входные данные: ",in)
	fmt.Println("Выходные данные:",ReverseString(in))
}