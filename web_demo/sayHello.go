package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 * @Project GoWebDemo
 * @File    sayHello.go
 * @Author  Augus Lee
 * @Date    2022/11/9 15:05
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// http.ResponseWriter：代表响应，传递到前端的
// *http.Request：表示请求，从前端传递过来的
func sayHello(w http.ResponseWriter, r *http.Request) {
	html, _ := os.ReadFile("./hello.html")
	_, _ = fmt.Fprintln(w, string(html))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed, err: ", err)
		return
	}
}
