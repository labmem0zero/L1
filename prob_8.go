package main

import (
	"errors"
	"fmt"
)

func invertBit(n int, bitNum int)(int,error){
	//создаем битовую маску, что бы получить число с битом=1 только на той позиции,
	//которую нужно инвертировать.
	bitMask:=1<<(bitNum-1)
	//проверяем, что бы в битовой записи входного числа имелся необходимый разряд
	//можно было сделать по другому: n<=pow(2,bitnum)
	if len(fmt.Sprintf("%b",n))<len(fmt.Sprintf("%b",bitMask)){
		//Показываю, что умею формировать кастомные ошибки
		return n, errors.New(fmt.Sprintf("В двоичной записи числа %v(%b) нет %v разряда",n,n,bitNum))
	}
	//пользуемся исключающим ИЛИ, которое работает так:
	//Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
	return n^bitMask, nil
}

func prob8(){
	n:=321
	i:=6
	fmt.Printf("До смены бита:\t\t%v=%b\n",n,n)
	fmt.Printf("Бит для инверсии: %v\n",i)
	res,err:=invertBit(n,i)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("После смены бита: \t%v=%b\n",res,res)
}
