package main

import (
	"fmt"
)

var justString string

func StringSlice(input string,start int, end int)string{
	tmp:=[]rune(input)
	var res []rune
	for i:=start;i<end;i++{
		res=append(res,tmp[i])
	}
	return string(res)
}

func createHugeString(n int)string{
	str:=""
	for i:=0;i<n;i++{
		str+="♥"
	}
	return str
}
func someFunc() {
	v := ""
	v=createHugeString(1<<10)
	fmt.Printf("Первые 100 'символов' хьюдж стринга:\t\t%s\n",v[:100])
	justString = v[:100]
	fmt.Printf("Джаст стринг курильщика:\t\t\t\t\t%s\n",justString)
	//делаю вывод, что указанным выше способом мы работаем не со срезом символов, а со срезом байтов.
	//если символ занимает места больше, чем 1 байт, то программа выдаст результат, которого мы не ожидаем:)
	//значит надо работать со срезом символов
	justString=string([]rune(v)[:100])
	fmt.Printf("Джаст стринг здорового человека:\t\t\t%s\n",justString)
}

func prob15(){
	someFunc()
}

