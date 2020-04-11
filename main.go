package main

import (
	"fmt"

	collectiontypes "github.com/msnsuresh/gotutorial/collectionTypes"
	constanttypes "github.com/msnsuresh/gotutorial/constantTypes"
	pointertypes "github.com/msnsuresh/gotutorial/pointerTypes"
	primitivetypes "github.com/msnsuresh/gotutorial/primitiveTypes"
)

func main() {
	fmt.Println("Hello world from Pluralsight !")

	primitivetypes.PrintPrimitiveTypes()

	pointertypes.PrintPointerTypes()
	constanttypes.PrintConstantTypes()
	collectiontypes.PrintBasicArray()
	collectiontypes.PrintSliceOperation()
	collectiontypes.PrintMapOperation()
	collectiontypes.PrintStructOperations()
}
