package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	arcgisAPI = "https://services1.arcgis.com/0MSEUqKaxRlEPj5g/ArcGIS/rest/services/ncov_cases/FeatureServer/1/query?where=1%3D1&outFields=*&f=pjson"
)

type Response struct {
	Features []Features `json:"features"`
}
type Features struct {
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	OBJECTID      int    `json:"OBJECTID"`
	ProvinceState string `json:"Province_State"`
	CountryRegion string `json:"Country_Region"`
	Confirmed     int    `json:"Confirmed"`
	Recovered     int    `json:"Recovered"`
	Deaths        int    `json:"Deaths"`
	Active        int    `json:"Active"`
}

func ConsumingAPI() Response {
	resp := responseArcgisAPI()

	respData := readDataFromResponse(resp)

	var respObj Response
	json.Unmarshal(respData, &respObj)

	return respObj
}
func responseArcgisAPI() *http.Response {
	resp, err := http.Get(arcgisAPI)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
func readDataFromResponse(resp *http.Response) []byte {
	respData, err := ioutil.ReadAll(responseArcgisAPI().Body)
	if err != nil {
		log.Fatal(err)
	}

	return respData
}

func ShowAttributes(respObj Response, i int) {
	fmt.Printf("\n> ID: %d\n", respObj.Features[i].Attributes.OBJECTID)
	fmt.Printf("\n> Country: %s\n", respObj.Features[i].Attributes.CountryRegion)
	fmt.Printf("> State: %s\n", respObj.Features[i].Attributes.ProvinceState)
	fmt.Printf("> Confirmed: %d\n", respObj.Features[i].Attributes.Confirmed)
	fmt.Printf("> Recovered: %d\n", respObj.Features[i].Attributes.Recovered)
	fmt.Printf("> Deaths: %d\n", respObj.Features[i].Attributes.Deaths)
	fmt.Printf("> Active: %d\n", respObj.Features[i].Attributes.Active)
}
func ShowCountryAttributes(respObj Response, searchByCountry string) {
	var (
		sum int
	)

	for i := 0; i < len(respObj.Features); i++ {
		if searchByCountry == respObj.Features[i].Attributes.CountryRegion {
			ShowAttributes(respObj, i)
			sum = sum + 1
		}
	}
	fmt.Printf("\nTotal: %d", sum)
}

func ShowStateAttributes(respObj Response, searchState string) {
	for i := 0; i < len(respObj.Features); i++ {
		if searchState == respObj.Features[i].Attributes.ProvinceState {
			ShowAttributes(respObj, i)
		}
	}
}
