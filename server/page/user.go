package page

import (
	"encoding/json"
	"fmt"
	"log"

	"../quill"
	websocket "github.com/gorilla/websocket"
)

// UserPacket :- protocol unit for user updates
type UserPacket struct {
	Type    string
	Message interface{}
}

// User :- user struct definition
type User struct {
	ID            int // use socket as id?
	Socket        websocket.Conn
	PageView      Page // use mutex, no need to use mutex if only one goroutine is updating this
	ChanDelta     chan quill.Operations
	InsertTracker Page
	DeleteTracker Page
}

var uid int

// NewUser :-
func NewUser(socket websocket.Conn, doc Page) User {
	u := User{ID: uid, ChanDelta: make(chan quill.Operations, 10), Socket: socket, PageView: doc}
	uid++
	go u.UserUpdate()
	return u
}

// ApplyDeltaOp :- update PageView of user on receiving a new delta
func (u *User) ApplyDeltaOp(delta quill.Operations) {
	return
}

// UserUpdate :- send delta from collaborators to user
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

// AckHandler :- handle user acknowledement of deltas
func (u *User) AckHandler() {

}
