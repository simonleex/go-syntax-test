package gotest

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/myteksi/go/messaging/grab-messaging/common/errors"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestClosedChannel(t *testing.T) {
	ch := make(chan int, 1)
	//ch <- 1
	close(ch)
	v, ok := <-ch
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)

	ch1 := make(chan int, 1)
	ch1 <- 1
	v1, ok := <-ch1
	assert.Equal(t, 1, v1)
	assert.Equal(t, true, ok)

	ch2 := make(chan int, 1)
	ch2 <- 1
	close(ch2)
	v2, ok := <-ch2
	assert.Equal(t, 1, v2)
	assert.Equal(t, true, ok)
}

func TestVariadic(t *testing.T) {
	f := func(a int, b ...int) int {
		return len(b)
	}
	assert.Equal(t, 0, f(1))
	assert.Equal(t, 2, f(1, 2, 3))
	var b []int
	assert.Nil(t, b)
	b = make([]int, 0)
	assert.Equal(t, 0, f(1, b...))
	b = append(b, 1, 2, 3)
	assert.Equal(t, 3, f(1, b...))
}

func TestDeferInInnerFun(t *testing.T) {
	var lock sync.Mutex
	te := func() {
		lock.Lock()
		defer lock.Unlock()
	}
	te()
	lock.Lock()
	lock.Unlock()
}

func TestSliceAppendBasic(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3}
	s3 := append(s1[:1], s2...)
	assert.Equal(t, s3, s1)
	assert.Equal(t, []int{1, 1, 2, 3}, s1)
	s4 := append(s1[:2], s2...)
	assert.NotEqual(t, s4, s1)
	assert.Equal(t, s1, s3)
	assert.Equal(t, []int{1, 1, 2, 3}, s1)
	assert.Equal(t, []int{1, 1, 1, 2, 3}, s4)
}

func TestSliceAppend(t *testing.T) {
	f1 := func() []int {
		s1 := []int{1, 2, 3, 4}
		s2 := []int{-1, -2, -3}
		return append(append(s1[:1], s2...), s1[1:]...)
	}

	f2 := func() []int {
		s1 := []int{1, 2, 3, 4}
		s2 := []int{-1, -2, -3, -4}
		return append(append(s1[:1], s2...), s1[1:]...)
	}

	assert.Equal(t, []int{1, -1, -2, -3, -1, -2, -3}, f1())

	assert.Equal(t, []int{1, -1, -2, -3, -4, 2, 3, 4}, f2())
}

func TestNew(t *testing.T) {
	type A struct {
		a interface{}
	}
	C := new(A)
	assert.Equal(t, &A{nil}, C)

}

func TestDefer(t *testing.T) {
	c := func() (i int) {
		i = 0
		defer func() {
			i++
		}()
		return
	}

	assert.Equal(t, 1, c())

	c1 := func() int {
		i := 0
		defer func() {
			i++
		}()
		return i
	}

	assert.Equal(t, 0, c1())

	c2 := func() int {
		i := 0
		defer func(i int) {
			i++
		}(i)
		return i
	}

	assert.Equal(t, 0, c2())

	c3 := func() (i int) {
		i = 0
		defer func(i int) {
			i++
		}(i)
		return
	}

	assert.Equal(t, 0, c3())
}

func TestSliceCap(t *testing.T) {
	sl := make([]int, 5, 6)
	assert.Equal(t, 5, len(sl))
	assert.Equal(t, 6, cap(sl))
	sl1 := []int{1}
	assert.Equal(t, 1, len(sl1))
	assert.Equal(t, 1, len(sl1))
}

func TestBreakLabel(t *testing.T) {

	v := 1
Label:
	for {
		break Label
	}
	assert.Equal(t, 1, v)

}

func TestForLoop(t *testing.T) {
	var i, j int
	for i = 0; i < 10 && j == 0; i++ {
		if i == 9 {
			j = 1
		}
	}
	if i < 10 && j == 0 {
		println("ddd")
	}
	println(i, j)

}

type User struct {
	Name string
	id   int
}

func TestGob(t *testing.T) {
	user := User{"test", 1}
	buf := bytes.NewBuffer([]byte{})
	Enc := gob.NewEncoder(buf)
	err := Enc.Encode(user)
	assert.Nil(t, err)

	fmt.Println(buf)

	newuser := User{"new", 2}
	Dec := gob.NewDecoder(buf)
	Dec.Decode(&newuser)
	assert.Equal(t, User{"test", 2}, newuser)

}

func TestString(t *testing.T) {
	var s string
	assert.Equal(t, "", s)
}

func TestIntLimit(t *testing.T) {
	println(-(math.MaxInt32 + 1))
	println(math.MinInt32)

}

// os.exit(1) actually can't stop by recover
func TestOsExitRecover(t *testing.T) {
	f := func() {
		defer func() {
			if err := recover(); err != nil {
				println(err)
			}
		}()
		os.Exit(1)
	}
	f()
	println("still working")
}

var token = `mutual?agent=grab&token=EAAHVAvYTQm8BAAqZBrdWGNjxU9gwKTWaZBVpZAAK8S3WmAMv3jJg
jjdZAev71yvPh9fwpbYZAVxE1wxErj6bxJwEaeV80n6xa1oJhrYb8VQtaUcgRZCyVBY3X3CCZAA
IGkvmpZCvXIW8skuRfG7eSE3t74iZBSJdq7yfT0ZAJoolqLTIqcjXsgBwZAn3zhOO6xxXZCDqzn
5VtNi4bzMZAHVOZBlljI1RWHly6ZAmyPSTAoOkP1FrqculFhTxWlK
OyNCvyVgRuijD6WJZC5FEoKFLwZDtual?agent=grab&token=EAAHVAvYTQm8BAAqZBrdWGNjxU9gwKTWaZBVpZAAK8S3WmAMv3jJg
jjdZAev71yvPh9fwpbYZAVxE1wxErj6bxJwEaeV80n6xa1oJhrYb8VQtaUcgRZCyVBY3X3CCZAA
IGkvmpZCvXIW8skuRfG7eSE3t74iZBSJdq7yfT0ZAJoolqLTIqcjXsgBwZAn3zhOO6xxXZCDqzn
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
1CcBIzWBW4uprQtCfZAmRWPLO3c8vxkggr7rPWg2ZA4EbwRVClrd45jAZBFhcqzeFZC3omEhuAT
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
OyNCvyVgRuijD6WJZC5FEoKFLwZD`

type trans func()

func BenchmarkSha256(b *testing.B) {
	fmt.Printf("%v\n", len(token))
}

func TestNotNil(t *testing.T) {
	var chana chan int
	assert.NotNil(t, chana)
}

var msg = `mutual?agent=grab&token=EAAHVAvYTQm8BAAqZBrdWGNjxU9gwKTWaZBVpZAAK8S3WmAMv3jJg
jjdZAev71yvPh9fwpbYZAVxE1wxErj6bxJwEaeV80n6xa1oJhrYb8VQtaUcgRZCyVBY3X3CCZAA
IGkvmpZCvXIW8skuRfG7eSE3t74iZBSJdq7yfT0ZAJoolqLTIqcjXsgBwZAn3zhOO6xxXZCDqzn
5VtNi4bzMZAHVOZBlljI1RWHly6ZAmyPSTAoOkP1FrqculFhTxWlK
OyNCvyVgRuijD6WJZC5FEoKFLwZDtual?agent=grab&token=EAAHVAvYTQm8BAAqZBrdWGNjxU9gwKTWaZBVpZAAK8S3WmAMv3jJg
jjdZAev71yvPh9fwpbYZAVxE1wxErj6bxJwEaeV80n6xa1oJhrYb8VQtaUcgRZCyVBY3X3CCZAA
IGkvmpZCvXIW8skuRfG7eSE3t74iZBSJdq7yfT0ZAJoolqLTIqcjXsgBwZAn3zhOO6xxXZCDqzn
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
1CcBIzWBW4uprQtCfZAmRWPLO3c8vxkggr7rPWg2ZA4EbwRVClrd45jAZBFhcqzeFZC3omEhuAT
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
kr7ThyJ3CJEv5Ab8VW3kX23p9jNsw5MUuotknTYbZA5hwJzfQHcw6Hr0v0fjn6Ra1VZABReXJaI
OyNCvyVgRuijD6WJZC5FEoKFLwZD`

var entity = msg + "-" + "123456789"

func BenchmarkT1(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		if strings.HasPrefix(entity, msg) {
			msgs := strings.Split(entity, "-")
			msgNum := len(msgs)
			msgID := msgs[msgNum-1]
			if len(entity) == len(msg)+1+len(msgID) {
				result = msgID
			}
		}
	}
	_ = result
}

func BenchmarkT2(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		if strings.HasPrefix(entity, msg) {
			msgs := strings.Split(entity, "-")
			msgNum := len(msgs)
			msgID := msgs[msgNum-1]
			result = msgID
		}
	}
	_ = result
}

func TestParseID(t *testing.T) {

	msgID := ""
	for i := len(entity) - 1; i >= 0; i-- {
		if entity[i:i+1] == "-" {
			msgID = entity[i+1:]
			break
		}
	}
	assert.NotEqual(t, "", msgID)
	assert.Equal(t, "123456789", msgID)
}

type server interface {
	pull()
	push()
}

type ATEST struct {
	a int
}

type Ser struct {
}

func (s *Ser) pull() {
	println("Ser pull")
}

func (s *Ser) push() {
	println("Ser push")
}

type Ser1 struct {
	Ser
}

func (s1 *Ser1) push() {
	println("Ser1 push")
}

func TestReflect(t *testing.T) {
	a := ATEST{1}
	var b interface{}
	b = a
	fmt.Printf("%v\n", reflect.TypeOf(b).Name())
	fmt.Printf("%d\n", time.Millisecond*10)

	var s server
	s1 := &Ser1{}
	s = s1
	s.pull()
	s.push()

}

func TestFanIn(t *testing.T) {
	buffer := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func(i int, b chan int) {
			for k := range b {
				println(i, k)
			}
		}(i, buffer)
	}

	for i := 0; i < 100; i++ {
		buffer <- i
	}
	close(buffer)
}

func TestWgWaitAndTimeout(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		time.Sleep(3 * time.Second)
		wg.Done()
	}()

	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		wg.Wait()
		println("wait")
	}()
	select {
	case <-doneChan:
		println("done")
	case <-time.After(5 * time.Second):
		println("timeout")
	}
}

var commonInt = []int{4, 5, 6}

func fun1(a, b int, c ...int) {
	println(a, b, c)
}

func TestVardic(t *testing.T) {
	fun1(1, 2, commonInt...)
}

func TestPrint(t *testing.T) {
	a := 0x10000000
	println(a)

}

type e struct {
}

func (*e) Error() string { return "boom" }
func f() *e              { return nil }

// PushNotificationRequest defines push notification request
type PushNotificationRequest struct {
	Sender        string                 `json:"sender"`
	Event         string                 `json:"event"`
	Message       string                 `json:"title"`
	MessageType   interface{}            `json:"messageType"`
	Body          string                 `json:"body"`
	Recipients    []int64                `json:"recipients"`
	RecipientType string                 `json:"recipientType"`
	PushData      map[string]interface{} `json:"pushData"`
}

func TestPrintln(t *testing.T) {
	err := errors.New("ys")
	err = f()
	println(err == nil)

	push := PushNotificationRequest{
		Sender:      "sender",
		Event:       "event",
		MessageType: "type200",
		Recipients:  []int64{1, 2, 3},
		PushData:    map[string]interface{}{"yes": "yes", "no": 1},
	}

	fmt.Printf("%v\n", push)

	c := make(chan int, 2)
	c <- 1
	c <- 2
	<-c
	close(c)
	fmt.Println(<-c)

}

func TestTimeWait(t *testing.T) {
	for i := 0; i <= 100; i++ {
		fmt.Printf("i:%d\n", i)
		if i%10 != 0 {
			continue
		}

		<-time.After(time.Second * 3)
	}
}

const (
	factor = 5
)

func TestMaxInt32Mod(t *testing.T) {
	var count int32 = 0
	count = math.MaxInt32 - 50
	for i := 0; i <= 100; i++ {
		count += 1
		if count%factor == 0 {
			fmt.Printf("i:%d count:%d\n", i, count)
		}
	}
	var f float32 = 16777216 // 1 << 24
	fmt.Printf("%v %v %v\n", f, f+1, f == f+1)
}

type TestTagStruct struct {
	Test int `json:"jtag" fuck:"nofuck"`
}

func TestTag(t *testing.T) {
	a := TestTagStruct{}
	tyo := reflect.TypeOf(a)
	fi := tyo.Field(0)
	loup, err := fi.Tag.Lookup("json")
	fmt.Printf("%v %v %v %v\n", loup, err, fi.Tag.Get("json"), fi.Tag.Get("fuck"))

}

var testSlice = []int{5, 6, 7, 8}

func TestRangeSlice(t *testing.T) {
	for i, j := range testSlice {
		fmt.Printf("%v %v", i, j)
	}

}

func TestMapAndConvert(t *testing.T) {
	m := map[string]interface{}{
		"a": true,
		"b": 1,
	}

	res, ok := m["a"].(bool)
	fmt.Printf("%v %v", res, ok)

}

func TestOperatorMul(t *testing.T) {
	a := 2
	b := 3
	c := 8
	c *= b - a
	fmt.Printf("%v", c)

}

func TestCutSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	subs1 := s[2:4]
	fmt.Printf("%v %v", s, subs1)
	s[3] = 10
	assert.Equal(t, 10, subs1[1])
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s1 := make([]int, 3, 5)
	copy(s1, s[2:5])
	s1[1] = 10
	s2 := s1[1:3]
	s2[1] = 1111
	s3 := append(s1, 6)
	s3[1] = 7
	fmt.Printf("%v %v %v %v", s, s1, s2, s3)
}

func TestSliceExpand(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[2:5]
	s = append(s, 8)
	s1[1] = 111
	fmt.Printf("%v %v %v", s, s1, s)
}

func TestHappenedBefore(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 1005; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			fmt.Printf("%v\n", count)
			count++
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%v", count)
}

func TestMapRetrieve(t *testing.T) {
	m := make(map[int]map[int]int)
	m[1] = make(map[int]int)
	m[1][1] = 1
	m[1][0] = 1
	x, y := m[1][2]
	fmt.Printf("%v %v\n", x, y)
	x, y = m[2][1]
	fmt.Printf("%v %v\n", x, y)
	if m[3] == nil {
		fmt.Printf("nil")
	}

}

type testtc struct {
	a, b int
}

func TestTypeConversion(t *testing.T) {
	a := 1
	b := float32(a)
	fmt.Printf("%v %v ", a, b)
}

func (t *testtc) add1() {

}

func (t *testtc) add2() {

}

type adder interface {
	add1()
}

type adderplus interface {
	add1()
	add2()
}

func TestTypeAssertion(t *testing.T) {
	var x adderplus
	x = &testtc{}
	x.add1()
	y := x.(adder)
	y.add1()
	fmt.Printf("%v", string(6767))

}

func TestAssignAbility(t *testing.T) {
	var x chan int
	var y <-chan int
	y = x
	_ = y
}

func TestConstantAssign(t *testing.T) {
	var a float32
	var b float64
	var c int32
	var d int64
	const untyped = 1
	a = untyped
	b = untyped
	c = untyped
	d = untyped
	_, _, _, _ = a, b, c, d
}

type testano struct {
	a int
}

type record struct {
	testtc
	testano
}

func (t *testano) add1() {

}

func TestAnonymousFieldFunc(t *testing.T) {
	a := &testtc{}
	a.add1()
	a.add2()
	var b adder
	b = a
	b.add1()
	v := reflect.ValueOf(b)
	fmt.Printf("%v", v)
}

func TestBreakLabel1(t *testing.T) {
	for i := 0; i < 10; i++ {
	OuterLoop:
		for j := 0; j < 10; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			break OuterLoop
		}
	}
	gotest := 2
	fmt.Printf("%v", gotest)
}

var (
	a int = b + 1
	b int = 1
)

func TestInitVar(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
}

type Employee struct {
	id     int
	salary int
}

func EmployeeByID(id int) *Employee { return &Employee{10, 10} }

func TestPointArray(t *testing.T) {
	x := []*record{}
	fmt.Printf("%v\n", x)
}

func TestReturnValue(t *testing.T) {
	x := fmt.Errorf("asd")
	fmt.Printf("%v", x)
}

func p(i int) {
	print(i + 10)
}

func myrecover() {
	print("4")
	if r := recover(); r != nil {
		fmt.Printf("rec %v", r)
	}
	print("5")
}

func TestPanicFlow(t *testing.T) {

	defer p(1)
	defer myrecover()
	print("1")
	panic("panic")
	print("2")
	defer p(2)
	print("3")
	defer p(3)
}

func return2() (int, int) {
	return 1, 2
}

func return3() (int, int) {
	return return2()
}

func TestGC(t *testing.T) {
	a, b := return3()

	fmt.Printf("%d %p %v", a, b, &b)
}
