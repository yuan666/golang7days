package main

import (
	"net/http"
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s \n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.RUN(":1234")

}

/***********
测试效果：

===========================
前缀树路由 trie router
	使用 Trie 树实现动态路由(dynamic route)解析。
	支持两种模式:name和*filepat

  main.go
	gee
		context.go
		router.go
		gee.go
		trie.go
**/
