package bigintpackages

import (
	"fmt"
	"strconv"
)

func Multiply(a, b Bigint) Bigint {

	var (
		resValue string = "0"
		resSign  bool   = true
	)

	if a.sign != b.sign {
		resSign = false
	}

	for i, pow := len(b.value)-1, 0; i >= 0; i, pow = i-1, pow+1 {

		bVal, _ := strconf.Atoi(string(b.value[i]))

		var (
			row  string = ""
			flag int    = 0
		)

		for j := len(a.value) - 1; j >= 0; j-- {
			aVal, _ := strconv.Atoi(string(a.value[j]))

			if j == 0 {
				row = fmt.Sprint(aVal*bVal+flag) + row
				continue
			}

			prod := (aVal*bVal + flag) % 10
			flag = (aVal*bVal + flag) / 10
			row = fmt.Sprint(prod) + row
		}

		for i := pow; i > 0; i-- {
			row = row + "0"
		}
		big, small, _ := bigSmall(Bigint{value: resValue}, Bigint{value: row})
		resValue = add(big.value, small.value)
	}

	return Bigint{value: resValue, sing: resSign}

}
