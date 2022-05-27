package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64

	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	if err != nil {
		return numbers, err
	}

	for i := 0; i < 3; i++ {
		scanner.Scan()
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, err
	}
	return numbers, nil
}
