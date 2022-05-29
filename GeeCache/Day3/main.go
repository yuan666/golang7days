package main

import (
	"fmt"
	"geecache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "879",
	"Jack": "233",
	"Same": "245",
}

func main() {
	geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)

			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:1234"
	peers := geecache.NewHTTPPool(addr)

	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
/*****
介绍如何使用 Go 语言标准库 http 搭建 HTTP Server
并实现 main 函数启动 HTTP Server 测试 API

测试结果：
~ curl http://localhost:1234/_geecache/scores/Tom
879%
➜  ~ curl http://localhost:1234/_geecache/scores/Tom
879%
➜  ~ curl http://localhost:1234/_geecache/scores/xx
xx not exist

Day3 git:(master) ✗ ./Day3
2022/05/29 22:56:36 geecache is running at localhost:1234
2022/05/29 22:57:35 [Server localhost:1234] GET /_geecache/scores/Tom
2022/05/29 22:57:35 [SlowDB] search key Tom
2022/05/29 22:57:45 [Server localhost:1234] GET /_geecache/scores/Tom
2022/05/29 22:57:45 [GeeCache] hit
2022/05/29 22:57:52 [Server localhost:1234] GET /_geecache/scores/xx
2022/05/29 22:57:52 [SlowDB] search key xx


**/
