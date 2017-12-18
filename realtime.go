package sl

import (
	"context"
	"errors"
	"reflect"
)

// realtimeEndpoint is the endpoint to the realtime api.
const realtimeEndpoint = "realtimedeparturesV4.json"

// RealtimeService handles communication with the realtime related
// methods of the SL API.
//
// SL API docs: https://www.trafiklab.se/node/15754/documentation
type RealtimeService service

// RealtimeResponse represents the realtime response from SL.
type RealtimeResponse struct {
	Buses []struct {
		Destination          string      `json:"Destination"`
		Deviations           interface{} `json:"Deviations"`
		DisplayTime          string      `json:"DisplayTime"`
		ExpectedDateTime     string      `json:"ExpectedDateTime"`
		GroupOfLine          interface{} `json:"GroupOfLine"`
		JourneyDirection     int         `json:"JourneyDirection"`
		JourneyNumber        int         `json:"JourneyNumber"`
		LineNumber           string      `json:"LineNumber"`
		StopAreaName         string      `json:"StopAreaName"`
		StopAreaNumber       int         `json:"StopAreaNumber"`
		StopPointDesignation string      `json:"StopPointDesignation"`
		StopPointNumber      int         `json:"StopPointNumber"`
		TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
		TransportMode        string      `json:"TransportMode"`
	} `json:"Buses"`
	DataAge      int    `json:"DataAge"`
	LatestUpdate string `json:"LatestUpdate"`
	Metros       []struct {
		Destination          string      `json:"Destination"`
		Deviations           interface{} `json:"Deviations"`
		DisplayTime          string      `json:"DisplayTime"`
		ExpectedDateTime     string      `json:"ExpectedDateTime"`
		GroupOfLine          string      `json:"GroupOfLine"`
		JourneyDirection     int         `json:"JourneyDirection"`
		JourneyNumber        int         `json:"JourneyNumber"`
		LineNumber           string      `json:"LineNumber"`
		StopAreaName         string      `json:"StopAreaName"`
		StopAreaNumber       int         `json:"StopAreaNumber"`
		StopPointDesignation string      `json:"StopPointDesignation"`
		StopPointNumber      int         `json:"StopPointNumber"`
		TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
		TransportMode        string      `json:"TransportMode"`
	} `json:"Metros"`
	Ships               []interface{} `json:"Ships"`
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
	Trains []struct {
		Destination string `json:"Destination"`
		Deviations  []struct {
			Consequence     string `json:"Consequence"`
			ImportanceLevel int    `json:"ImportanceLevel"`
			Text            string `json:"Text"`
		} `json:"Deviations"`
		DisplayTime              string      `json:"DisplayTime"`
		ExpectedDateTime         string      `json:"ExpectedDateTime"`
		GroupOfLine              interface{} `json:"GroupOfLine"`
		JourneyDirection         int         `json:"JourneyDirection"`
		JourneyNumber            int         `json:"JourneyNumber"`
		LineNumber               string      `json:"LineNumber"`
		SecondaryDestinationName interface{} `json:"SecondaryDestinationName"`
		StopAreaName             string      `json:"StopAreaName"`
		StopAreaNumber           int         `json:"StopAreaNumber"`
		StopPointDesignation     string      `json:"StopPointDesignation"`
		StopPointNumber          int         `json:"StopPointNumber"`
		TimeTabledDateTime       string      `json:"TimeTabledDateTime"`
		TransportMode            string      `json:"TransportMode"`
	} `json:"Trains"`
	Trams []interface{} `json:"Trams"`
}

// ResponseData represents the response data SL API.
type ResponseData struct {
	ExecutionTime int               `json:"ExecutionTime"`
	Message       interface{}       `json:"Message"`
	ResponseData  *RealtimeResponse `json:"ResponseData"`
	StatusCode    int               `json:"StatusCode"`
}

// RealtimeSearchOptions specifies optional parameters to the RealtimeService.Search
type RealtimeSearchOptions struct {
	// Exclude buses if true. Default is false that are reversed to true.
	Bus bool

	// Exclude metros if true. Default is false that are reversed to true.
	Metro bool

	// API Key.
	Key string

	// Station ID.
	SiteID string

	// Exclude ships if true. Default is false that are reversed to true.
	Ship bool

	// Exclude train if true. Default is false that are reversed to true.
	Train bool

	// Exclude trams if true. Default is false that are reversed to true.
	Tram bool

	// Time window to search departures within. Max 60 minutes.
	TimeWindow int
}

// Search does a realtime search.
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

	var resp *ResponseData
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if reflect.ValueOf(resp.Message).Kind() == reflect.String {
		return nil, errors.New(resp.Message.(string))
	}

	return resp.ResponseData, nil
}
