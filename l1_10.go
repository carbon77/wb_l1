package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, t := range temperatures {
		// Округляем число до ближайшего числа кратного 10 вниз и добавляем в группу
		group := int(t/10) * 10
		groups[group] = append(groups[group], t)
	}

	fmt.Println(groups)
}
