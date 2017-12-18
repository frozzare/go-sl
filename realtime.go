package sl

import (
	"context"
	"errors"
)

// realtimeEndpoint is the endpoint to the realtime api.
const realtimeEndpoint = "realtimedeparturesV4.json"

// RealtimeService handles communication with the realtime related
// methods of the SL API.
//
// SL API docs: https://www.trafiklab.se/node/15754/documentation
type RealtimeService service

// Transport represents a transport type (bus, metro, ship, tram, train) struct.
type Transport struct {
	Destination string `json:"Destination"`
	Deviations  []struct {
		Consequence     string `json:"Consequence"`
		ImportanceLevel int    `json:"ImportanceLevel"`
		Text            string `json:"Text"`
	} `json:"Deviations"`
	DisplayTime          string `json:"DisplayTime"`
	ExpectedDateTime     string `json:"ExpectedDateTime"`
	GroupOfLine          string `json:"GroupOfLine"`
	JourneyDirection     int    `json:"JourneyDirection"`
	JourneyNumber        int    `json:"JourneyNumber"`
	LineNumber           string `json:"LineNumber"`
	StopAreaName         string `json:"StopAreaName"`
	StopAreaNumber       int    `json:"StopAreaNumber"`
	StopPointDesignation string `json:"StopPointDesignation"`
	StopPointNumber      int    `json:"StopPointNumber"`
	TimeTabledDateTime   string `json:"TimeTabledDateTime"`
	TransportMode        string `json:"TransportMode"`
}

// RealtimeResponse represents the realtime response from SL.
type RealtimeResponse struct {
	Buses               []*Transport `json:"Buses"`
	DataAge             int          `json:"DataAge"`
	LatestUpdate        string       `json:"LatestUpdate"`
	Metros              []*Transport `json:"Metros"`
	Ships               []*Transport `json:"Ships"`
	StopPointDeviations []struct {
		Deviation struct {
			Consequence     interface{} `json:"Consequence"`
			ImportanceLevel int         `json:"ImportanceLevel"`
			Text            string      `json:"Text"`
		} `json:"Deviation"`
		StopInfo struct {
			GroupOfLine    string `json:"GroupOfLine"`
			StopAreaName   string `json:"StopAreaName"`
			StopAreaNumber int    `json:"StopAreaNumber"`
			TransportMode  string `json:"TransportMode"`
		} `json:"StopInfo"`
	} `json:"StopPointDeviations"`
	Trains []*Transport `json:"Trains"`
	Trams  []*Transport `json:"Trams"`
}

// RealtimeResponseData represents the realtime response data SL API.
type RealtimeResponseData struct {
	ExecutionTime int               `json:"ExecutionTime"`
	Message       string            `json:"Message"`
	ResponseData  *RealtimeResponse `json:"ResponseData"`
	StatusCode    int               `json:"StatusCode"`
}

// RealtimeSearchOptions specifies optional parameters to the RealtimeService.Search.
type RealtimeSearchOptions struct {
	// Exclude buses if true. Default is false that are reversed to true.
	Bus bool `url:"bus,omitempty"`

	// Exclude metros if true. Default is false that are reversed to true.
	Metro bool `url:"metro,omitempty"`

	// API Key.
	Key string `url:"key,omitempty"`

	// Station ID.
	SiteID string `url:"siteId,omitempty"`

	// Exclude ships if true. Default is false that are reversed to true.
	Ship bool `url:"ship,omitempty"`

	// Exclude train if true. Default is false that are reversed to true.
	Train bool `url:"train,omitempty"`

	// Exclude trams if true. Default is false that are reversed to true.
	Tram bool `url:"tram,omitempty"`

	// Time window to search departures within. Max 60 minutes.
	TimeWindow int `url:"timeWindow,omitempty"`
}

// Search does a realtime search and response with the realtime list or a error.
func (s *RealtimeService) Search(ctx context.Context, opt *RealtimeSearchOptions) (*RealtimeResponse, error) {
	// Reverse transport options.
	opt.Bus = !opt.Bus
	opt.Metro = !opt.Metro
	opt.Ship = !opt.Ship
	opt.Train = !opt.Train
	opt.Tram = !opt.Tram

	if len(opt.Key) == 0 {
		return nil, ErrNoKey
	}

	if len(opt.SiteID) == 0 {
		return nil, ErrNoSiteID
	}

	r, err := addOptions(realtimeEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *RealtimeResponseData
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Message) > 0 {
		return nil, errors.New(resp.Message)
	}

	return resp.ResponseData, nil
}
