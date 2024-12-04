package main

import "fmt"

// Для реализации структуры данных множество будем использовать мапу,
// где ключом является элементы множеста
type MySet struct {
	data map[string]bool
}

func NewMySet() *MySet {
	return &MySet{
		data: make(map[string]bool),
	}
}

func (ms *MySet) Add(s string) {
	ms.data[s] = true
}

func (ms *MySet) Delete(s string) {
	delete(ms.data, s)
}

func (ms *MySet) Exists(s string) bool {
	_, ok := ms.data[s]
	return ok
}

func (ms MySet) String() string {
	result := "[ "
	for s, _ := range ms.data {
		result += s + " "
	}
	result += "]"
	return result
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewMySet()

	for _, word := range words {
		set.Add(word)
	}

	fmt.Printf("Set: %v\n", set)
	fmt.Printf("Does 'cat' exists: %v\n", set.Exists("cat"))
	fmt.Printf("Does 'car' exists: %v\n", set.Exists("car"))
	fmt.Println("Deleting 'cat'")
	set.Delete("cat")
	fmt.Printf("Does 'cat' exists: %v\n", set.Exists("cat"))
}
