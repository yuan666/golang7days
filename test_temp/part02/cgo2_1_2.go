package part02

//#include <stdio.h>
import "C"

func Cgo2_1_2()  {
	C.puts(C.CString("hello world"))
}
