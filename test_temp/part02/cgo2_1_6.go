package part02
//注意声明得时候，要使用c得代码哦  char *s


//void SayHELLO(char *s);
import "C"
import "fmt"

func Cgo2_1_6()  {
	C.SayHELLO(C.CString("hello world!.....\n"))
}

//export SayHELLO
func SayHELLO(s *C.char)  {
	fmt.Print(C.GoString(s))
}

