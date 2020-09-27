package fleaflicker

import (
	"fmt"
	"net/http"
	"time"
)

const apiRoot = "https://www.fleaflicker.com/api/"
const leagueID = 170474

// GetLeagueStandings gets the league of
func GetLeagueStandings() error {
	url := apiRoot + "FetchLeagueStandings" + getFiltering()
	result, err := http.Get(url)
	if err != nil {
		return err
	}

	fmt.Printf("result: %v\n", result)
	return nil
}

// GetLeagueScoreboard returns the league scoreboard
func GetLeagueScoreboard() error {
	url := apiRoot + "FetchLeagueScoreboard" + getFiltering()
	fmt.Printf("url: %v\n", url)
	result, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Printf("result: %v", result)
	return nil
}

func getFiltering() string {
	return "?sport=NFL&season=" + fmt.Sprintf("%v", time.Now().Year()) + "&league_id=" + fmt.Sprintf("%v", leagueID)
}
