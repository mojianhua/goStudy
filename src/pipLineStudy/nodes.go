package pipLineStudy

import "sort"

//其中 a ... int表示 定义a参数，有不定个int类型的数据 chan int 定义管道保存类型是a
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

			}
		}
	}()
}