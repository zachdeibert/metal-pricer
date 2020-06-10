package api

import (
	"net/http"
)

// API for MetalByTheFoot website
type API struct {
	client *http.Client
}

// NewAPI creates a new API instance
func NewAPI() *API {
	return &API{
		client: &http.Client{},
	}
}
