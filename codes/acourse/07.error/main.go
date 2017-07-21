package main

import "fmt"
import "errors"

var (
	errAgeTooLow  = errors.New("age too low")
	errAgeTooHigh = errors.New("age too high")
)

func validateAge(age int) error {
	if age < 18 {
		// return false
		// return fmt.Errorf("age too low")
		return errAgeTooLow
	} else if age > 60 {
		// return false
		// return fmt.Errorf("age too high")
		return errAgeTooHigh
	}

	return nil
}

func main() {
	// fmt.Println(validateAge(5))
	// fmt.Println(validateAge(61))

	err := validateAge(70)

	// if err == errAgeTooLow {
	// 	fmt.Println("Can not enter")
	// 	return
	// }

	// if err == errAgeTooHigh {
	// 	fmt.Println(":)")
	// 	return
	// }

	if err != nil {
		fmt.Println(":(")
		return
	}
}
