package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joelsunny/docstore/server/page"
	"github.com/joelsunny/docstore/server/quill"
)

type DeltaOp struct {
	Operation string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var document page.Page

func deltaHandler(delta []byte) {
	operation := quill.GetDelta(delta)
	log.Println("operation : ", operation)
	document.ApplyDeltaOperation(*operation)
}

// will be cleaned up
func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
				if ipnet.IP.String()[0:3] == "192" {
					return (ipnet.IP.String())
				}
			}
		}
	}
	return ("")
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

		err = c.WriteMessage(mt, document.GetContentAsByte())
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.SetFlags(0)
	http.HandleFunc("/", docserve)
	log.Fatal(http.ListenAndServe(getIP()+":8080", nil))
}
