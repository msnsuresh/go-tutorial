package main

import (
	"fmt"

	constanttypes "github.com/msnsuresh/gotutorial/constantTypes"
	pointertypes "github.com/msnsuresh/gotutorial/pointerTypes"
	primitivetypes "github.com/msnsuresh/gotutorial/primitiveTypes"
)

func main() {
	fmt.Println("Hello world from Pluralsight !")

	primitivetypes.PrintPrimitiveTypes()

	pointertypes.PrintPointerTypes()
	constanttypes.PrintConstantTypes()
}
