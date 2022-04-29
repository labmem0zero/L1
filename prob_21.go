package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

//сервер с биг датой
type server struct {
	bigData []byte
}
//интерфейс с методом получения данных
type OperationalSystem interface {
	ReceiveJSON([]byte)
}

//метод для отправки данных выбранной "OS"
func (s *server)SendData(os OperationalSystem){
	os.ReceiveJSON(s.bigData)
}

//формат структуры данных, получаемых "Windows"
type Windows struct {
	BigData struct{
		Id int	`json:"id"`
		Name string	`json:"name"`
		Time time.Time	`json:"time"`
	}
}

//Метод получения данных "Windows"
func (w *Windows)ReceiveJSON(data []byte){
	fmt.Println("PC получил JSON")
	json.Unmarshal(data, &w.BigData)
}

//формат структуры данных, получаемых "Android"
type Android struct {
	DickData struct{
		Id int	`xml:"id"`
		Name string	`xml:"name"`
		Time time.Time	`xml:"time"`
	}
}

//Метод получения данных "Android"
func (a *Android) ReceiveXML(data []byte){
	fmt.Println("Android получил XML")
	xml.Unmarshal(data,&a.DickData)
}

//Структура для нашего адаптера
type androidPCAdapter struct {
	androidPC *Android
}

//Адаптер: олучаем данные методом "ReceiveJSON", не прописанным для структуры "Android"
//на практике конкретно такой вариант мы вряд ли увидим, но на бумаги вот он, АДАПТЕР
func (apA *androidPCAdapter)ReceiveJSON(data []byte){
	fmt.Println("Адаптер получил JSON, адаптируем для Android")
	var tmp testStruct
	json.Unmarshal(data,&tmp)
	newData,_:=xml.Marshal(tmp)
	apA.androidPC.ReceiveXML(newData)
}


type testStruct struct {
	Id int	`json:"id" xml:"id"`
	Name string	`json:"name" xml:"name"`
	Time time.Time	`json:"time" xml:"time"`
}

func prob21(){
	testData:=testStruct{
		15,
		"Vasya",
		time.Now(),
	}
	fmt.Println("Тестовая структура: ",testData)
	server:=&server{}
	server.bigData,_=json.Marshal(testData)
	w:=&Windows{}
	server.SendData(w)
	fmt.Println("Win data:",w.BigData)
	a:=&Android{}
	aA:=&androidPCAdapter{a}
	server.SendData(aA)
	fmt.Println("Android data:",a.DickData)
}

