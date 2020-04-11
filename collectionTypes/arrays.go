package collectiontypes

import "fmt"

/*
PrintBasicArray - Prints basic array exploration
*/
func PrintBasicArray() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr[:]

	arr[1] = 10
	slice[2] = 30
	fmt.Println(arr, slice)
}
