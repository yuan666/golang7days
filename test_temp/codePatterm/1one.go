package codePatterm

import (
	"bytes"
	"fmt"
	"reflect"
)

func printSliceInfo(name string, b []byte) {
	fmt.Printf("[%s] len:%d cap:%d \n", name, len(b), cap(b))
}
func TestSlice() {
	path := []byte("AAAA/BBBBBBBB") //切片 len=13 cap=13
	printSliceInfo("path", path)

	sepIdx := bytes.IndexByte(path, '/') //4
	dir1 := path[:sepIdx]                //len=4 cap 13
	dir2 := path[sepIdx+1:]              //len=8 cap=8
	printSliceInfo("dir1", dir1)
	printSliceInfo("dir2", dir2)

	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))
	//"suffix"... 将数组转换为单成员，然后追加进去
	dir1 = append(dir1, "suffix"...) //len=10 cap 13
	printSliceInfo("dir1", dir1)
	printSliceInfo("dir2", dir2)

	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

}

func TestDeepEqual() {
	m1 := map[string]string{"one": "1", "two": "2"}
	m2 := map[string]string{"two": "2", "one": "1"}

	fmt.Println("m1==m2:", reflect.DeepEqual(m1, m2))
}

//把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦
type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

var _ Stringable = (*Country)(nil)
var _ Stringable = (*City)(nil)
func (c *Country) ToString()string  {
	return "Country="+c.Name
}
func (c *City)ToString()string  {
	return "City="+c.Name
}

func PrintStr(p Stringable)  {
	fmt.Println(p.ToString())
}

func TestStringable()  {
	d1 := &Country{"China"}
	d2 := &City{"beijing"}

	PrintStr(d1)
	PrintStr(d2)
}