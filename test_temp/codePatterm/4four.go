package codePatterm

import (
	"errors"
	"fmt"
)

//委托和反转控制
/***
其主要的思想是把控制逻辑与业务逻辑分享，不要在业务逻辑里写控制逻辑，这样会让控制逻辑依赖于业务逻辑，
而是反过来，让业务逻辑依赖控制逻辑。
**/

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}
func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}

	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil //For garbage collect
	}
	*undo = functions[:index]
	return nil
}

//IntSet 中嵌入Undo
type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}
func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func TestUndo() {
	intSet := NewIntSet()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(30)

	fmt.Println(intSet)

	intSet.Undo()
	fmt.Println(intSet)

	intSet.Add(20)
	intSet.Add(20)
	fmt.Println(intSet)
	intSet.Undo()
	fmt.Println(intSet)
}
