package service

type Foo struct {
	Name string
}

func NewFoo() *Foo {
	return &Foo{Name: "I am Foo"}
}