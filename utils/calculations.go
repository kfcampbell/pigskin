package utils

import (
	"fmt"

	"github.com/kfcampbell/pigskin/responses"
)

const topNScorers = 3

// GetBiggestWin returns the single biggest win in a slice of FantasyGames
func GetBiggestWin(games []responses.FantasyGame) (responses.FantasyGame, float32) {
	biggestWin := games[0]
	difference := GetDifferenceInScore(games[0])
	for i := 1; i < len(games); i++ {
		currDifference := GetDifferenceInScore(games[i])
		if currDifference > difference {
			difference = currDifference
			biggestWin = games[i]
		}
	}
	return biggestWin, difference
}

// GetClosestGame returns the closest game in a slice of FantasyGames
func GetClosestGame(games []responses.FantasyGame) (responses.FantasyGame, float32) {
	closestGame := games[0]
	difference := GetDifferenceInScore(games[0])
	for i := 1; i < len(games); i++ {
		currDifference := GetDifferenceInScore(games[i])
		if currDifference < difference {
			difference = currDifference
			closestGame = games[i]
		}
	}

	return closestGame, difference
}

// GetTopScorers returns the topNScorers in a slice of FantasyGames
func GetTopScorers(games []responses.FantasyGame) (*Scorers, error) {
	topScorers, err := NewScorers(true, topNScorers)
	if err != nil {
		return nil, fmt.Errorf("error getting top scorers: %v", err)
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

	return topScorers, nil
}

// GetBottomScorers returns the bottom "topNScorers" in a slice of FantasyGames
func GetBottomScorers(games []responses.FantasyGame) (*Scorers, error) {
	bottomScorers, err := NewScorers(false, topNScorers)
	if err != nil {
		return nil, fmt.Errorf("error getting top scorers: %v", err)
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
	return bottomScorers, nil
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
