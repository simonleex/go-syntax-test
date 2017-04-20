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

const (
	consta = iota
	constb = iota
	constc = iota
	constd
	conste
)

func read(m map[int]int, ch chan int) {
	k := 0
	for i := 1; i < 100000; i++ {
		k += m[1]
	}
	ch <- k
}

func main() {
	var m map[int]int = make(map[int]int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	m[1] = 1

	go read(m, ch1)
	go read(m, ch2)
	println(<-ch1, <-ch2)
	println(consta, constb, constc, constd, conste)
	reqParam := map[string]interface{}{
		"GnDer": "123",
		"ID":    123,
	}
	data, _ := json.Marshal(reqParam)
	fmt.Printf("%v\n", data)

	req := Request{}

	json.Unmarshal(data, &req)

	fmt.Printf("%v\n", req)

	var i, j int
	for i = 0; i < 5 && j != 1; i++ {
		if i == 2 {
			j = 1
		}
	}
	p := []int{0, 1, 2, 3, 4, 5}
	d := p[3:]
	fmt.Printf(",%v", d)
	println(i)

	str := fmt.Sprintf("%v %s", 123, "321")
	println(str)

}
