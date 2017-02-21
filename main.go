package main


func consumer(ch chan int, tag int) {
	for {
		println(<- ch, tag)
	}
}

func main() {
	ch := make(chan int)

	go consumer(ch, 1)
	go consumer(ch, 2)
	go consumer(ch, 3)

	for i := 0; i < 1000; i ++ {
		ch <- i
	}

}
