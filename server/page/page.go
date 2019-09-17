package page

import (
	"fmt"
	"log"
	"strconv"

	"../quill"
)

const PAGELEN = 100

// Page : Page of a Document
type Page struct {
	Content [PAGELEN]rune
	eoc     int
}

func (p *Page) ApplyDeltaOperation(delta quill.Operations) {
	retain := delta.Retain
	delete := delta.Delete
	insert := []rune(delta.Insert)

	if p.eoc+len(insert)-delete >= 1024 {
		// panic
		return
	}
	p.DeleteString(retain, delete)
	p.InsertString(retain, insert)

	p.Print()
}

func (p *Page) FreeSpace() int {
	return (PAGELEN - p.eoc)
}

func (p *Page) rightShift(index int, shiftValue int) {
	if p.eoc <= index || shiftValue == 0 {
		//panic("error in rightShift, c1")
		return
	}
	if p.eoc+shiftValue > PAGELEN-1 {
		panic("page overflow, rightShift")
	}

	for i := p.eoc; i >= index; i-- {
		p.Content[i+shiftValue] = p.Content[i]
	}
}

func (p *Page) leftShift(index int, shiftValue int) {
	if shiftValue == 0 {
		//panic("error in leftShift, c1")
		return
	}
	for i := index; i <= p.eoc; i++ {
		p.Content[i-shiftValue] = p.Content[i]
	}
}

func (p *Page) DeleteString(retain int, deleteLength int) {
	p.leftShift(retain+deleteLength, deleteLength)
	p.eoc -= deleteLength
}

func (p *Page) InsertString(retain int, str []rune) {
	if p.eoc+len(str) >= PAGELEN {
		log.Println("ERROR: error while insert")
		return
	}
	p.rightShift(retain, len(str))
	for i := 0; i < len(str); i++ {
		p.Content[retain+i] = str[i]
	}

	p.eoc = p.eoc + len(str)
}

func (p *Page) Print() {
	fmt.Println("Content: " + string(p.Content[0:p.eoc+1]))
	fmt.Println("eoc    : " + strconv.Itoa(p.eoc))
}

func (p *Page) GetContentAsByte() []byte {
	return []byte(string(p.Content[:p.eoc]))
}
