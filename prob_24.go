package main

import (
	"L1/points"
	"fmt"
	"math"
)

func CalcDistance(p1,p2 points.Point) float64{
	x1,y1:=p1.GetXY()
	x2,y2:=p2.GetXY()
	distance:=math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2))
	return distance
}

func prob24(){
	a:=points.NewPoint(7.07106781186547524400,-7.07106781186547524400)
	b:=points.NewPoint(0,0)
	fmt.Println("Точка 1:",a)
	fmt.Println("Точка 2:",b)
	fmt.Println("Расстояние=",CalcDistance(a,b))
}
