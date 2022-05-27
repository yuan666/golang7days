package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.HTML(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main()  {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.RUN(":1234")
}

/********
测试
[root@localhost ~]# curl http://localhost:1234/v2/hello/geektutu
hello geektutu, you're at /v2/hello/geektutu


2022/05/27 08:33:29 Route  GET - /v2/hello/:name
2022/05/27 08:33:32 [0] /v2/hello/geektutu in 208ns for group v2
2022/05/27 08:33:32 [200] /v2/hello/geektutu in 40.653µs

day5
	设计并实现 Web 框架的中间件(Middlewares)机制。
	实现通用的Logger中间件，能够记录请求到响应所花费的时间

	中间件(middlewares)，简单说，就是非业务的技术类组件。

	中间件，考虑两个关键点：
		插入点在哪？使用框架的人并不关心底层逻辑的具体实现，如果插入点太底层，中间件逻辑就会非常复杂。如果插入点离用户太近，那和用户直接定义一组函数，每次在 Handler 中手工调用没有多大的优势了。
		中间件的输入是什么？中间件的输入，决定了扩展能力。暴露的参数太少，用户发挥空间有限。

	接收到请求后，应查找所有应作用于该路由的中间件，保存在Context中，依次进行调用。

*/
