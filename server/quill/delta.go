package quill

import (
	"encoding/json"
	"log"
)

type delta map[string]interface{}

type Operations struct {
	retain float64
	insert string
	delete float64
}

// GetDelta: function to get delta
func GetDelta(data []byte) *Operations {

	var arr []delta
	_ = json.Unmarshal(data, &arr)

	var op Operations

	for _, d := range arr {
		r := d["retain"]
		i := d["insert"]
		d := d["delete"]

		rf, ok := r.(float64)
		if ok {
			op.retain = rf
		}

		iS, ok := i.(string)
		if ok {
			op.insert = iS
		}

		df, ok := d.(float64)
		if ok {
			op.delete = df
		}
	}
	log.Println("delta: ", op)
	return &op
}
