package main

import (
	"fmt"
	"math/big"
)

func prob22(){
	//создаем два больших числа и работаеим с ними при помощи встроенного пакета math/big
	a:=big.NewFloat(2<<21)
	b:=big.NewFloat(2<<21+5)
	res:=big.NewFloat(0)
	res.Add(a,b)
	fmt.Printf("%v+%v=%.0f\n",a.String(),b.String(),res)
	res.Sub(a,b)
	fmt.Printf("%v-%v=%.0f\n",a.String(),b.String(),res)
	res.Mul(a,b)
	fmt.Printf("%v*%v=%.0f\n",a.String(),b.String(),res)
	res.Quo(a,b)
	fmt.Printf("%v*%v=%.10f\n",a.String(),b.String(),res)
}
