package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s \n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.RUN(":1234")
}
/****
测试效果：
[root@localhost ~]#  curl -i http://localhost:1234
HTTP/1.1 200 OK
Content-Type: text/html
Date: Sun, 22 May 2022 03:25:54 GMT
Content-Length: 18

<h1>hello Gee</h1>[root@localhost ~]# curl "http://localhost:1234/hello?name=geektutu"
hello geektutu,you are at /hello

[root@localhost ~]# curl "http://localhost:1234/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}
[root@localhost ~]#
[root@localhost ~]#
[root@localhost ~]# curl "http://localhost:1234/world"
404 NOT FOUND /world
=======================================
将路由(router)独立出来，方便之后增强。
设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
  main.go
	gee
		context.go
		router.go
		gee.go
*/