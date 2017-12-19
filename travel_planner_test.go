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

func TestTravelPlannerReconstruction(t *testing.T) {
	client, mux, _, teardown := setupClient()
	defer teardown()
	mux.HandleFunc("/TravelplannerV3/reconstruction.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"Trip":[{"ServiceDays":[{"planningPeriodBegin":"2017-12-17","planningPeriodEnd":"2018-01-31","sDaysR":"Mo - Fr","sDaysI":"nicht 25. Dez 2017 bis 5. Jan 2018, 31. Jan","sDaysB":"7C0003E7CF98"}],"LegList":{"Leg":[{"Origin":{"name":"Centralen (Klarabergsviad.)","type":"ST","id":"A=1@O=Centralen (Klarabergsviad.)@X=18057369@Y=59330873@U=74@L=400110537@","extId":"400110537","lon":18.057369,"lat":59.330873,"prognosisType":"PROGNOSED","time":"09:11:00","date":"2017-12-19","track":"R","rtTime":"09:12:00","rtDate":"2017-12-19","hasMainMast":true,"mainMastId":"A=1@O=Centralen (Stockholm)@X=18057657@Y=59331134@U=74@L=300101002@","mainMastExtId":"300101002"},"Destination":{"name":"Sergels torg","type":"ST","id":"A=1@O=Sergels torg@X=18062790@Y=59332985@U=74@L=400110307@","extId":"400110307","lon":18.06279,"lat":59.332985,"prognosisType":"PROGNOSED","time":"09:13:00","date":"2017-12-19","track":"M","rtTime":"09:13:00","rtDate":"2017-12-19","hasMainMast":true,"mainMastId":"A=1@O=Sergels torg (Stockholm)@X=18064327@Y=59332563@U=74@L=300101000@","mainMastExtId":"300101000"},"JourneyDetailRef":{"ref":"1|12526|0|74|19122017"},"JourneyStatus":"P","Product":{"name":"BUSS  54","num":"28891","line":"54","catOut":"BUS     ","catIn":"BUS","catCode":"3","catOutS":"BUS","catOutL":"BUSS ","operatorCode":"SL","operator":"Storstockholms Lokaltrafik","admin":"100054"},"idx":"0","name":"BUSS  54","number":"28891","category":"BUS","type":"JNY","reachable":true,"direction":"Storängsbotten"},{"Origin":{"name":"Sergels torg","type":"ST","id":"A=1@O=Sergels torg@X=18062790@Y=59332985@U=74@L=400110307@","extId":"400110307","lon":18.06279,"lat":59.332985,"time":"09:15:00","date":"2017-12-19","hasMainMast":true,"mainMastId":"A=1@O=Sergels torg (Stockholm)@X=18064327@Y=59332563@U=74@L=300101000@","mainMastExtId":"300101000"},"Destination":{"name":"Hötorget","type":"ST","id":"A=1@O=Hötorget@X=18062960@Y=59335610@U=74@L=400101111@","extId":"400101111","lon":18.06296,"lat":59.33561,"time":"09:20:00","date":"2017-12-19","hasMainMast":true,"mainMastId":"A=1@O=Kungsgatan / Sveavägen (Stockholm)@X=18063868@Y=59335529@U=74@L=300101019@","mainMastExtId":"300101019"},"GisRef":{"ref":"G|1|G@F|A=1@O=Sergels torg@X=18062790@Y=59332985@U=74@L=400110307@|A=1@O=Hötorget@X=18062960@Y=59335610@U=74@L=400101111@|19122017|91500|92000|ft|ft@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$§bt@0@2000@120@-1@100@1@1000@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$§tt@0@5000@120@-1@100@1@2500@0@@@@@false@0@-1@$f@$f@$f@$f@$f@$§|"},"idx":"1","name":"","type":"WALK","duration":"PT5M","dist":398}]},"TariffResult":{"fareSetItem":[{"fareItem":[{"name":"Reskassa","desc":"Helt pris","price":3000,"cur":"SEK"},{"name":"Övriga försäljningsställen","desc":"Helt pris","price":4300,"cur":"SEK"},{"name":"Konduktör på Djurgårds- och Roslagsbanan","desc":"Helt pris","price":6000,"cur":"SEK"},{"name":"Reskassa","desc":"Reducerat pris","price":2000,"cur":"SEK"},{"name":"Övriga försäljningsställen","desc":"Reducerat pris","price":2900,"cur":"SEK"},{"name":"Konduktör på Djurgårds- och Roslagsbanan","desc":"Reducerat pris","price":4000,"cur":"SEK"}],"name":"ONEWAY","desc":"SL"}]},"idx":0,"tripId":"C-0","ctxRecon":"T$A=1@O=Centralen (Klarabergsviad.)@L=400110537@a=128@$A=1@O=Sergels torg@L=400110307@a=128@$201712190911$201712190913$        $§G@F$A=1@O=Sergels torg@L=400110307@a=128@$A=1@O=Hötorget@L=400101111@a=128@$201712190915$201712190920$$","duration":"PT9M","checksum":"076D4066_4"}]}`)
	})

	trip, err := client.TravelPlanner.Reconstruction(context.Background(), &ReconstructionOptions{
		Key: "XXXX",
		Ctx: "T$A=1@O=Centralen%20(Klarabergsviad.)@L=400110537@a=128@$A=1@O=Sergels%20torg@L=400110307@a=128@$201712190911$201712190913$%20$§G@F$A=1@O=Sergels%20torg@L=400110307@a=128@$A=1@O=Hötorget%20(Stockholm)@L=300109119@a=128@$201712190915$201712190921$$",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if trip.LegList.Leg[0].Name != "BUSS  54" {
		t.Errorf("Expected 'BUSS  54' got %s", trip.LegList.Leg[0].Name)
	}
}
