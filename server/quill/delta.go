package quill

import (
	"encoding/json"
	"fmt"
)

type delta map[string]interface{}

type Operations struct {
	Retain int    `json:"retain"`
	Insert string `json:"insert"`
	Delete int    `json:"delete"`
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
			op.Retain = int(rf)
		}

		iS, ok := i.(string)
		if ok {
			op.Insert = iS
		}

		df, ok := d.(float64)
		if ok {
			op.Delete = int(df)
		}
	}
	return &op
}

func ConvertToQuillFormat(delta Operations) interface{} {
	var ops [3]map[string]interface{}
	var d map[string]interface{}

	ops[0] = map[string]interface{}{"retain": delta.Retain}
	ops[1] = map[string]interface{}{"insert": delta.Insert}
	ops[2] = map[string]interface{}{"delete": delta.Delete}

	d = map[string]interface{}{"ops": ops}
	fmt.Println("delta: ")
	fmt.Println(d)
	return &d
}
