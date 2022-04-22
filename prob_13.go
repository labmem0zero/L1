package main

import "fmt"

func prob13(){
	a:=10
	b:=5
	fmt.Printf("Начальные данные:\t\t\t\tA=%v B=%v\n",a,b)
	a,b=b,a
	fmt.Printf("Сахарносинтаксический способ:\tA=%v B=%v\n",a,b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Printf("Арифметический способ:\t\t\tA=%v B=%v(возвращаем обратно)\n",a,b)
	a=a^b
	b=a^b
	a=a^b
	fmt.Printf("Бинарнологический способ:\t\tA=%v B=%v\n",a,b)
}
