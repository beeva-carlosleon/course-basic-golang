package main

import (
	"fmt"
	"math/rand"
	"time"
)

func count() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Count finished")
}
func countBlocked(c chan string) {
	count()
	c <- "finished"
}

func producer(ch chan bool) {
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Duration(10 * rand.Float64() * float64(time.Microsecond)))
		ch <- true
	}
	close(ch)
}

func consumer(ch chan bool) {
	for _ = range ch {
		time.Sleep(time.Duration(10 * rand.Float64() * float64(time.Microsecond)))
	}
}

func unbuffered() {
	ch := make(chan bool)
	go producer(ch)
	consumer(ch)
}

func buffered() {
	ch := make(chan bool, 1)
	go producer(ch)
	consumer(ch)
}
func example1() {
	go count()
	var input string
	fmt.Println("Count in process. Press a key for continue...")
	fmt.Scanln(&input)
}
func example2() {
	//With blocking channel
	channel := make(chan string)
	go countBlocked(channel)
	fmt.Printf("Count '%s'\n", <-channel)
}
func example3() {
	//select example
	c1 := make(chan string)
	c2 := make(chan string)
	timeout := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second * 4):
				timeout <- "timeout"
				//default:
				//	time.Sleep(time.Second)
				//	fmt.Println("nothing ready")
			}
		}
	}()
	fmt.Printf("%s received\n", <-timeout)
}
func example4() {
	//buffered channels example
	startTime := time.Now()
	unbuffered()
	fmt.Printf("unbuffered: %v\n", time.Since(startTime))
	startTime = time.Now()
	buffered()
	fmt.Printf("buffered: %v\n", time.Since(startTime))
}
func main() {
	example1()
	example2()
	example3()
	example4()
}
