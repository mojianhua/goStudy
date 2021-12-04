package main

import (
	"fmt"
	//"time"
)

func main() {
	//定义一个管道
	ch := make(chan string)
	//并发执行
	for i := 0; i < 5000; i++ {
		go printhelloworld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}
	//因为并发太快所以。让他睡1毫秒
	//time.Sleep(time.Microsecond)
	//time.Sleep(time.Second)
}

////go中跟func传参
//func printhelloworld(i int) {
//	//fmt.Print("hello world")
//
//	//因为用了formait参数所以要用printf
//	fmt.Printf("hello world %d\n", i)
//}

//并发中返回之值，定义chan管道，可以实现管道之间通信，
func printhelloworld(i int, ch chan string) {
	for {
		//把打印发送到ch管道里面
		ch <- fmt.Sprintf("hello world form "+"go rountine %d!\n", i)
		//
		////把v发送到管道ch
		//ch <-v
		////把ch接收数据，并且赋值到v上
		//v := <-ch
	}
}
