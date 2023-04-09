package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	*Human // встраиваем родительскую структуру
}

func (a *Action) Walk() {
	fmt.Printf("%s is walking.\n", a.GetName())
}

func (a *Action) Run() {
	fmt.Printf("%s is running.\n", a.GetName())
}

func main() {
	h := &Human{name: "Pasha", age: 21}
	a := &Action{Human: h} // инициализируем вложенную структуру Action с родительской структурой Human

	a.Walk()
	a.Run()
}
