package main

import "fmt"

func main() {
	var input Year
	fmt.Print("Input: ")
	fmt.Scan(&input.firstInput)

	fmt.Println(input.firstInput)
}

type Year struct {
	firstInput string
}

func (y Year) leapYearDeterminer() (string, error) {
	return y.firstInput, nil
}
