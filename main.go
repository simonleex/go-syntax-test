package main

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Gender string `json:"gender"`
	ID     int64  `json:"id"`
}

func pint64(a int64) {
	println(a)
}

const c = 123

func main() {

	reqParam := map[string]interface{}{
		"GnDer": "123",
		"ID":    123,
	}
	data, _ := json.Marshal(reqParam)
	fmt.Printf("%v\n", data)

	req := Request{}

	json.Unmarshal(data, &req)

	fmt.Printf("%v\n", req)

	str := fmt.Sprintf("%v %s", 123, "321")
	println(str)

	var a int64
	a = 1012412471982738920012
	var b int
	b := int(a)

	pint64(c)
	println(b)

}
