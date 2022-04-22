package main

import (
	"fmt"
	"reflect"
)

func CheckType(in interface{})string{
	return fmt.Sprintf(reflect.TypeOf(in).String())
}

func prob14(){
	//Вообще можно не объявлять интерфейсы, но решил явно указать
	var a interface{}
	var b interface{}
	var c interface{}
	a=1
	b=2.1
	c="three"
	fmt.Printf("A=%v\t\t\tB=%v\t\t\tC=%v\n",a,b,c)
	fmt.Printf("A=%v\t\tB=%v\t\tC=%v\t-\tопределение через рефлект\n",CheckType(a),CheckType(b),CheckType(c))
	fmt.Printf("A=%T\t\tB=%T\t\tC=%T\t-\tопределение с помощью форматированных строк\n",a,b,c)
}