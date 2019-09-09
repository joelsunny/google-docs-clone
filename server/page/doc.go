package page

import "../quill"

// Page: Page of a Document
type Doc struct {
	Index [100]int64 // 64 bit address space assumption
}

func (d *Doc) ApplyDeltaOperation(delta quill.Operations) {
	retain := delta.Retain
	delete := delta.Delete
	insertLength := len(delta.Insert)
	insert := []byte(delta.Insert)

	startPage := retain % 1024
}
