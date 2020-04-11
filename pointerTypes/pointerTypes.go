package pointertypes

import "fmt"

func printPointerFromAddressReference() {
	secondName := "Neelakandan"

	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Change text"
	fmt.Println(ptr, *ptr)
}

/*
PrintPointerTypes - Prints basic pointer exploration
*/
func PrintPointerTypes() {
	var firstName *string = new(string)
	*firstName = "Suresh from pointer"
	fmt.Println(*firstName)

	printPointerFromAddressReference()
}
