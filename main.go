package main

import (
	"fmt"
	"strings"

	api "./connection/api"
	_ "./connection/db"
)

func main() {

	loopAnswer := "YES"
	for loopAnswer == "YES" {
		var (
			userAnswer int
		)

		fmt.Print("|| [1] View all || [2] Search by ID || [3] Search by country || [0] Cancel\n->")
		fmt.Scanf("%d", &userAnswer)

		respObj := api.ConsumingAPI()

		switch userAnswer {
		case 1:
			for i := 0; i < len(respObj.Features); i++ {
				api.ShowAttributes(respObj, i)
			}
			break

		case 2:
			var searchByID int

			fmt.Print("Type the ID: \n->")
			fmt.Scanf("%d", &searchByID)

			for i := 0; i < len(respObj.Features); i++ {
				if searchByID == respObj.Features[i].Attributes.OBJECTID {
					api.ShowAttributes(respObj, i)
				}
			}
			break

		case 3:
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

	// conn.Connection()
}
