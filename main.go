package main

import (
	"fmt"
	"log"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
	"github.com/kfcampbell/pigskin/responses"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalf("error: %v", err)
	}

}

func realMain() error {
	scores, err := fleaflicker.GetLeagueScoreboard()
	if err != nil {
		return err
	}

	for i := 0; i < len(scores.Games); i++ {
		score := formatScore(scores.Games[i])
		fmt.Println(score)
	}

	return nil
}

func formatScore(game responses.FantasyGame) string {
	homeTeam := game.Home.Name
	awayTeam := game.Away.Name

	homeScore := game.HomeScore.Score.Value
	awayScore := game.AwayScore.Score.Value

	if homeScore > awayScore {
		return fmt.Sprintf("%v beat %v with a score of %v-%v.\n", homeTeam, awayTeam, homeScore, awayScore)
	} else if awayScore > homeScore {
		return fmt.Sprintf("%v beat %v with a score of %v-%v.\n", awayTeam, homeTeam, awayScore, homeScore)
	} else {
		return fmt.Sprintf("Whaaaat....%v and %v tied with a score of %v-%v.\n", homeTeam, awayTeam, homeScore, awayScore)
	}
}
