package page

import (
	"encoding/json"
	"fmt"
	"log"

	"../quill"
	websocket "github.com/gorilla/websocket"
)

type UserPacket struct {
	Type    string
	Message interface{}
}

// User :- user struct definition
type User struct {
	ID        int
	Socket    websocket.Conn
	PageView  Page // use mutex, no need to use mutex if only one goroutine is updating this
	ChanDelta chan quill.Operations
}

var uid int

func NewUser(socket websocket.Conn, doc Page) User {
	u := User{ID: uid, ChanDelta: make(chan quill.Operations, 10), Socket: socket, PageView: doc}
	uid++
	go u.UserUpdate()
	return u
}

// ApplyDeltaOp :- update PageView of user on receiving a new delta
// :param ownUpdate :- boolean indicating whether the update is done by the user itself
func (u *User) ApplyDeltaOp(delta quill.Operations, ownUpdate bool) {
	return
}

func (u *User) UserUpdate() {

	m := UserPacket{Type: "delta"}
	for d := range u.ChanDelta {
		fmt.Print("sending delta to user ")
		fmt.Println(u.ID)
		m.Message = d
		b, _ := json.Marshal(m)
		err := u.Socket.WriteMessage(1, b)
		if err != nil {
			log.Println("write:", err)
			break // remove break?
		}
	}
}
