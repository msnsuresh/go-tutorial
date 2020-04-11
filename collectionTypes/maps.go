package collectiontypes

import "fmt"

/*
PrintMapOperation - Prints Map exploration
*/
func PrintMapOperation() {
	m := map[string]int{"foo": 42, "foobar": 30}

	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}
