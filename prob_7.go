package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func prob7(){
	var prob7map sync.Map
	var wg sync.WaitGroup
	for i:=1;i<6;i++{
		wg.Add(1)
		go func(wrkrID int){
			for i:=1;i<1000;i++{
				prob7map.LoadOrStore(rand.Intn(99999),rand.Intn(99999))
			}
			fmt.Printf("Worker #%v:1000 записей в катру сделаны\n",wrkrID)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
