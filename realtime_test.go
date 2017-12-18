package sl

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRealtimeSearch(t *testing.T) {
	client, mux, _, teardown := setupClient()
	defer teardown()
	mux.HandleFunc("/realtimedeparturesV4.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"StatusCode":0,"Message":null,"ExecutionTime":621,"ResponseData":{"LatestUpdate":"2017-12-18T20:10:19","DataAge":12,"Metros":[{"GroupOfLine":"tunnelbanans blå linje","DisplayTime":"Nu","TransportMode":"METRO","LineNumber":"11","Destination":"Akalla","JourneyDirection":1,"StopAreaName":"T-Centralen","StopAreaNumber":1051,"StopPointNumber":3051,"StopPointDesignation":"5","TimeTabledDateTime":"2017-12-18T20:10:45","ExpectedDateTime":"2017-12-18T20:11:03","JourneyNumber":30531,"Deviations":null}],"Trams":[],"Ships":[],"StopPointDeviations":[{"StopInfo":{"StopAreaNumber":1051,"StopAreaName":"T-Centralen","TransportMode":"METRO","GroupOfLine":"tunnelbanans gröna linje"},"Deviation":{"Text":"För din säkerhet, var uppmärksam på risken för ficktjuvar. *** For your safety, please be aware of pickpockets","Consequence":null,"ImportanceLevel":2}}]}}`)
	})

	realtime, err := client.Realtime.Search(context.Background(), &RealtimeSearchOptions{
		Key:    "XXXX",
		SiteID: "1002",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if realtime.Metros[0].GroupOfLine != "tunnelbanans blå linje" {
		t.Errorf("Expected 'tunnelbanans blå linje' got %s", realtime.Metros[0].GroupOfLine)
	}
}
