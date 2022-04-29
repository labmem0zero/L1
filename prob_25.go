package main

import (
	"fmt"
	"time"
)

func MySleep(dur time.Duration){
	//таймер. Ждет окончание отсчета, завершает функцию.
	//Так, как ожидает канал timer.C, блокирует последующее выполнение функции
	//Ничем не отличается от обычного слипа
	//прям с нуля написать соип не вижу воможности, потому что не умею
	//погдлядеть код без шансов, потому что основные функции времени - built-in
	//Мы можем, конечно, чподглядеть пакет time и пакет runtime/time, который он использует
	//но упремся во все те же built-in функции
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