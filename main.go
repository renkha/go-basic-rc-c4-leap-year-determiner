package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input Year
	fmt.Print("Input first year: ")
	fmt.Scan(&input.firstInput)
	fmt.Print("Input last year: ")
	fmt.Scan(&input.lastInput)

	leapYear, err := input.leapYearDeterminer()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Leap's year: ")
	for _, value := range leapYear {
		fmt.Println(value)
	}
}

type Year struct {
	firstInput string
	lastInput  string
}

func (y Year) leapYearDeterminer() ([]int, error) {
	firstInput, err1 := strconv.Atoi(y.firstInput)
	lastInput, err2 := strconv.Atoi(y.lastInput)
	var output []int

	//errror handler
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Input isn't number")
	}

	//swap value to make the lower value only for first variable
	if lastInput < firstInput {
		firstInput, lastInput = lastInput, firstInput
	}

	//mod by 4 equal to 0
	for i := firstInput; i <= lastInput; i++ {
		if i%4 == 0 {
			output = append(output, i)
		}
	}

	//for every new century that mod by 100 and 400 equal to 0
	for j := 0; j < len(output); j++ {
		if output[j]%100 == 0 && output[j]%400 != 0 {
			output = append(output[:j], output[j+1:]...)
		}
	}

	return output, nil
}
