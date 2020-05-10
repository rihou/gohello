package calc

import "errors"
import "github.com/fatih/color"

func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		errorMessage := color.RedString("provide more than 2 numbers")
		return errors.New(errorMessage), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum

	}
	
}