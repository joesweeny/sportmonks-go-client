package sportmonks

import (
	"context"
	"fmt"
	"net/url"
)

const bookmakersUri = "/bookmakers"

type Bookmaker struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Logo *string `json:"logo"`
}

// Bookmakers sends a request and returns a slice of Bookmaker resource struct and supporting meta data.
func (c *HTTPClient) Bookmakers(ctx context.Context) ([]Bookmaker, *Meta, error) {
	response := struct {
		Data []Bookmaker `json:"data"`
		Meta *Meta       `json:"meta"`
	}{}

	err := c.getResource(ctx, bookmakersUri, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// BookmakerById sends a request and returns a single Bookmaker resource.
func (c *HTTPClient) BookmakerById(ctx context.Context, id int) (*Bookmaker, *Meta, error) {
	path := fmt.Sprintf(bookmakersUri+"/%d", id)

	response := struct {
		Data *Bookmaker `json:"data"`
		Meta *Meta      `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
