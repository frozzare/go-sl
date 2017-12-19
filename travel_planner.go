package sl

import (
	"context"
	"errors"
	"fmt"
)

// travelPlannerEndpoint is the endpoint to the travel planner api.
const travelPlannerEndpoint = "TravelplannerV3/%s.json"

var (
	ErrNoTripFound = errors.New("No trip found")
)

// Trip represents a trip.
type Trip struct {
	LegList struct {
		Leg []struct {
			Destination struct {
				Date          string  `json:"date"`
				ExtID         string  `json:"extId"`
				HasMainMast   bool    `json:"hasMainMast"`
				ID            string  `json:"id"`
				Lat           float64 `json:"lat"`
				Lon           float64 `json:"lon"`
				MainMastExtID string  `json:"mainMastExtId"`
				MainMastID    string  `json:"mainMastId"`
				Name          string  `json:"name"`
				PrognosisType string  `json:"prognosisType"`
				Time          string  `json:"time"`
				Track         string  `json:"track"`
				Type          string  `json:"type"`
			} `json:"Destination"`
			JourneyDetailRef struct {
				Ref string `json:"ref"`
			} `json:"JourneyDetailRef"`
			JourneyStatus string `json:"JourneyStatus"`
			Origin        struct {
				Date          string  `json:"date"`
				ExtID         string  `json:"extId"`
				HasMainMast   bool    `json:"hasMainMast"`
				ID            string  `json:"id"`
				Lat           float64 `json:"lat"`
				Lon           float64 `json:"lon"`
				MainMastExtID string  `json:"mainMastExtId"`
				MainMastID    string  `json:"mainMastId"`
				Name          string  `json:"name"`
				PrognosisType string  `json:"prognosisType"`
				Time          string  `json:"time"`
				Track         string  `json:"track"`
				Type          string  `json:"type"`
			} `json:"Origin"`
			Product struct {
				Admin        string `json:"admin"`
				CatCode      string `json:"catCode"`
				CatIn        string `json:"catIn"`
				CatOut       string `json:"catOut"`
				CatOutL      string `json:"catOutL"`
				CatOutS      string `json:"catOutS"`
				Line         string `json:"line"`
				Name         string `json:"name"`
				Num          string `json:"num"`
				Operator     string `json:"operator"`
				OperatorCode string `json:"operatorCode"`
			} `json:"Product"`
			Category  string `json:"category"`
			Direction string `json:"direction"`
			Idx       string `json:"idx"`
			Name      string `json:"name"`
			Number    string `json:"number"`
			Reachable bool   `json:"reachable"`
			Type      string `json:"type"`
		} `json:"Leg"`
	} `json:"LegList"`
	ServiceDays []struct {
		PlanningPeriodBegin string `json:"planningPeriodBegin"`
		PlanningPeriodEnd   string `json:"planningPeriodEnd"`
		SDaysB              string `json:"sDaysB"`
		SDaysI              string `json:"sDaysI"`
		SDaysR              string `json:"sDaysR"`
	} `json:"ServiceDays"`
	TariffResult struct {
		FareSetItem []struct {
			Desc     string `json:"desc"`
			FareItem []struct {
				Cur   string `json:"cur"`
				Desc  string `json:"desc"`
				Name  string `json:"name"`
				Price int    `json:"price"`
			} `json:"fareItem"`
			Name string `json:"name"`
		} `json:"fareSetItem"`
	} `json:"TariffResult"`
	Checksum string `json:"checksum"`
	CtxRecon string `json:"ctxRecon"`
	Duration string `json:"duration"`
	Idx      int    `json:"idx"`
	TripID   string `json:"tripId"`
}

// TripResponseData represents the travel planner trip response data SL API.
type TripResponseData struct {
	ErrorCode string  `json:"errorCode"`
	ErrorText string  `json:"errorText"`
	Message   string  `json:"Message"`
	Trip      []*Trip `json:"Trip"`
	ScrB      string  `json:"scrB"`
	ScrF      string  `json:"scrF"`
}

// TravelPlannerService handles communication with the travel planner related
// methods of the SL API.
//
// SL API docs: https://www.trafiklab.se/node/16717/documentation
type TravelPlannerService service

// TripOptions specifies optional parameters to the TravelPlannerService.Trip.
type TripOptions struct {
	// Number of minutes added to estimated turnaround time.
	AddChangeTime int `url:"addChangeTime,omitempty"`

	// Separated by; in the following format: avoidId | avoidStatus
	// - avoidId, internal or external id for stop / station to avoid,
	// - avoidStatus, one of NPAVO (do not pass), NCAVO (do not change). Optionally.
	Avoid string `url:"avoid,omitempty"`

	// Internal or external ID for stop / station to avoid change.
	AvoidID string `url:"avoidID,omitempty"`

	// Percentage of original estimated time to handle a change.
	// Ex, 200 doubles the time the system will use for the traveler to catch up with a change.
	// Default 100.
	ChangeTimePercent int `url:"changeTimePercent,omitempty"`

	// Parameter that specifies the starting point for searching later or earlier trips.
	Context string `url:"context,omitempty"`

	// Trip date. Example: 2014-08-23. Default is today.
	Date string `url:"date,omitempty"`

	// Destination station id. Example: 300109600, 9600.
	DestID string `url:"destId,omitempty"`

	// Can either be a website id or an alias, website or acronym. Examples: 300109001, 9001, TCE.
	DestExtID string `url:"destExtId,omitempty"`

	// The destination coordinate in latitude.
	DestCoordLat string `url:"destCoordLat,omitempty"`

	// The destination coordinate in longitude.
	DestCoordLong string `url:"destCoordLong,omitempty"`

	// Indicates whether a trip can start with a walking distance. For distance sharing, min and max number of meters can be specified as 1, [min distance], [max distance].
	// Default is 1.
	DestWalk string `url:"destWalk,omitempty"`

	// API Key.
	Key string `url:"key,omitempty"`

	// Response language. Default 'sv', can be 'en' or 'de'.
	Lang string `url:"lang,omitempty"`

	// Line or lines separated by commas to be used to filter results, exclamation points are used for exclusion of lines.
	// Example: lines=55,122
	Lines string `url:"lines,omitempty"`

	// Max change (0-11).
	MaxChange int `url:"maxChange,omitempty"`

	// Max change time in minutes.
	MaxChangeTime int `url:"maxChangeTime,omitempty"`

	// Min change time in minutes.
	MinChangeTime int `url:"minChangeTime,omitempty"`

	// Min number of trips by the specified start time, default 4.
	// NumF and NumB together can not exceed 6.
	NumB string `url:"numB,omitempty"`

	// Min number of trips before the specified start time, default 1.
	// NumF and NumB together can not exceed 6.
	NumF string `url:"numF,omitempty"`

	// Limit the number of returned trips. Note that this is an approximate number. Default = 5
	NumTrips int `url:"numTrips,omitempty"`

	// The station to start at. Example: 300109600, 9600.
	OriginID string `url:"originId,omitempty"`

	// Can either be a website id or an alias, website or acronym. Examples: 300109001, 9001, TCE.
	OriginExtID string `url:"originExtId,omitempty"`

	// The start coordinate in latitude.
	OriginCoordLat string `url:"originCoordLat,omitempty"`

	// The start coordinate in longitude.
	OriginCoordLong string `url:"originCoordLong,omitempty"`

	// Indicates whether a trip can start with a walking distance. For distance sharing, min and max number of meters can be specified as 1, [min distance], [max distance].
	// Default is 1.
	OriginWalk string `url:"originWalk,omitempty"`

	// Indicates whether stops / stations passed on the trip should be retrieved. Default 0.
	Passlist int `url:"passlist,omitempty"`

	// Indicates whether detailed routes should be calculated for the results. 0 or 1. Default is 0.
	Poly int `url:"poly,omitempty"`

	// Combination value of desired traffic mode if not all will be used when traveling.
	Products int `url:"products,omitempty"`

	// By default, you are searching for the time you want the trip to resign.
	// By setting searchForArrival = 1, you will instead travel based on the time you want to reach.
	// Default = 0.
	SearchForArrival int `url:"searchForArrival,omitempty"`

	// Time. Example hour = 19:06. Default now.
	Time string `url:"time,omitempty"`

	// The station to pass through the station. Example: 300109600, 9600.
	ViaID string `url:"viaId,omitempty"`

	// Number of minutes to be spent on the via station indicated by ViaID.
	ViaWaitTime int `url:"viaWaitTime,omitempty"`
}

// Trip does a trip request to SL API and response with the trip list or a error.
func (s *TravelPlannerService) Trip(ctx context.Context, opt *TripOptions) ([]*Trip, error) {
	if len(opt.Key) == 0 {
		return nil, ErrNoKey
	}

	r, err := addOptions(fmt.Sprintf(travelPlannerEndpoint, "trip"), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *TripResponseData
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.ErrorText) > 0 {
		return nil, errors.New(resp.ErrorText)
	}

	if len(resp.Message) > 0 {
		return nil, errors.New(resp.Message)
	}

	return resp.Trip, nil
}

// Journey represents a journey.
type Journey struct {
	ErrorCode  string `json:"errorCode"`
	ErrorText  string `json:"errorText"`
	Message    string `json:"Message"`
	Directions struct {
		Direction []struct {
			RouteIdxFrom int    `json:"routeIdxFrom"`
			RouteIdxTo   int    `json:"routeIdxTo"`
			Value        string `json:"value"`
		} `json:"Direction"`
	} `json:"Directions"`
	JourneyStatus string `json:"JourneyStatus"`
	Names         struct {
		Name []struct {
			Product struct {
				Admin        string `json:"admin"`
				CatCode      string `json:"catCode"`
				CatIn        string `json:"catIn"`
				CatOut       string `json:"catOut"`
				CatOutL      string `json:"catOutL"`
				CatOutS      string `json:"catOutS"`
				Line         string `json:"line"`
				Name         string `json:"name"`
				Num          string `json:"num"`
				Operator     string `json:"operator"`
				OperatorCode string `json:"operatorCode"`
			} `json:"Product"`
			Category     string `json:"category"`
			Name         string `json:"name"`
			Number       string `json:"number"`
			RouteIdxFrom int    `json:"routeIdxFrom"`
			RouteIdxTo   int    `json:"routeIdxTo"`
		} `json:"Name"`
	} `json:"Names"`
	ServiceDays []struct {
		SDaysB string `json:"sDaysB"`
		SDaysI string `json:"sDaysI"`
		SDaysR string `json:"sDaysR"`
	} `json:"ServiceDays"`
	Stops struct {
		Stop []struct {
			DepDate          string  `json:"depDate"`
			DepPrognosisType string  `json:"depPrognosisType"`
			DepTime          string  `json:"depTime"`
			DepTrack         string  `json:"depTrack"`
			ExtID            string  `json:"extId"`
			HasMainMast      bool    `json:"hasMainMast"`
			ID               string  `json:"id"`
			Lat              float64 `json:"lat"`
			Lon              float64 `json:"lon"`
			MainMastExtID    string  `json:"mainMastExtId"`
			MainMastID       string  `json:"mainMastId"`
			Name             string  `json:"name"`
			RouteIdx         int     `json:"routeIdx"`
		} `json:"Stop"`
	} `json:"Stops"`
	LastPassRouteIdx int    `json:"lastPassRouteIdx"`
	LastPassStopRef  int    `json:"lastPassStopRef"`
	Ref              string `json:"ref"`
}

// JourneyOptions specifies optional parameters to the TravelPlannerService.Journey.
type JourneyOptions struct {
	// Trip date. Example: 2014-08-23. Default is today.
	Date string `url:"date,omitempty"`

	// The reference from Trip, see above.
	ID string `url:"id,omitempty"`

	// API Key.
	Key string `url:"key,omitempty"`

	// Indicates whether detailed routes should be calculated for the results. 0 or 1. Default is 0.
	Poly int `url:"poly,omitempty"`
}

// Journey does a journey request to SL API and response with the journey list or a error.
func (s *TravelPlannerService) Journey(ctx context.Context, opt *JourneyOptions) (*Journey, error) {
	if len(opt.Key) == 0 {
		return nil, ErrNoKey
	}

	r, err := addOptions(fmt.Sprintf(travelPlannerEndpoint, "journeydetail"), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *Journey
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.ErrorText) > 0 {
		return nil, errors.New(resp.ErrorText)
	}

	if len(resp.Message) > 0 {
		return nil, errors.New(resp.Message)
	}

	return resp, nil
}

// ReconstructionOptions specifies optional parameters to the TravelPlannerService.Reconstruction.
type ReconstructionOptions struct {
	// Trip date. Example: 2014-08-23. Default is today.
	Date string `url:"date,omitempty"`

	// The value I CtxRecon as I get the response from travel.
	Ctx string `url:"ctx,omitempty"`

	// API Key.
	Key string `url:"key,omitempty"`

	// Indicates whether detailed routes should be calculated for the results. 0 or 1. Default is 0.
	Poly int `url:"poly,omitempty"`
}

// Reconstruction does a reconstruction request to SL API and response with a trip or a error.
func (s *TravelPlannerService) Reconstruction(ctx context.Context, opt *ReconstructionOptions) (*Trip, error) {
	if len(opt.Key) == 0 {
		return nil, ErrNoKey
	}

	r, err := addOptions(fmt.Sprintf(travelPlannerEndpoint, "reconstruction"), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *TripResponseData
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.ErrorText) > 0 {
		return nil, errors.New(resp.ErrorText)
	}

	if len(resp.Message) > 0 {
		return nil, errors.New(resp.Message)
	}

	if len(resp.Trip) == 0 {
		return nil, ErrNoTripFound
	}

	return resp.Trip[0], nil
}
