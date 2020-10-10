package utils

import (
	"fmt"
	"strings"

	"github.com/kfcampbell/pigskin/responses"
)

const topNScorers = 3

// FormatScore prettyprints the score with the following format:
// %v beat %v with a score of %v-%v
func FormatScore(game responses.FantasyGame) string {
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

// GetBiggestWin returns the single biggest win in a slice of FantasyGames
func GetBiggestWin(games []responses.FantasyGame) string {
	biggestWin := games[0]
	difference := GetDifferenceInScore(games[0])
	for i := 1; i < len(games); i++ {
		currDifference := GetDifferenceInScore(games[i])
		if currDifference > difference {
			difference = currDifference
			biggestWin = games[i]
		}
	}
	if IsHomeWinning(biggestWin) > 0 {
		return fmt.Sprintf("- **Biggest win**: _%v_ over _%v_, %v-%v (difference of %v points).",
			biggestWin.Home.Name, biggestWin.Away.Name, biggestWin.HomeScore.Score.Value, biggestWin.AwayScore.Score.Value, difference)
	}
	return fmt.Sprintf("- **Biggest win**: _%v_ over _%v_, %v-%v (difference of %v points).",
		biggestWin.Away.Name, biggestWin.Home.Name, biggestWin.AwayScore.Score.Value, biggestWin.HomeScore.Score.Value, difference)
}

// GetTopScorers returns the topNScorers in a slice of FantasyGames
func GetTopScorers(games []responses.FantasyGame) (string, error) {
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

	formatted := "- **Top scorers** :"
	for i := 0; i < topScorers.GetLength(); i++ {
		formatted += fmt.Sprintf("_%v_ (%v points), ", topScorers.GetScorers()[i].Team.Name, topScorers.GetScorers()[i].Score)
	}
	formatted = strings.TrimSuffix(formatted, ", ")

	return formatted, nil
}

// GetBottomScorers returns the bottom "topNScorers" in a slice of FantasyGames
func GetBottomScorers(games []responses.FantasyGame) (string, error) {
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

	formatted := "- **Bottom scorers**: "
	for i := 0; i < bottomScorers.GetLength(); i++ {
		formatted += fmt.Sprintf("_%v_ (%v points), ", bottomScorers.GetScorers()[i].Team.Name, bottomScorers.GetScorers()[i].Score)
	}
	formatted = strings.TrimSuffix(formatted, ", ")

	return formatted, nil
}

// IsHomeWinning returns 1 if true, -1 if false, and 0 if tied
func IsHomeWinning(game responses.FantasyGame) int {
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

// GetDifferenceInScore returns the absolute value of the difference between
// the home and away teams
func GetDifferenceInScore(game responses.FantasyGame) float32 {
	if IsHomeWinning(game) > 0 {
		return game.HomeScore.Score.Value - game.AwayScore.Score.Value
	}
	return game.AwayScore.Score.Value - game.HomeScore.Score.Value
}
