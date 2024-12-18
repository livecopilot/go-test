package main

import "fmt"

func main() {
	bar := InitializeBar()
	fmt.Println(bar.Foo.Name)
}