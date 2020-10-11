package fleaflicker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kfcampbell/pigskin/responses"
)

const apiRoot = "https://www.fleaflicker.com/api/"

// GetLeagueStandings gets the league of
func GetLeagueStandings(leagueID string) error {
	url := apiRoot + "FetchLeagueStandings" + getFiltering(leagueID)
	result, err := http.Get(url)
	if err != nil {
		return err
	}

	fmt.Printf("standings: %v\n", result)
	return nil
}

// GetLeagueScoreboard returns the league scoreboard
func GetLeagueScoreboard(leagueID string) (*responses.LeagueScoreboard, error) {
	response := &responses.LeagueScoreboard{}
	url := apiRoot + "FetchLeagueScoreboard" + getFiltering(leagueID)
	res, err := http.Get(url)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	defer res.Body.Close()

	if err != nil {
		return response, err
	}

	return response, nil
}

// GetLeagueBoxscore returns a boxscore from a given game
func GetLeagueBoxscore(leagueID string, fantasyGameID string) (*responses.LeagueBoxscore, error) {
	response := &responses.LeagueBoxscore{}
	url := apiRoot + "FetchLeagueBoxscore" + getBoxScoreFiltering(leagueID, fantasyGameID)
	res, err := http.Get(url)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	defer res.Body.Close()

	if err != nil {
		return response, err
	}
	return response, nil
}

func getBoxScoreFiltering(leagueID string, fantasyGameID string) string {
	return "?sport=NFL&league_id=" + fmt.Sprintf("%v", leagueID) + "&fantasy_game_id=" + fantasyGameID + "&scoring_period=0"
}

func getFiltering(leagueID string) string {
	return "?sport=NFL&season=" + fmt.Sprintf("%v", time.Now().Year()) + "&league_id=" + fmt.Sprintf("%v", leagueID)
}
