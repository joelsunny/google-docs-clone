package main

import (
	"./page"
	"./quill"
)

func testIndividual() {
	var p page.Page
	p.InsertString(0, []byte("hello"))
	p.Print()
	p.InsertString(5, []byte("world"))
	p.Print()
	p.DeleteString(5, 2)
	p.Print()
	p.InsertString(5, []byte("wo"))
	p.Print()
}

func main() {
	var p page.Page
	d := quill.Operations{Retain: 0, Delete: 0, Insert: "hello"}
	p.ApplyDeltaOperation(d)
	p.Print()

	d = quill.Operations{Retain: 5, Delete: 0, Insert: "world"}
	p.ApplyDeltaOperation(d)
	p.Print()

	d = quill.Operations{Retain: 4, Delete: 3, Insert: "wo"}
	p.ApplyDeltaOperation(d)
	p.Print()
}
