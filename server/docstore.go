package main

import (
	"fmt"
	"log"
	"net/http"

	"./page"
	"./quill"
	"github.com/gorilla/websocket"
)

type DeltaOp struct {
	Operation string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func deltaHandler(delta []byte, user page.User) {
	operation := quill.GetDelta(delta)
	p := page.Packet{Delta: *operation, UserData: user}
	// log.Println("operation : ", operation)
	d.ChanDelta <- p
}

func docserve(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// initialize the new users pageview to that of server
	//user := page.User{Socket: *c, PageView: d.DocPage}
	user := page.NewUser(*c, d.DocPage)

	// TODO: add the user to the document users list
	d.AddUser(&user)
	fmt.Println("Added new user")
	fmt.Println(user.ID)

	for {
		// TODO: find out significance of message type return value
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			if err.Error() == "websocket: close 1001 (going away)" {
				fmt.Println("connection terminated")
				d.RemoveUser(&user)
				return
			}
		}

		deltaHandler(message, user)
	}
}

// defining this as global variable for simplicity, need to handle doc initilaization on user request
// probably need an updater routine per user with a lock on the document for mutex
var d page.Doc

func main() {
	log.SetFlags(0)
	d = page.NewDoc()
	go d.DocUpdater() // go routine for document synchronization
	http.HandleFunc("/", docserve)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
