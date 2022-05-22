package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path:%q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})

	r.RUN(":1234")
}

/***
测试效果：
[root@localhost ~]# curl http://localhost:1234
URL.Path:"/"
[root@localhost ~]# curl http://localhost:1234/hello
Header["User-Agent"]=["curl/7.29.0"]
Header["Accept"]=["*//*"]
===========================================
	整个Gee框架的原型
	用户注册静态路由
	包装了启动服务的函数
*/