package part02

import "C"

import "fmt"

//下面得函数，不带参数哦
//export SayHelloGo
func SayHelloGo(s *C.char)  {
	fmt.Print(C.GoString(s))
}
