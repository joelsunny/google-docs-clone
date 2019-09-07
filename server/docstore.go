package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"./quill"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

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

func deltaHandler(delta []byte, doc *Document) {
	operation := quill.GetDelta(delta)
	log.Println("operation : ", operation)
	doc.mux.Lock()
	defer doc.mux.Unlock()
	doc.Message = doc.Message + operation.Insert
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	document := Document{Message: "This is a Sentence.", LastModifiedUserName: "Nintappan"}
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

		deltaHandler(message, &document)

		err = c.WriteMessage(mt, []byte(document.Message))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe("192.168.0.103:8080", nil))
}
