package main

import (
	"fmt"
	"time"
)

func MySleep(dur time.Duration){
	timer:=time.NewTimer(dur)
	<-timer.C
	return
}

func prob25(){
	dur:=time.Duration(5*time.Second)
	fmt.Printf("Начальное время:\t\t\t\t\t\t\t\t%v\t\t\tвремя для сна:%v\n",time.Now(),dur)
	MySleep(dur)
	fmt.Printf("MySleep завершил свое действие, текущее время:\t%v\n",time.Now())
}