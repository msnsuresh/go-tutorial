package collectiontypes

import "fmt"

/*
PrintStructOperations - Print Struct exploration
*/
func PrintStructOperations() {
	fmt.Println("*** Struct Exploration ***")
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "First"
	u.LastName = "Last"
	fmt.Println(u)

	u2 := user{ID: 2,
		FirstName: "Suresh",
		LastName:  "Neelakandan",
	}
	fmt.Println(u2)
}
