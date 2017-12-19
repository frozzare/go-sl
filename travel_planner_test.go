package sl

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestTravelPlannerTrip(t *testing.T) {
	client, mux, _, teardown := setupClient()
	defer teardown()
	mux.HandleFunc("/TravelplannerV3/trip.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"Trip":[{"ServiceDays":[],"LegList":{"Leg":[{"Origin":{"name":"Slussen","type":"ST","id":"A=1@O=Slussen@X=18071491@Y=59319511@U=74@L=400102011@","extId":"400102011","lon":18.071491,"lat":59.319511,"prognosisType":"PROGNOSED","time":"22:58:00","date":"2017-12-18","track":"2","hasMainMast":true,"mainMastId":"A=1@O=Slussen (Stockholm)@X=18071860@Y=59320284@U=74@L=300109192@","mainMastExtId":"300109192"},"Destination":{"name":"T-Centralen","type":"ST","id":"A=1@O=T-Centralen@X=18061477@Y=59331358@U=74@L=400101051@","extId":"400101051","lon":18.061477,"lat":59.331358,"prognosisType":"PROGNOSED","time":"23:02:00","date":"2017-12-18","track":"3","hasMainMast":true,"mainMastId":"A=1@O=Sergels torg (Stockholm)@X=18064327@Y=59332563@U=74@L=300101000@","mainMastExtId":"300101000"},"JourneyDetailRef":{"ref":"1|4455|1|74|18122017"},"JourneyStatus":"P","Product":{"name":"TUNNELBANA  13","num":"20765","line":"13","catOut":"METRO   ","catIn":"MET","catCode":"1","catOutS":"MET","catOutL":"TUNNELBANA ","operatorCode":"SL","operator":"Storstockholms Lokaltrafik","admin":"101013"},"idx":"0","name":"TUNNELBANA  13","number":"20765","category":"MET","type":"JNY","reachable":true,"direction":"Ropsten"}]},"TariffResult":{"fareSetItem":[{"fareItem":[{"name":"Reskassa","desc":"Helt pris","price":3000,"cur":"SEK"},{"name":"Övriga försäljningsställen","desc":"Helt pris","price":4300,"cur":"SEK"},{"name":"Konduktör på Djurgårds- och Roslagsbanan","desc":"Helt pris","price":6000,"cur":"SEK"},{"name":"Reskassa","desc":"Reducerat pris","price":2000,"cur":"SEK"},{"name":"Övriga försäljningsställen","desc":"Reducerat pris","price":2900,"cur":"SEK"},{"name":"Konduktör på Djurgårds- och Roslagsbanan","desc":"Reducerat pris","price":4000,"cur":"SEK"}],"name":"ONEWAY","desc":"SL"}]},"idx":0,"tripId":"C-0","ctxRecon":"T$A=1@O=Slussen@L=400102011@a=128@$A=1@O=T-Centralen@L=400101051@a=128@$201712182258$201712182302$        $","duration":"PT4M","checksum":"A26A97EE_4"}],"serverVersion":"1.2","dialectVersion":"1.23","requestId":"xxx","scrB":"1|OB|MTµ11µ5698µ5698µ5703µ5703µ0µ0µ5µ5698µ1µ-2147483646µ0µ1µ2|PDHµc80e9bba6a4bc6bff038782eae38123c","scrF":"1|OF|MTµ11µ5713µ5713µ5718µ5718µ0µ0µ5µ5711µ5µ-2147483646µ0µ1µ2|PDHµc80e9bba6a4bc6bff038782eae38123c"}`)
	})

	trips, err := client.TravelPlanner.Trip(context.Background(), &TripOptions{
		Key:      "XXXX",
		DestID:   "1002",
		OriginID: "9192",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if trips[0].LegList.Leg[0].Name != "TUNNELBANA  13" {
		t.Errorf("Expected 'TUNNELBANA  13' got %s", trips[0].LegList.Leg[0].Name)
	}
}

func TestTravelPlannerJourney(t *testing.T) {
	client, mux, _, teardown := setupClient()
	defer teardown()
	mux.HandleFunc("/TravelplannerV3/journeydetail.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"Stops":{"Stop":[{"name":"Fruängen","id":"A=1@O=Fruängen@X=17964852@Y=59286754@U=74@L=400102851@","extId":"400102851","routeIdx":0,"lon":17.964852,"lat":59.286754,"depPrognosisType":"PROGNOSED","depTime":"08:04:00","depDate":"2017-12-19","depTrack":"1","hasMainMast":true,"mainMastId":"A=1@O=Fruängen (Stockholm)@X=17965454@Y=59285594@U=74@L=300109260@","mainMastExtId":"300109260"}]},"Names":{"Name":[{"Product":{"name":"TUNNELBANA  14","num":"20101","line":"14","catOut":"METRO   ","catIn":"MET","catCode":"1","catOutS":"MET","catOutL":"TUNNELBANA ","operatorCode":"SL","operator":"Storstockholms Lokaltrafik","admin":"107014"},"name":"TUNNELBANA  14","number":"20101","category":"MET","routeIdxFrom":0,"routeIdxTo":18}]},"Directions":{"Direction":[{"value":"Mörby centrum","routeIdxFrom":0,"routeIdxTo":18}]},"JourneyStatus":"P","ServiceDays":[{"sDaysR":"Mo - Do","sDaysI":"nicht 25. Dez 2017 bis 4. Jan 2018, 31. Jan","sDaysB":"780003C78F18"}],"lastPassRouteIdx":9,"lastPassStopRef":9,"ref":"1|5258|0|74|19122017"}`)
	})

	journey, err := client.TravelPlanner.Journey(context.Background(), &JourneyOptions{
		Key: "XXXX",
		ID:  "1|5258|0|74|19122017",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if journey.Stops.Stop[0].Name != "Fruängen" {
		t.Errorf("Expected 'Fruängen' got %s", journey.Stops.Stop[0].Name)
	}
}
