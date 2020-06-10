package api

import (
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

// GetPrice gets the price information from the document
func (doc *Document) GetPrice() float64 {
	nodes := htmlquery.Find(doc.Response, `//div[@id="show_price2"]/div[span[1]/text() = "Estimated Total for Item :"]/span[2]/text()`)
	if len(nodes) != 1 {
		panic("Error parsing price document")
	}
	priceStr := nodes[0].Data
	priceStr = strings.TrimSpace(priceStr)
	lenBefore := len(priceStr)
	priceStr = strings.TrimLeft(priceStr, "$")
	if len(priceStr) != lenBefore-1 {
		panic("Error parsing price")
	}
	priceStr = strings.TrimSpace(priceStr)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		panic(err)
	}
	return price
}
