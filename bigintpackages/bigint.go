package bigintpackages

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Bigint struct {
	value string
	sing  bool
}

var Badinput error = errors.New("Bad input")

func NewBigInt(son string) Bigint {

	err := validate(son)
	if err != nil {
		fmt.Println(err)
		return Bigint{}
	}

	numSimplified, ishora := simplify(son)

	return Bigint{
		value: numSimplified,
		sing:  ishora,
	}
}

func validate(num string) error {

	allowed := "+-0123456789"

	if !strings.Contains(allowed, string(num[0])) {
		return Badinput
	}

	for i := 1; i < len(num); i++ {
		if !unicode.IsDigit(rune(num[i])) {
			return Badinput
		}
	}
	return nil
}

func simplify(num string) (string, bool) {

	var sign bool = true

	if num[0] == '-' {
		num = num[1:]
		sign = false
	}

	if num[0] == '+' {
		num = num[1:]
	}

	for num[0] == '0' {
		num = num[1:]
	}

	return num, sign
}
