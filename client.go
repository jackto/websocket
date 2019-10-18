package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

var orgin = "http://127.0.0.1:8000"
var url = "ws://127.0.0.1:8000/ws"
var signal = make(chan bool)

func main() {

	for i := 0; i < 10000; i++ {
		go call(string(i))
		time.Sleep(10 * time.Millisecond)
	}
	signal <- true
}

func call(threadNum string) {

	var msg = make([]byte, 100000)
	//data:=map[string]string{"email":"jack.to"+threadNum+"@163.com","username":"user"+threadNum,"message":"hi, this is a test"}

	ws, err := websocket.Dial(url, "", orgin)
	if err != nil {
		fmt.Println("%v", err)
	}
	defer ws.Close()

	for {

		/*
			index := rand.Intn(300)
			if index%300==0{
				message,err:= Marshal(data)
				if err!=nil{
					fmt.Println("%v",err)
					break
				}

				_,err = ws.Write(message)
				if err!=nil{
					fmt.Println("%v", err)
					break
				}
			}
		*/

		m, err := ws.Read(msg)
		if err != nil {
			fmt.Println("%v", err)
			break
		}

		fmt.Printf("Receive %s", string(msg[:m]))
		time.Sleep(1 * time.Second)
	}

}
