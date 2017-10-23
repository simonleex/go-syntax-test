package main

import (
	"fmt"
	"github.com/myteksi/go/commons/util/log/logging"
	"github.com/myteksi/go/staples/gredis3"
	"github.com/myteksi/go/staples/gredis3/conman"
	"github.com/myteksi/go/staples/gredis3/gredisapi"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"time"
)

type Request struct {
	Gender string `json:"gender"`
	ID     int64  `json:"id"`
}

func pint64(a int64) {
	println(a)
}

func init() {
	fmt.Printf("init1")

}

func init() {

	fmt.Printf("init2")

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
	/*
		var m map[int]int = make(map[int]int)
		ch1 := make(chan int)
		ch2 := make(chan int)
		m[1] = 1

		//go read(m, ch1)
		//go read(m, ch2)
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

		//testMgo()
	*/
	//redisTest()
	defer myrecover()
	go letmepanic()
	//f()
	fmt.Println("Returned normally from f.")
	time.Sleep(15)
}

type Book struct {
	ID   int32
	Name string
}

func testMgo() {
	DBSess, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	DBSess.SetMode(mgo.Monotonic, true)
	DB := DBSess.DB("test").C("book")
	err = DB.Insert(&Book{1, "abc"}, &Book{2, "robot"})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	DBSess.Close()
}

func redisTest() {
	//redis client init
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	conMan := conman.NewSingleHost(context.Background(),
		conman.Host("localhost"),
		conman.Port(6379),
		conman.MaxActive(5),
		conman.MaxIdle(5))

	var err error
	redisClient, err := gredis3.NewClient(ctx, conMan, gredis3.Logger(logging.GetDefaultLogger()), gredis3.Name("borders"))
	if err != nil {
		logging.Error("gredis3", "init gredis3 failed, err% %v", err)
		return
	}

	/*
		fmt.Printf("\n")
		for i := 0 ; i < 10; i ++ {
			res, err := redisClient.Do(context.Background(), "SET", "test1", "whatever","EX",int64(2),"NX")
			if err != nil {
				fmt.Printf("resssserr:%v",err)
				continue
			}

			str, ok := res.(string)
			fmt.Printf("resss:%v %v %v\n", res,str,ok)
			time.Sleep(1*time.Second)
			if ok && str == "OK" {
				fmt.Printf("ok")
			}
		}
	*/
	res, err := gredisapi.Strings(redisClient.Do(context.Background(), "ZRANGE", "lddl", int64(-1), int64(-1), "WITHSCORES"))
	fmt.Printf("\nres%v %v\n", res, err)

}

func p(i int) {
	fmt.Println(i + 10)
}

func myrecover() {
	fmt.Println(4)
	if r := recover(); r != nil {
		fmt.Printf("rec %v", r)
	}
	fmt.Println(5)
}

func letmepanic() {
	//defer p(1)
	defer p(2)
	//defer myrecover()
	defer p(3)

	fmt.Println("1")
	panic("panic")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
