package main

import (
	"fmt"
	"pipLineStudy"
)

func main() {
	p := pipLineStudy.ArraySource(3,2,4,56,7,8)
	for v := range p{
		fmt.Print(v)
	}
}
