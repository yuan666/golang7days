package part02

//void SayHello(const char *s);		//声明也有分号
import "C"

func Cgo2_1_3()  {
	C.SayHello(C.CString("test 12345\n"))
}
