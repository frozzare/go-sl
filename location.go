package sl

import (
	"context"
	"errors"
)

// typeaheadEndpoint is the endpoint to the typeahead api.
const typeaheadEndpoint = "typeahead.json"

var (
	ErrNoSearchString = errors.New("Search string is empty")
)

// LocationService handles communication with the location related
// methods of the SL API.
//
// SL API docs: https://www.trafiklab.se/api/sl-platsuppslag/dokumentation
type LocationService service

// Location represents a location from the SL API.
type Location struct {
	Name   string `json:"Name"`
	SiteID string `json:"SiteId"`
	Type   string `json:"Type"`
	X      string `json:"X"`
	Y      string `json:"Y"`
}

// TypeaheadResponseData represents the typeahead response data SL API.
type TypeaheadResponseData struct {
	ExecutionTime int         `json:"ExecutionTime"`
	Message       string      `json:"Message"`
	ResponseData  []*Location `json:"ResponseData"`
	StatusCode    int         `json:"StatusCode"`
}

// LocationSearchOptions specifies optional parameters to the LocationSearch.Search.
type LocationSearchOptions struct {
	// Exclude buses if true. Default is false that are reversed to true.
	Bus bool `url:"bus,omitempty"`

	// API Key.
	Key string `url:"key,omitempty"`

	// Max results. Default is 10. Max 50.
	MaxResults bool `url:"maxResults,omitempty"`

	// SearchString.
	SearchString string `url:"searchstring,omitempty"`

	// Include only stations. Default is false.
	StationsOnly bool `url:"stationsonly,omitempty"`
}

// Search does a location lookup and response with the location list or a error.
func (s *LocationService) Search(ctx context.Context, opt *LocationSearchOptions) ([]*Location, error) {
	opt.StationsOnly = !opt.StationsOnly

	if len(opt.Key) == 0 {
		return nil, ErrNoKey
	}

	if len(opt.SearchString) == 0 {
		return nil, ErrNoSearchString
	}

	r, err := addOptions(typeaheadEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *TypeaheadResponseData
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Message) > 0 {
		return nil, errors.New(resp.Message)
	}

	return resp.ResponseData, nil
}
