package api

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

// GetShapeOptions gets a list of all options available for a shape
func (doc *Document) GetShapeOptions(name string) []string {
	nodes := htmlquery.Find(doc.Response, fmt.Sprintf(`//form/div[@id="%s"]//option/@value[.!=""]`, name))
	options := make([]string, len(nodes))
	for i, node := range nodes {
		options[i] = fmt.Sprint(node.FirstChild.Data)
	}
	return options
}

// GetShapeOptionsNames gets the names of all the options available for a shape
func (doc *Document) GetShapeOptionsNames() []string {
	nodes := htmlquery.Find(doc.Response, `//form/div[.//select]/@id`)
	options := make([]string, len(nodes))
	for i, node := range nodes {
		options[i] = fmt.Sprint(node.FirstChild.Data)
	}
	return options
}

// GetNextShapeOptionsName gets the name of the next option available for a shape
func (doc *Document) GetNextShapeOptionsName() string {
	nodes := htmlquery.Find(doc.Response, `//form/div[.//select[not(option/@selected)]][1]/@id`)
	switch len(nodes) {
	case 0:
		return ""
	case 1:
		return nodes[0].FirstChild.Data
	default:
		panic("Invalid XPath result")
	}
}

// GetShapeTextFieldNames gets the names of all the text fields available for a shape
func (doc *Document) GetShapeTextFieldNames() []string {
	nodes := htmlquery.Find(doc.Response, `//form/div[.//input[@type="text"]]/@id`)
	options := make([]string, len(nodes))
	for i, node := range nodes {
		options[i] = fmt.Sprint(node.FirstChild.Data)
	}
	return options
}
