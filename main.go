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
	doc2, err := doc.Mutate(api.PriceEndpoint, map[string]string{
		"length":          "1200",
		"width":           "0",
		"oValueType":      "Aluminum",
		"oValueShape":     "Angle",
		"oValueSpecific":  "6061",
		"oValueThickness": ".125  (1/8)",
		"oValueSize":      ".75 x .75  (3/4 x 3/4)",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(doc.GetShapeOptionsNames())
	fmt.Println(doc.GetShapeTextFieldNames())
	fmt.Println(doc.GetShapeOptions("type"))
	fmt.Println(doc.GetShapeOptions("shape"))
	fmt.Println(doc.GetShapeOptions("specific"))
	fmt.Println(doc.GetShapeOptions("thickness"))
	fmt.Println(doc2.GetPrice())
}
