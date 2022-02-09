package main

import (
	"fmt"
	"io"
	"net/http"
)

func aboutHttp() {
	fmt.Println("===")

	// 测试1
	// httpTest1()

	// 测试2
	// httpTest2()
}

// 测试1
func httpTest1() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// 回调函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")

	// 请求方式
	fmt.Println("method: ", r.Method)

	// /go
	fmt.Println("url: ", r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)

	// 回复
	w.Write([]byte("www.baidu.com"))
}


// 测试2
func httpTest2() {
	resp, _ := http.Get("http://www.0.0.1:8000/go")

	defer resp.Body.Close()

	// 200 ok
	fmt.Println("status: ", resp.Status)
	fmt.Println("header: ", resp.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}




