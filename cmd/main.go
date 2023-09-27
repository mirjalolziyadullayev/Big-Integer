package main

import (
	"bignum/bigintpackages"
	"fmt"
)

func main() {
	
	var a bigintpackages.Bigint = bigintpackages.NewBigInt("+546")
	var b bigintpackages.Bigint = bigintpackages.NewBigInt("-54")

	b.Abs()
	fmt.Println(bigintpackages.Add(a,b))
}