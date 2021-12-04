package main

import (
	"fmt"
	"net/http"
)

func main() {
	//默认页
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintln(writer, "<p>hello world</p>")
	//})

	//建立一个传参的页码
	//访问路径，http://127.0.0.1:9999/?name=123，结果打印：hello world 123
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<p>hello world %s</p>", request.FormValue("name"))
	})

	//建立服务
	http.ListenAndServe(":9999", nil)
}
