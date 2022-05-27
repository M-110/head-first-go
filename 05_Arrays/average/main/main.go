package main

import (
	"average/datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	averageValue := mean(numbers)
	fmt.Printf("Average: %0.2f\n", averageValue)
}

func mean(numbers [3]float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
