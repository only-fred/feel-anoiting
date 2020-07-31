package main

import (
	"fmt"
	"strings"
	"time"

	api "./connection/api"
	control "./control"
)

func main() {

	loopAnswer := "YES"
	for loopAnswer == "YES" {
		var (
			userAnswer int
		)

		fmt.Print("[1] View all \n[2] World \n[3] Search by ID \n[4] Search by country \n\n[5] Update and Read Database \n[6] Compare datas\n\n[0] Cancel\n->")
		fmt.Scanf("%d", &userAnswer)

		respObj := api.ConsumingAPI()

		switch userAnswer {
		case 1:
			for i := 0; i < len(respObj.Features); i++ {
				api.ShowAttributes(respObj, i)
			}
			fmt.Printf("\nTotal: %d", len(respObj.Features))
			break

		case 2:
			api.ShowSumAttributes(respObj)
			break

		case 3:
			var searchByID int

			fmt.Print("Type the ID: \n->")
			fmt.Scanf("%d", &searchByID)

			for i := 0; i < len(respObj.Features); i++ {
				if searchByID == respObj.Features[i].Attributes.OBJECTID {
					api.ShowAttributes(respObj, i)
				}
			}
			break

		case 4:
			var (
				userAnswer string

				searchByCountry string
				searchState     string
			)

			// searchByCountry := bufio.NewScanner(os.Stdin)
			// searchState := bufio.NewScanner(os.Stdin)

			fmt.Print("Country: (ex.: Brazil)\n->")
			// searchByCountry.Scan()
			fmt.Scanf("%s", &searchByCountry)
			searchByCountry = strings.Title(searchByCountry)

			fmt.Print("Search a specific state? (YES/NO)\n->")
			fmt.Scanf("%s", &userAnswer)
			userAnswer = strings.ToUpper(userAnswer)

			if userAnswer != "YES" {
				api.ShowCountryAttributes(respObj, searchByCountry)
			} else {
				fmt.Print("State (ex.: Ceara)\n->")
				// searchState.Scan()
				fmt.Scanf("%s", &searchState)
				searchState = strings.Title(searchState)

				api.ShowStateAttributes(respObj, searchState)
			}
			break

		case 5:
			timeNow := timeNow()
			lastUpdate := control.Read()

			if timeNow != lastUpdate {
				for i := 0; i < len(respObj.Features); i++ {
					id := respObj.Features[i].Attributes.OBJECTID
					state := respObj.Features[i].Attributes.ProvinceState
					country := respObj.Features[i].Attributes.CountryRegion
					confirmed := respObj.Features[i].Attributes.Confirmed
					recovered := respObj.Features[i].Attributes.Recovered
					deaths := respObj.Features[i].Attributes.Deaths
					active := respObj.Features[i].Attributes.Active

					control.Update(state, country, timeNow, confirmed, recovered, deaths, active, id)
				}
				fmt.Print("Successfully updated")

			} else {
				fmt.Print("Nothing to update")
			}

			control.ReadAll()
			break

		case 0:
			break

		default:
			userAnswer = 99
		}

		if userAnswer != 99 {
			fmt.Print("\nTry again? (YES/NO)\n->")
			fmt.Scanf("%s", &loopAnswer)
			loopAnswer = strings.ToUpper(loopAnswer)

		}
	}
}

func timeNow() string {
	const layout = "02/01/2006"
	timeNow := time.Now()

	realTimeNow := timeNow.Format(layout) //Convert to string, IDK D:

	return realTimeNow
}
