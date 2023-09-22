package main

import (
	"fmt"
	"math"
)

func main() {
	// задаем исходную последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаем карту для хранения групп температурных колебаний
	groups := make(map[int][]float64)

	// проходим по всем температурам и добавляем их в соответствующую группу
	for _, temperature := range temperatures {
		group := int(math.Floor(temperature/10.0)) * 10
		groups[group] = append(groups[group], temperature)
	}

	// выводим результаты
	for group, temperatures := range groups {
		fmt.Printf("%d: %v\n", group, temperatures)
	}
}
