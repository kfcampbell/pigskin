package utils

import (
	"fmt"

	"github.com/kfcampbell/pigskin/responses"
)

// FormatScore prettyprints the score with the following format:
// %v beat %v with a score of %v-%v
func FormatScore(game responses.FantasyGame) string {
	homeTeam := game.Home.Name
	awayTeam := game.Away.Name

	homeScore := game.HomeScore.Score.Value
	awayScore := game.AwayScore.Score.Value

	if homeScore > awayScore {
		return fmt.Sprintf("%v beat %v with a score of %.2f-%.2f.\n", homeTeam, awayTeam, homeScore, awayScore)
	} else if awayScore > homeScore {
		return fmt.Sprintf("%v beat %v with a score of %.2f-%.2f.\n", awayTeam, homeTeam, awayScore, homeScore)
	} else {
		return fmt.Sprintf("Whaaaat....%v and %v tied with a score of %.2f-%.2f.\n", homeTeam, awayTeam, homeScore, awayScore)
	}
}

// FormatBiggestWin prettyprints the winner of a fantasy game
func FormatBiggestWin(biggestWin responses.FantasyGame, difference float32) string {
	if IsHomeWinning(biggestWin) > -1 {
		return fmt.Sprintf("- **Biggest win**: \n_%v_ over _%v_, \n%.2f-%.2f (difference of %.2f points).",
			biggestWin.Home.Name, biggestWin.Away.Name, biggestWin.HomeScore.Score.Value, biggestWin.AwayScore.Score.Value, difference)
	}
	return fmt.Sprintf("- **Biggest win**: \n_%v_ over _%v_, \n%.2f-%.2f (difference of %.2f points).",
		biggestWin.Away.Name, biggestWin.Home.Name, biggestWin.AwayScore.Score.Value, biggestWin.HomeScore.Score.Value, difference)
}

// FormatClosestGame prettyprints the closest game
func FormatClosestGame(closestGame responses.FantasyGame, difference float32) string {
	if IsHomeWinning(closestGame) > 0 {
		return fmt.Sprintf("- **Closest game**: \n_%v_ over %v, \n%.2f-%.2f (difference of %.2f points).", closestGame.Home.Name, closestGame.Away.Name, closestGame.HomeScore.Score.Value, closestGame.AwayScore.Score.Value, difference)
	}

	return fmt.Sprintf("- **Closest game**: \n_%v_ over %v, \n%.2f-%.2f (difference of %.2f points).", closestGame.Away.Name, closestGame.Home.Name, closestGame.AwayScore.Score.Value, closestGame.HomeScore.Score.Value, difference)
}

// FormatTopScorers prettyprints a given *Scorers object
func FormatTopScorers(topScorers *Scorers) string {
	formatted := "- **Top scorers** :\n"
	for i := 0; i < topScorers.GetLength(); i++ {
		formatted += fmt.Sprintf("_%v_ (%.2f points)\n", topScorers.GetScorers()[i].Team.Name, topScorers.GetScorers()[i].Score)
	}

	return formatted
}

// FormatBottomScorers prettyprints the bottom scorers
func FormatBottomScorers(bottomScorers *Scorers) string {
	formatted := "- **Bottom scorers**: \n"
	for i := 0; i < bottomScorers.GetLength(); i++ {
		formatted += fmt.Sprintf("_%v_ (%.2f points)\n", bottomScorers.GetScorers()[i].Team.Name, bottomScorers.GetScorers()[i].Score)
	}

	return formatted
}
