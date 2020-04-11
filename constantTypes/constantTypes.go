package constanttypes

import "fmt"

const (
	one = iota + 1
	two
	four = iota + 2
	five
)

/*
PrintConstantTypes - Print constant exploration
*/
func PrintConstantTypes() {
	const (
		zero = iota
		first
		second
	)

	fmt.Println(zero, first, second)
	fmt.Println(one, two, four, five)
}
