package main

import (
	"bufio"
	"fmt"
	"os"
	"pipLineStudy"
)

func main() {
	//讲数字放入管道后取值
	//p := pipLineStudy.ArraySource(3,2,4,56,7,8)
	//for v := range p{
	//	fmt.Print(v)
	//}

	////从管道里面取出数据，然后重新进行排序放入管道里面（外部排序）
	//p1 := pipLineStudy.InMemSort(
	//	pipLineStudy.ArraySource(3,2,6,9,2,1))
	//for v := range p1{
	//	fmt.Print(v)
	//}

	////归并排序
	//p2 := pipLineStudy.Merge(
	//		pipLineStudy.InMemSort(
	//			pipLineStudy.ArraySource(6,5,2,6,7,8,7)),
	//		pipLineStudy.InMemSort(
	//			pipLineStudy.ArraySource(15,12,13,10,18,6,8,33)))
	//for v := range p2{
	//	fmt.Print(v)
	//}

	////生成随机数
	//p3 := pipLineStudy.RandomSource(99)
	//for v := range p3 {
	//	fmt.Print(v)
	//}

	//写入大文件
	const filename  = "small.in"
	const n  = 64

	//新建文件，返回可用的文件描述
	file ,err := os.Create(filename)
	if err != nil{
		panic(err)
	}

	//生成随机数
	p4 := pipLineStudy.RandomSource(n)

	//写入数据文件
	//包文件描述符，使用缓存机制，提高读写速度
	writer := bufio.NewWriter(file)
	pipLineStudy.WriteSink(writer,p4)
	writer.Flush()

	//上面的文件描述符，offset在末端，不能用于读取
	//打开文件
	file,err = os.Open(filename)
	if err !=nil{
		panic(err)
	}
	defer file.Close()

	//读取数据
	p5 := pipLineStudy.RandomSourceFile(bufio.NewReader(file),-1)
	count := 0
	for v := range p5{
		fmt.Println(v)
		count ++
		if count >= 100{
			break
		}
	}
}
