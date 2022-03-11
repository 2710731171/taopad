package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------")
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_ long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello GoWeb!") //这个写入到w的是输出到客户端的
}
func main() {
	/**
	  第一个参数：pattern string,
	  第二个参数：handler func(ResponseWriter, *Request)
	*/
	http.HandleFunc("/", sayHello) // 设置访问的路由
	/**
	  第一个参数addr：监听地址
	  第二个参数handler：通常为空，意味着服务端调用http.DefaultServerMux进行处理，而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMux中
	*/
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
