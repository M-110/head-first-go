package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	averageValue := mean()
	fmt.Printf("Average: %0.2f\n", averageValue)
}

func mean() float64 {
	args := os.Args[1:]
	var sum float64 = 0
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	return sum / float64(len(args))
}
