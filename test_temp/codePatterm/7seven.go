package codePatterm

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func TestDecorateHello() {
	decorator(Hello)("hello world")

	hello := decorator(Hello)
	hello("hello world")
}

//再看一个和计算运行时间的例子
type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timeSumFunc(f SumFunc) SumFunc {
	return func(start int64, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("----Time Elapsed (%s):%v\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

func sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}

	for i := start; i < end; i++ {
		sum += i
	}
	return sum
}

func sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (start + end) / 2
}

func TestDecorateSum() {
	s1 := timeSumFunc(sum1)
	s2 := timeSumFunc(sum2)

	fmt.Printf("%d,%d\n", s1(-10000, 100000),
		s2(-10000, 100000))
}

/////////////////////////////////////////////////////////////////
func withServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("---WithServerHeader()")
		w.Header().Set("Server", "Hello Server v1.0.0")
		h(w, r)
	}
}

func withAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("---withAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

func withBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--- withBasicAuth()")
		cookie, err := r.Cookie("Auth")
		if err != nil || cookie.Value != "Pass" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func withDebugLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("---withDebugLog()")
		r.ParseForm()
		log.Println(r.Form)
		log.Println("Path", r.URL.Path)
		log.Println("schema", r.URL.Scheme)
		log.Println(r.Form["url_log"])

		for k, v := range r.Form {
			log.Println("key:", k)
			log.Println("val:", strings.Join(v, ""))
		}
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World!"+r.URL.Path)
}
func TesthelloWithServer() {
	http.HandleFunc("/v1/hello", withServerHeader(withAuthCookie(hello)))
	http.HandleFunc("/v2/hello", withServerHeader(withBasicAuth(hello)))
	http.HandleFunc("/v3/hello", withServerHeader(withBasicAuth(withDebugLog(hello))))

	//pipeline
	http.HandleFunc("/v4/hello", Handler(hello, withServerHeader, withBasicAuth, withDebugLog))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenAndServe")
	}
}

//多个修饰器pipline
type HttpHandlerDecorator func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] //逆向
		h = d(h)
	}
	return h
}
