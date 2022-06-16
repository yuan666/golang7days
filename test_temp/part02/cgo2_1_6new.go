package part02

//void SaYHELLO(_GoString_ s);
import "C"
import "fmt"

func Cgo2_1_6_new() {
	C.SaYHELLO("hello xxxx")
}

//export SaYHELLO
func SaYHELLO(s string) {
	fmt.Print(s)
}
