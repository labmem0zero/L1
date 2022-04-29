package main

import (
	"fmt"
	"strings"
)

func ReverseStringWords(input string)string{
	tmp:=strings.Split(input," ")
	//идентично предидущему заданию
	for i:=0;i<len(tmp)/2;i++{
		tmp[i],tmp[len(tmp)-i-1]=tmp[len(tmp)-i-1],tmp[i]
	}
	return strings.Join(tmp," ")
}

func prob20(){
	in:="Раз два три четыре пять"
	fmt.Println("Входные данные: ", in)
	fmt.Println("Выходные данные:", ReverseStringWords(in))
}
