package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s\n", h.Name)
}

func (a *Action) GetInfo() {
	fmt.Printf("Name=%s, age=%d\n", a.Name, a.Age)
}

func main() {
	action := &Action{
		Human{
			Name: "Igor",
			Age:  21,
		},
	}

	action.Speak()
	action.GetInfo()
}
