package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func makeSum(input *string, sign1 uint8, sign2 uint8) error {
	res := strings.Split(*input, "+")
	a, err := strconv.Atoi(strings.Trim(res[0], " "))
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(strings.Trim(res[1], " "))
	if err != nil {
		return err
	}
	if sign1 == '-' {
		a = -a
	}
	if sign2 == '-' {
		b = -b
	}
	*input = strconv.Itoa(a + b)
	return nil
}

func StringSum(input string) (output string, err error) {
	if (len(input)) == 0 {
		return "", fmt.Errorf("[StringSum] internal error: %w", errorEmptyInput)
	}
	if len(input) < 3 {
		return "", fmt.Errorf("[StringSum] internal error: %w", errorNotTwoOperands)
	}
	input = strings.Trim(input, " ")
	var sign uint8
	if (input[0] == '-') || (input[0] == '+') {
		sign = input[0]
		input = input[1:]
	}
	var e error
	if len(strings.Split(input, "+")) == 2 {
		e = makeSum(&input, sign, '+')
	} else if len(strings.Split(input, "-")) == 2 {
		e = makeSum(&input, sign, '-')
	} else {
		return "", fmt.Errorf("[StringSum] internal error: %w", errorNotTwoOperands)
	}
	if e != nil {
		return "", fmt.Errorf("[StringSum] internal error: %w", e)
	}
	return input, nil
}
