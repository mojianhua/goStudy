package pipLineStudy

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

//建立管道并且快速排序
func InMemSort(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		//读取内存,定义变量a 是一个int 类型的数组
		a := []int{}
		for v := range in{
			a = append(a )
		}
	}()
}