package part02

import "C"

//go build命令会在编译和链接阶段启动gcc编译器
func FirstCgo()  {
	println("hello cgo")
}
