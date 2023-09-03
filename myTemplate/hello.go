package myTemplate

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("myTemplate/hello.html")
	if err != nil {
		fmt.Println("create myTemplate failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "枯藤",
		Gender: "男",
		Age:    18,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}
func Server() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
