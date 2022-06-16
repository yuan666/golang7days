package part02

//包含 下面.h文件，还有hello.go

//#include <hellogo.h>
import "C"

func Cgo2_1_5()  {
	C.SayHelloGo(C.CString("1111111111111\n"))
}