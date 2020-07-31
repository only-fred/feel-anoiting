package api

import (
	"encoding/json"
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
