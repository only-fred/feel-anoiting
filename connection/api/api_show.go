package api

import "fmt"

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
func ShowSumAttributes(respObj Response) {
	var (
		sumConfirmed int
		sumRecovered int
		sumDeaths    int
		sumActive    int
	)
	for i := 0; i < len(respObj.Features); i++ {
		sumConfirmed = sumConfirmed + respObj.Features[i].Attributes.Confirmed
		sumRecovered = sumRecovered + respObj.Features[i].Attributes.Recovered
		sumDeaths = sumDeaths + respObj.Features[i].Attributes.Deaths
		sumActive = sumActive + respObj.Features[i].Attributes.Active
	}

	fmt.Printf("\n> Confirmed: %d\n", sumConfirmed)
	fmt.Printf("> Recovered: %d\n", sumRecovered)
	fmt.Printf("> Deaths: %d\n", sumDeaths)
	fmt.Printf("> Active: %d\n", sumActive)
}
