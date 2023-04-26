package main

import (
	"errors"
	"fmt"
)

var (
	errUhOh = errors.New("uh oh")
)

type ValueError struct {
	Value int
	Err   error
}

func newValueError(value int, err error) *ValueError {
	return &ValueError{
		Value: value,
		Err:   err,
	}
}

func (ve *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", ve.Err)
}

func (ve *ValueError) Unwrap() error {
	return ve.Err
}

func validateValue(number int) error {
	if number == 1 {
		return newValueError(number, fmt.Errorf("this odd"))
	} else if number == 2 {
		return newValueError(number, errUhOh)
	}
	return nil

}

func runValidation(number int) error {
	err := validateValue(number)
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}

func main() {
	for num := 1; num <= 3; num++ {
		fmt.Printf("validating %d... ", num)
		err := runValidation(num)
		var valueErr *ValueError
		if errors.Is(err, errUhOh) {
			fmt.Println("oh no!")
		} else if errors.As(err, &valueErr) {
			fmt.Printf("value error (%d): %v\n", valueErr.Value, valueErr.Err)
		} else {
			fmt.Println("valid!")
		}
	}
}
