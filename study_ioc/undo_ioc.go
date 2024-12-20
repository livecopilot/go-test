package main

import "fmt"


type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo

	if len(functions) == 0 {
		return fmt.Errorf("no functions to undo")
	}

	index := len(functions) - 1

	if function := functions[index]; function != nil {
		functions[index] = nil
		function()
	}

	*undo = functions[:index]

	return nil
}


type IntSet struct {
	data map[int]bool
	undo Undo
}


func NewIntSet() IntSet {
	return IntSet{make(map[int]bool), nil}
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}


func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}
