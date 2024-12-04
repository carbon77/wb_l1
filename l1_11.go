package main

import "fmt"

// Для реализации структуры данных множество (Set) будем использовать мапу,
// где ключом является элементы множеста
type Set map[int]bool

// Метод для вывода множества
func (s Set) String() string {
	result := "[ "

	for key, _ := range s {
		result += fmt.Sprintf("%d ", key)
	}

	result += "]"
	return result
}

func main() {
	s1 := Set{
		5: true,
		1: true,
		2: true,
	}

	s2 := Set{
		5: true,
		2: true,
		4: true,
		3: true,
	}

	intersection := Set{}

	// Для получение пересечения множеств удобнее итерироваться множеству, с меньшим количеством элементов
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	for key, _ := range s1 {
		// Если элемент, если в первом и во втором множестве, то добавляем в пересечением
		if s2[key] {
			intersection[key] = true
		}
	}

	fmt.Printf("Set 1: %v\n", s1)
	fmt.Printf("Set 2: %v\n", s2)
	fmt.Printf("Intersection: %v\n", intersection)
}
