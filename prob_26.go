package main

import (
	"fmt"
	"strings"
)

func CheckRepeat(in string)bool{
	//сначала привожу все к маленьким буквам, что бы не быть регистрозависимым
	in=strings.ToLower(in)
	//работаю со срезом рон, что бы учесть символы размером больше 1 байта(юникод)
	inRune:=[]rune(in)
	chkMap:=make(map[rune]int)
	//если символ встречается больше 1 раза, то будет выявлен во время итерирования.
	for _,c:=range inRune{
		chkMap[c]++
		if chkMap[c]>1{
			return true
		}
	}
	return false
}

func prob26(){
	input:=[]string{
		"abcd",
		"abba",
		"SAaB",
		"Oops!",
	}
	fmt.Println("Входной срез:",input)
	for _,s:=range input{
		fmt.Printf("%s\t-\t%v\n",s,CheckRepeat(s))
	}
}
