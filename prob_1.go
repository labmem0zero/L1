package main

import "fmt"

type Human struct {
	gender string
	age int
}

func (h *Human) SetParams(gender string, age int){
	h.gender=gender
	h.age=age
}

func (h Human) ReturnParams() string {
	return fmt.Sprintf("пол: %s\t|\tвозраст:%v", h.gender,h.age)
}

type Action struct {
	//при указании структуры A как часть структуры B, мы получваем доступ к полям и методам структуры A из структуры B
	Human
}

func prob1(){
	var a Action
	a.SetParams("муж.", 28)
	fmt.Println(a.ReturnParams())
}
