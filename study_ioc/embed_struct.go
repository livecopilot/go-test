package main

import (
	"fmt"
)

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
}

type Button struct {
	Label // embedding
}

type ListBox struct {
	Widget
	Texts []string
	Index int
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	println("Label Paint", label.X, label.Y, label.Text)
}

func (button Button) Paint() {
	println("Button Paint", button.X, button.Y, button.Text)
}

func (button Button) Click() {
	println("Button Click", button.X, button.Y, button.Text)
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}

func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func NewButton(x, y int, text string) Button {
	return Button{Label{Widget{x, y}, text}}
}

// type IntSet struct {
// 	data map[int]bool
// }

// func NewIntSet() IntSet {
// 	return IntSet{make(map[int]bool)}
// }

// func (set *IntSet) Add(x int) {
// 	set.data[x] = true
// }

// func (set *IntSet) Delete(x int) {
// 	delete(set.data, x)
// }

// func (set *IntSet) Contains(x int) bool {
// 	return set.data[x]
// }

// type UndoAbleIntSet struct {
// 	IntSet
// 	functions []func()
// }

// func NewUndoAbleIntSet() UndoAbleIntSet {
// 	return UndoAbleIntSet{NewIntSet(), nil}
// }

// func (set *UndoAbleIntSet) Add(x int) {
// 	if !set.Contains(x) {
// 		set.data[x] = true
// 		set.functions = append(set.functions, func() { set.Delete(x) })
// 	} else {
// 		set.functions = append(set.functions, nil)
// 	}
// }

// func (set *UndoAbleIntSet) Delete(x int) {
// 	if set.Contains(x) {
// 		delete(set.data, x)
// 		set.functions = append(set.functions, func() { set.Add(x) })
// 	} else {
// 		set.functions = append(set.functions, nil)
// 	}
// }

// func (set *UndoAbleIntSet) Undo() error {
// 	if len(set.functions) == 0 {
// 		return fmt.Errorf("no functions to undo")
// 	}

// 	index := len(set.functions) - 1

// 	if function := set.functions[index]; function != nil {
// 		function()
// 		set.functions[index] = nil
// 	}

// 	set.functions = set.functions[:index]

// 	return nil
// }
