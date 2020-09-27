package fleaflicker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kfcampbell/pigskin/responses"
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
func GetLeagueScoreboard() (*responses.LeagueScoreboard, error) {
	response := &responses.LeagueScoreboard{}
	url := apiRoot + "FetchLeagueScoreboard" + getFiltering()
	fmt.Printf("url: %v\n", url)
	res, err := http.Get(url)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	defer res.Body.Close()

	return response, nil
}

func getFiltering() string {
	return "?sport=NFL&season=" + fmt.Sprintf("%v", time.Now().Year()) + "&league_id=" + fmt.Sprintf("%v", leagueID)
}
