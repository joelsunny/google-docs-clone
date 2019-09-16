package page

import (
	"fmt"
	"log"

	"../quill"
)

const ULIMIT = 10

// Packet :- struct representing a single unit of update
type Packet struct {
	Delta    quill.Operations
	UserData User // only need the socket? any disadvantage in passing around structs?
}

// Doc :- struct representing a single document
type Doc struct {
	DocPage     Page // no need to use mutex if only one goroutine updates this
	UserList    [ULIMIT]*User
	activeUsers int
	ChanDelta   chan Packet
}

// NewDoc :- call this function to initialize a new Doc object
func NewDoc() Doc {
	d := Doc{ChanDelta: make(chan Packet, 50)}
	return d
}

// AddUser :- add a new user to the active user list of document
func (d *Doc) AddUser(u *User) {
	d.UserList[d.activeUsers] = u
	d.activeUsers++
}

// RemoveUser :- remove user from the active user list of document
func (d *Doc) RemoveUser(u *User) {
	for i := 0; i < ULIMIT; i++ {
		if d.UserList[i].ID == u.ID {
			d.UserList[i] = nil
			for j := i; j < d.activeUsers-1; j++ {
				d.UserList[j] = d.UserList[j+1]
			}
			d.activeUsers--
			fmt.Println("removed user")
			fmt.Println(u.ID)
			return
		}
	}
}

// DocUpdater :- goroutine for hanlding document and user view updates
func (d *Doc) DocUpdater() {

	// listen on channel for updates
	for i := range d.ChanDelta {
		// apply update to document
		d.DocPage.ApplyDeltaOperation(i.Delta)

		for j := 0; j < d.activeUsers; j++ {
			err := d.UserList[j].Socket.WriteMessage(1, d.DocPage.GetContentAsByte())
			if err != nil {
				log.Println("write:", err)
				break // remove break?
			}

			// send delta via channel for remaining users
			if i.UserData.ID != d.UserList[j].ID {
				d.UserList[j].ChanDelta <- i.Delta
			}
		}

	}

	return
}
