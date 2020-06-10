package api

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Document represents one document response from the server
type Document struct {
	API        *API
	Path       string
	Parameters map[string]string
	Response   *html.Node
}

// GetDocument retrieves a document from the server
func (api *API) GetDocument(path string, parameters map[string]string) (*Document, error) {
	params := make([]string, len(parameters))
	i := 0
	for k, v := range parameters {
		params[i] = fmt.Sprintf("%s=%s", k, v)
		i++
	}
	res, err := api.client.Get(fmt.Sprintf("http://www.metalbythefoot.com/configrator_ajax.php/%s?%s", path, strings.Join(params, "&")))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := parseXMLBody(res.Body)
	if err != nil {
		return nil, err
	}
	return &Document{
		API:        api,
		Path:       path,
		Parameters: parameters,
		Response:   doc,
	}, nil
}

// Mutate changes a set of parameters about the document and gets a new document from the server
func (doc *Document) Mutate(path string, parameters map[string]string) (*Document, error) {
	newParams := map[string]string{}
	for k, v := range doc.Parameters {
		newParams[k] = v
	}
	for k, v := range parameters {
		newParams[k] = v
	}
	if path == "" {
		path = doc.Path
	}
	return doc.API.GetDocument(path, newParams)
}
