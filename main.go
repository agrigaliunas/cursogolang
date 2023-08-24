package main

import (
	"fmt"

	"github.com/agrigaliunas/cursogolang/ejercicios"
)

func main() {
	var returnedInt int
	var returnedText string

	returnedInt, returnedText = ejercicios.StringToInt("1000")

	fmt.Println(returnedInt, " ", returnedText)

	returnedInt, returnedText = ejercicios.StringToInt("10")

	fmt.Print(returnedInt, " ", returnedText)

}
