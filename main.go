package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
	"github.com/kfcampbell/pigskin/responses"
)

const topNScorers = 3

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

	biggestWin := getBiggestWin(scores.Games)
	fmt.Println(biggestWin)

	topScorers, err := getTopScorers(scores.Games)
	if err != nil {
		return err
	}
	fmt.Println(topScorers)

	bottomScorers, err := getBottomScorers(scores.Games)
	if err != nil {
		return err
	}
	fmt.Println(bottomScorers)

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

func getBiggestWin(games []responses.FantasyGame) string {
	biggestWin := games[0]
	difference := getDifferenceInScore(games[0])
	for i := 1; i < len(games); i++ {
		currDifference := getDifferenceInScore(games[i])
		if currDifference > difference {
			difference = currDifference
			biggestWin = games[i]
		}
	}
	if isHomeWinning(biggestWin) > 0 {
		return fmt.Sprintf("Biggest win: %v over %v, %v-%v (difference of %v points).",
			biggestWin.Home.Name, biggestWin.Away.Name, biggestWin.HomeScore.Score.Value, biggestWin.AwayScore.Score.Value, difference)
	}
	return fmt.Sprintf("Biggest win: %v over %v, %v-%v (difference of %v points).",
		biggestWin.Away.Name, biggestWin.Home.Name, biggestWin.AwayScore.Score.Value, biggestWin.HomeScore.Score.Value, difference)
}

func getTopScorers(games []responses.FantasyGame) (string, error) {
	topScorers, err := NewScorers(true, topNScorers)
	if err != nil {
		return "Error getting top scorers", fmt.Errorf("error getting top scorers: %v", err)
	}

	for i := 0; i < len(games); i++ {
		homeScorer := GetScorer(true, games[i])
		awayScorer := GetScorer(false, games[i])

		if topScorers.ShouldAddScorer(*homeScorer) {
			topScorers.AddScorer(*homeScorer)
		}
		if topScorers.ShouldAddScorer(*awayScorer) {
			topScorers.AddScorer(*awayScorer)
		}
	}

	formatted := "Top scorers: "
	for i := 0; i < topScorers.length; i++ {
		formatted += fmt.Sprintf("%v (%v points), ", topScorers.scorers[i].Team.Name, topScorers.scorers[i].Score)
	}
	formatted = strings.TrimSuffix(formatted, ", ")

	return formatted, nil
}

func getBottomScorers(games []responses.FantasyGame) (string, error) {
	bottomScorers, err := NewScorers(false, topNScorers)
	if err != nil {
		return "Error getting top scorers", fmt.Errorf("error getting top scorers: %v", err)
	}

	for i := 0; i < len(games); i++ {
		homeScorer := GetScorer(true, games[i])
		awayScorer := GetScorer(false, games[i])

		if bottomScorers.ShouldAddScorer(*homeScorer) {
			bottomScorers.AddScorer(*homeScorer)
		}
		if bottomScorers.ShouldAddScorer(*awayScorer) {
			bottomScorers.AddScorer(*awayScorer)
		}
	}

	formatted := "Bottom scorers: "
	for i := 0; i < bottomScorers.length; i++ {
		formatted += fmt.Sprintf("%v (%v points), ", bottomScorers.scorers[i].Team.Name, bottomScorers.scorers[i].Score)
	}
	formatted = strings.TrimSuffix(formatted, ", ")

	return formatted, nil
}

// isHomeWinning returns 1 if true, -1 if false, and 0 if tied
func isHomeWinning(game responses.FantasyGame) int {
	homeScore := game.HomeScore.Score.Value
	awayScore := game.AwayScore.Score.Value

	if homeScore > awayScore {
		return 1
	} else if awayScore > homeScore {
		return -1
	} else {
		return 0
	}
}

func getDifferenceInScore(game responses.FantasyGame) float32 {
	if isHomeWinning(game) > 0 {
		return game.HomeScore.Score.Value - game.AwayScore.Score.Value
	}
	return game.AwayScore.Score.Value - game.HomeScore.Score.Value
}
