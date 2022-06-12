package codePatterm

import (
	"fmt"
	"reflect"
	"strings"
)

/***
MAP-REDUCE
	数据处理非常有用
map / filter / reduce 是相当好理解的：先变换，再过滤，最后总结。
***/
//写入两个Map函数，这两个函数都需要两个参数
//一个是字符串数组【】string，说明需要处理的数据
//另一个是一个函数
//map 转换
func MapStrToStr(arr []string, fn func(s string) string) []string {
	newArray := []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func TestMap01() {
	var list = []string{"Hao", "BabyCare", "Girl"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Printf("%v\n", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v\n", y)
}

//总结
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}

	return sum
}

func TestReduce() {
	var list = []string{"Hao", "BabyCare", "Girl"}

	x := Reduce(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v\n", x)
}

//过滤
func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func TestFilter() {
	var intSet = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out := Filter(intSet, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)

	out = Filter(intSet, func(n int) bool {
		return n >= 5
	})
	fmt.Printf("%v\n", out)
}

/////////////////////////////////////////////////////////////////
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

//reduce + filter
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) {
			count++
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList = []Employee{}
	for i, _ := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}
	return sum
}

//
func TestEmployee() {
	//统计有多少员工大于40岁
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people:%d\n", old)

	//统计多少员工的薪水大于6000
	highPay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary perople:%d\n", highPay)

	//列出没有休假的员工
	noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no Vacation:%v\n", noVacation)

	//统计30岁以下的员工的薪水总和
	youngerPay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("Younger People Salary Sum:%d\n", youngerPay)
}

/////////////////////////////////////////////////////////
//泛型版本 的Map-Reduce  go目前不支持泛型，目前使用 interface{} + reflect来实现

//非常简单，不做任何类型检查的泛型Map函数
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)

	result := make([]interface{}, vdata.Len())
	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func TestMap() {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	square_arr := Map(nums, square)
	fmt.Println(square_arr)

	upcase := func(s string) string {
		return strings.ToUpper(s)
	}

	strs := []string{"Hao", "BabyCare", "Girl"}
	upStrs := Map(strs, upcase)
	fmt.Println(upStrs)
}

//健壮版的 Generic Map 增加类型的检测
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {

	//check the <code data-enlighter-language="raw" class="EnlighterJSRAW">slice</code> type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

func TestGeneraMap() {
	list := []string{"1", "2", "3", "4", "5"}
	result := Transform(list, func(s string) string {
		return s + s + s
	})

	fmt.Println(result)
}


//健壮版的 Generic Reduce
func GenericReduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}
	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}
	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}
	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]
	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

//健壮版 Generic Filter
func GenericFilter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}
func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}
var boolType = reflect.ValueOf(true).Type()
func filter(slice, function interface{}, inPlace bool) (interface{}, int) {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}
	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}
	out := sliceInType
	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}
	return out.Interface(), len(which)
}