package main

import "fmt"

func SliceToSet(input []string)[]string{
	//использую медленный способ(с использованием карты)
	//можно сделать быстрый способ, но для этого нужно сначала упорядочивать
	//входной срез по алфавиту, что бы можно было итерировать по 2 переменным и пропускать дубликаты.
	//Пример быстрого создания множества из упорядоченного среза:
	/*
	res:=[]string{a[0]}
		i:=0
		j:=1
		for i<len(a){
			if a[j]!=a[i]{
				res=append(res,a[j])
				i=j
			}
			j++
			if j>=len(a){
				break
			}
		}
	 */
	temp:=make(map[string]bool)
	for _,i:=range input{
		temp[i]=true
	}
	var res []string
	for k,_:=range temp{
		res=append(res,k)
	}
	return res
}

func prob12(){
	a:=[]string{"cat","dog","cow","cat"}
	fmt.Println("Входные данные:        ",a)
	fmt.Println("Результирующие данные: ",SliceToSet(a))
}
