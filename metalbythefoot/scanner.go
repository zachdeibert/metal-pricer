package metalbythefoot

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zachdeibert/metal-pricer/metalbythefoot/api"
)

func isBySquareFeet(doc *api.Document) (bool, error) {
	hasLength := false
	hasWidth := false
	for _, field := range doc.GetShapeTextFieldNames() {
		switch field {
		case "length":
			hasLength = true
			break
		case "width":
			hasWidth = true
			break
		default:
			return false, fmt.Errorf("Unknown text field '%s'", field)
		}
	}
	if !hasLength {
		return false, errors.New("Document does not have a length field")
	}
	return hasWidth, nil
}

func scanPrice(doc *api.Document) (float64, error) {
	isSquare, err := isBySquareFeet(doc)
	if err != nil {
		return -1, err
	}
	width := "0"
	if isSquare {
		width = "12"
	}
	priceDoc, err := doc.Mutate(api.PriceEndpoint, map[string]string{
		"length": "1200",
		"width":  width,
	})
	if err != nil {
		return -1, err
	}
	return priceDoc.GetPrice() / 100, nil
}

func scanPriceRecursive(doc *api.Document, set map[string]interface{}) ([]map[string]interface{}, error) {
	next := doc.GetNextShapeOptionsName()
	if next == "" {
		var err error
		if set["price"], err = scanPrice(doc); err != nil {
			return nil, err
		}
		if set["isSquare"], err = isBySquareFeet(doc); err != nil {
			return nil, err
		}
		return []map[string]interface{}{set}, nil
	}
	options := doc.GetShapeOptions(next)
	res := []map[string]interface{}{}
	for _, option := range options {
		time.Sleep(100 * time.Millisecond)
		newDoc, err := doc.Mutate(api.ShapeEndpoint, map[string]string{
			fmt.Sprintf("oValue%s%s", strings.ToUpper(next[0:1]), next[1:]): option,
		})
		if err != nil {
			return nil, err
		}
		newSet := map[string]interface{}{next: option}
		for k, v := range set {
			newSet[k] = v
		}
		newRes, err := scanPriceRecursive(newDoc, newSet)
		if err != nil {
			return nil, err
		}
		res = append(res, newRes...)
	}
	return res, nil
}

// ScanPrices scans the server to get all the pricing information
func ScanPrices() ([]map[string]interface{}, error) {
	a := api.NewAPI()
	doc, err := a.GetDocument(api.ShapeEndpoint, map[string]string{})
	if err != nil {
		return nil, err
	}
	return scanPriceRecursive(doc, map[string]interface{}{})
}
