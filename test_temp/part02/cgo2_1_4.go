package part02

//void SayHelloCPP(const char *s);
import "C"

func Cgo2_1_4()  {
	C.SayHelloCPP(C.CString("test 12345\n"))
}
