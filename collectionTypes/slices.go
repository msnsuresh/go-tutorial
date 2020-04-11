package collectiontypes

import "fmt"

/*
PrintSliceOperation - Print Slice Operation exploration
*/
func PrintSliceOperation() {
	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)
}
