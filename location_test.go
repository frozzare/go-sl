package sl

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestLocationSearch(t *testing.T) {
	client, mux, _, teardown := setupClient()
	defer teardown()
	mux.HandleFunc("/typeahead.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"StatusCode":0,"Message":null,"ExecutionTime":0,"ResponseData":[{"Name":"Södra station (på Rosenlundsg) (Stockholm)","SiteId":"1365","Type":"Station","X":"18057738","Y":"59312688"},{"Name":"Södra station (Stockholm)","SiteId":"9530","Type":"Station","X":"18061405","Y":"59313389"},{"Name":"Södra station(på Swedenborgsg) (Stockholm)","SiteId":"1339","Type":"Station","X":"18065064","Y":"59314099"},{"Name":"Södra Nånö (Norrtälje)","SiteId":"6603","Type":"Station","X":"18671099","Y":"59779381"},{"Name":"Södra Muskö (Haninge)","SiteId":"8623","Type":"Station","X":"18061783","Y":"58969821"},{"Name":"Södra grinden (Solna)","SiteId":"3428","Type":"Station","X":"17998912","Y":"59393105"},{"Name":"Södra Träskö (Värmdö)","SiteId":"174","Type":"Station","X":"18790943","Y":"59382867"},{"Name":"Södra Grinda (Värmdö)","SiteId":"147","Type":"Station","X":"18552900","Y":"59407254"},{"Name":"Södra Långgatan (Solna)","SiteId":"3451","Type":"Station","X":"18006904","Y":"59363378"},{"Name":"Södra Evlinge (Värmdö)","SiteId":"4353","Type":"Station","X":"18493786","Y":"59256757"}]}`)
	})

	locations, err := client.Location.Search(context.Background(), &LocationSearchOptions{
		Key:          "XXXX",
		SearchString: "Södra",
	})

	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if locations[0].Name != "Södra station (på Rosenlundsg) (Stockholm)" {
		t.Errorf("Expected 'Södra station (på Rosenlundsg) (Stockholm)' got %s", locations[0].Name)
	}
}
