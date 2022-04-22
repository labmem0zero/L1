package main

import (
	"errors"
	"fmt"
)

func invertBit(n int, bitNum int)(int,error){
	bitMask:=1<<(bitNum-1)
	if len(fmt.Sprintf("%b",n))<len(fmt.Sprintf("%b",bitMask)){
		return n, errors.New(fmt.Sprintf("В двоичной записи числа %v(%b) нет %v разряда",n,n,bitNum))
	}
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
