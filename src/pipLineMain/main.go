package main

import (
	"fmt"
	"pipLineStudy"
)

func main() {
	//讲数字放入管道后取值
	//p := pipLineStudy.ArraySource(3,2,4,56,7,8)
	//for v := range p{
	//	fmt.Print(v)
	//}

	//从管道里面取出数据，然后重新进行排序放入管道里面
	p1 := pipLineStudy.InMemSort(
		pipLineStudy.ArraySource(3,2,6,9,2,1))
	for v := range p1{
		fmt.Print(v)
	}
}
