package main

import (
	"fmt"
	"myapp/internal/calculator"
)

func main() {
	result, err := calculator.Divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}
}
