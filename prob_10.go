package main

import "fmt"

func gradationBy10(temp float64)string{
	x:=int(temp/10)
	if temp<0{
		x--
	}
	return fmt.Sprintf("от %.1f до %.1f",float64(x*10),float64((x+1)*10))
}
func gradateTemps(temps []float64) map[string][]float64{
	tmpMap:=make(map[string][]float64)
	for _,temp:=range temps{
		gradation:=gradationBy10(temp)
		tmpMap[gradation]=append(tmpMap[gradation],temp)
	}
	return tmpMap
}
func prob10(){
	input:=[]float64{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
	result:=gradateTemps(input)
	for k,v:=range result{
		fmt.Println(k,":",v)
	}
}