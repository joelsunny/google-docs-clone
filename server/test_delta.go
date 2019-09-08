package main

import "./quill"

func test() {
	jsondata := `[{"retain" : 1}, {"insert" : "r"}]`
	quill.GetDelta([]byte(jsondata))
	jsondata = `[{"retain" : 1}, {"delete" : 2}]`
	quill.GetDelta([]byte(jsondata))
}
