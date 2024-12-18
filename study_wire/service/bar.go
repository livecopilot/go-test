package service

type Bar struct {
	Foo *Foo
}

func NewBar(foo *Foo) *Bar {
	return &Bar{Foo: foo}
}