package bigintpackages

func Add(a, b Bigint) Bigint {

	result := add(a.value , b.value)

	return Bigint{value: result, sing: true}
}