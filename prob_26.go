package main

import (
	"fmt"
	"strings"
)

func CheckRepeat(in string)bool{
	in=strings.ToLower(in)
	inRune:=[]rune(in)
	chkMap:=make(map[rune]int)
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
