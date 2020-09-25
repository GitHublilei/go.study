package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 当需要频繁的发送请求的时候：定义一个全局的client,后续发请求的操作都使用这个全局的client

func f1(w http.ResponseWriter, r *http.Request) {
	var str []byte
	html, err := ioutil.ReadFile("./index.html")
	if err != nil {
		str = []byte(fmt.Sprintf("err:%v", err))
	}
	str = html
	w.Write(str)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo/", f2)
	err := http.ListenAndServe("127.0.0.1:9080", nil)
	if err != nil {
		fmt.Printf("server start failed, err:%v\n", err)
		return
	}
}
