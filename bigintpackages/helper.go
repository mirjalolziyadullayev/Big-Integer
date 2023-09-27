package bigintpackages

import (
	"fmt"
	"strconv"
)

func add(big, small string) string {

	for len(small) < len(big) {
		small = "0" + small
	}

	var flag int = 0
	var resString string
	var res int

	for i := len(small) - 1; i >= 0; i-- {
		bigI, err := strconv.Atoi(string(big[i]))

		if err != nil {
			fmt.Println(err)
			return ""
		}
		smallI, err := strconv.Atoi(string(small[i]))

		res = (bigI + smallI + flag) % 10
		flag = (bigI + smallI + flag) / 10

		if i == 0 {
			res = bigI + smallI + flag
		}

		resString = fmt.Sprint(res) + resString
	}

	return resString
}

func sub(big, small string) string {
	return ""
}
