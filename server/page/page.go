package page

type Page struct {
	Content [1024]byte
}

type PageReader interface {
	ApplyDeltaOperation(Operations) bool
}
