package main

import (
	"fmt"

	"github.com/vishymails/golangexamples/simpleshape"
)

func main() {
	r := simpleshape.NewRectangle(9, 6)
	t := simpleshape.NewTriangle(3, 6)

	fmt.Println("Area : ", simpleshape.ShapeArea(r))

	fmt.Println("Area : ", simpleshape.ShapeArea(t))

}
