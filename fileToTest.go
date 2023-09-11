package main

import (
	"errors"
	"fmt"
)

func divider() {
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("The result of dividing is: ", result)
}

func divide(first, second float32) (float32, error) {
	var result float32

	if second == 0.0 {
		return result, errors.New("Cannot divide by 0")
	}
	result = first / second

	return result, nil
}
