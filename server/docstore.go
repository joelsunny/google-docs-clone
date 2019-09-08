package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"./quill"
	"github.com/gorilla/websocket"
)

type Document struct {
	Message              string
	LastModifiedUserName string
	mux                  sync.Mutex
}

type DeltaOp struct {
	Operation string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var document = Document{Message: "", LastModifiedUserName: "Nintappan"}

func deltaHandler(delta []byte) {
	operation := quill.GetDelta(delta)
	log.Println("operation : ", operation)
	document.mux.Lock()
	defer document.mux.Unlock()
	// perform operation
	retain := operation.Retain
	delete := operation.Delete
	insertLength := len(operation.Insert)
	docLength := len(document.Message)
	if document.Message == "" && retain == 0 && delete == 0 {
		// a new message starts
		document.Message = operation.Insert
	} else if retain == docLength && delete == 0 {
		// appending to message
		document.Message = document.Message + operation.Insert
	} else if delete+retain == docLength-1 {
		// deleteing from the end and replacing
		document.Message = document.Message[:retain] + operation.Insert
	} else if retain+delete < docLength && delete > insertLength {
		// replacing and message size reduces
		document.Message = document.Message[:retain] + operation.Insert + document.Message[retain+delete:]
	} else if retain+delete < docLength && delete <= insertLength {
		// replacing and message size reduces
		document.Message = document.Message[:retain] + operation.Insert + document.Message[retain+delete:]
	}
}

func docserve(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		deltaHandler(message)

		err = c.WriteMessage(mt, []byte(document.Message))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", docserve)
	log.Fatal(http.ListenAndServe("192.168.0.103:8080", nil))
}
