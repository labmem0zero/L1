package main

import (
	"fmt"
	"strings"
)

func ReverseStringWords(input string)string{
	tmp:=strings.Split(input," ")
	var res []string
	for _,s:=range tmp{
		res=append([]string{s},res...)
	}
	return strings.Join(res," ")
}

func prob20(){
	in:="Раз два три четыре пять"
	fmt.Println("Входные данные: ", in)
	fmt.Println("Выходные данные:", ReverseStringWords(in))
}
