package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(res http.ResponseWriter, req *http.Request) {

	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	req.ParseForm()

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(res, "Hello astaxie!") //这个写入到w的是输出到客户端的

}

func login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	fmt.Println("我这里路由了login")

	req.ParseForm()
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(res, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
	}

	fmt.Println("我这里login路由结束啦")
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	println("Starting server on :9090")      //设置监听的端口')
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
