package pipLineStudy

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
)

//其中 a ... int表示 定义a参数，有不定个int类型的数据 chan int 定义管道保存类型是a
//内部排序
func ArraySource( a ...int) chan int {
	out := make(chan int)
	go func() {
		for _,v := range a{
			out <- v
		}
		close(out)
	}()
	return out
}

//内部排序
func InMemSort(in <-chan int) <-chan int{
	//定义一个管道
	out := make(chan int)
	go func() {
		//读取内存,定义变量a 是一个int 类型的数组
		a := []int{}
		for v := range in{
			//往a数组里面添加元素
			a = append(a,v)
		}
		//排序
		sort.Ints(a)
		//输出
		for _,v := range a{
			//将来V放入到out 管道里面
			out <- v
		}
		//关闭管道,否则报错
		close(out)
	}()
	//返回
	return out
}

//归并排序
func Merge(in1,in2 <-chan int)  <-chan int {
	out := make(chan int)
	go func() {
		v1,ok1 := <-in1
		v2,ok2 := <-in2
		for ok1 || ok2{
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				v1,ok1 = <-in2
			}else{
				out <- v2
				v2,ok2 = <-in2
			}
		}
		fmt.Print("合并结束")
		close(out)
	}()
	return out
}

//随机数生成生成count个int型数据
func RandomSource(count int) <-chan int  {
	out := make(chan int)
	go func() {
		for i := 0; i < count ;i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//读取数据
//第一个参数是读的来源对象，第二个参数是读取长度，（-1全读），输出一个chan
func RandomSourceFile(reader io.Reader , chunksize int) <- chan int  {
	out := make(chan int ,1024)
	go func() {
		// 64位系统的int类型大小是64，所以用一个64位buffer = byte(8) * 8
		buffer := make([]byte,8)
		// 读取长度的控制变量
		bytesread := 0
		for {
			// n是读取的长度
			n ,err := reader.Read(buffer)
			bytesread += n
			// 可能读取最好4字节数据，nul = EOF,再判读nil
			if n > 0{
				// 大端还是小端，发送和接收端统一即可
				out <- int(binary.BigEndian.Uint64(buffer))
			}
			if err !=nil || (chunksize != -1 && bytesread >= chunksize){
				break
			}
		}
		close(out)
	}()
	return out
}

//写数据，第一个参数写的目的对象，第二个参数是写数据channel
func WriteSink(wirter io.Writer,in <- chan int)  {
	for v:= range in{
		buffer :=make([]byte,8)
		binary.BigEndian.PutUint64(buffer,uint64(v))
		wirter.Write(buffer)
	}
}
