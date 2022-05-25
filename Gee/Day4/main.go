package main

import (
	"gee"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("xxx")

	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.RUN(":1234")
}

/***
===测试结果：
[root@localhost ~]# curl "http://localhost:1234/v1/hello?name=geektutu"
hello geektutu, you are at /v1/hello
[root@localhost ~]# curl "http://localhost:1234/v2/hello/geektutu"
hello geektutu, you're at /v2/hello/geektutu

分组，是指路由的分组。大部分情况下的路由分组，是以相同的前缀来区分的。

*/
