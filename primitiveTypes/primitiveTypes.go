package primitivetypes

import "fmt"

/*
PrintPrimitiveTypes - Prints all primitive types for exploration
*/
func PrintPrimitiveTypes() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Suresh"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
