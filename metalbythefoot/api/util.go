package api

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parseXMLBody(body io.Reader) (*html.Node, error) {
	d := xml.NewDecoder(body)
	d.Strict = false
	var buf *bytes.Buffer = &bytes.Buffer{}
	e := xml.NewEncoder(buf)
	t, err := d.Token()
	for t != nil {
		if err != nil {
			return nil, err
		}
		if err = e.EncodeToken(t); err != nil {
			return nil, err
		}
		t, err = d.Token()
	}
	if err != io.EOF {
		return nil, err
	}
	if err = e.Flush(); err != nil {
		return nil, err
	}
	buf = bytes.NewBuffer(buf.Bytes())
	return htmlquery.Parse(buf)
}
