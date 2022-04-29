package main

import "fmt"

func ReverseString(input string)string{
	//тут даже комментировать нечего
	//итерируем по срезу []rune(strung), что бы не было проблем с символами юникода.
	tmp:=[]rune(input)
	//в цикле от 1 символа до середины строки мы меняем местами символы i и len-i
	//в отличии от варианта с циклом с декрементом,
	//мы не используем дополнительную переменную для append
	for i:=0;i<len(tmp)/2;i++{
		tmp[i],tmp[len(tmp)-i-1]=tmp[len(tmp)-i-1],tmp[i]
	}
	return string(tmp)
}

func prob19(){
	in:="Арифметик"
	fmt.Println("Входные данные: ",in)
	fmt.Println("Выходные данные:",ReverseString(in))
}