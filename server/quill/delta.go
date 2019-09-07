package quill

import (
	"encoding/json"
)

type delta map[string]interface{}

type Operations struct {
	Retain float64
	Insert string
	Delete float64
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
			op.Retain = rf
		}

		iS, ok := i.(string)
		if ok {
			op.Insert = iS
		}

		df, ok := d.(float64)
		if ok {
			op.Delete = df
		}
	}
	return &op
}
