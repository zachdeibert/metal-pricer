package main

import (
	"fmt"

	"github.com/zachdeibert/metal-pricer/metalbythefoot/api"
)

func main() {
	a := api.NewAPI()
	doc, err := a.GetDocument(api.ShapeEndpoint, map[string]string{})
	if err != nil {
		panic(err)
	}
	fmt.Println(doc.GetShapeOptionsNames())
	fmt.Println(doc.GetShapeTextFieldNames())
	fmt.Println(doc.GetShapeOptions("type"))
	fmt.Println(doc.GetShapeOptions("shape"))
	fmt.Println(doc.GetShapeOptions("specific"))
	fmt.Println(doc.GetShapeOptions("thickness"))
}
