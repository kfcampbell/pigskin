package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
	"github.com/kfcampbell/pigskin/clients/groupme"
	"github.com/kfcampbell/pigskin/responses"
	"github.com/kfcampbell/pigskin/utils"
)

const topNScorers = 3

func main() {
	if err := realMain(); err != nil {
		log.Fatalf("error: %v", err)
	}

}

func realMain() error {
	fleaflickerLeagueID := os.Getenv("LEAGUE_ID")
	groupmeChatID := os.Getenv("GROUPME_CHAT_ID")
	groupmeAPIKey := os.Getenv("GROUPME_API_KEY")

	if fleaflickerLeagueID == "" || groupmeChatID == "" || groupmeAPIKey == "" {
		return fmt.Errorf("missing either LEAGUE_ID or GROUPME_CHAT_ID or GROUPME_API_KEY")
	}

	scores, err := fleaflicker.GetLeagueScoreboard(fleaflickerLeagueID)
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

	body := fmt.Sprintf("%v\n%v\n%v\n", biggestWin, topScorers, bottomScorers)
	fmt.Printf("body message: \n%v", body)

	err = groupme.PostMessage(body, groupmeChatID, groupmeAPIKey)
	if err != nil {
		return err
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
		return fmt.Sprintf("- **Biggest win**: _%v_ over _%v_, %v-%v (difference of %v points).",
			biggestWin.Home.Name, biggestWin.Away.Name, biggestWin.HomeScore.Score.Value, biggestWin.AwayScore.Score.Value, difference)
	}
	return fmt.Sprintf("- **Biggest win**: _%v_ over _%v_, %v-%v (difference of %v points).",
		biggestWin.Away.Name, biggestWin.Home.Name, biggestWin.AwayScore.Score.Value, biggestWin.HomeScore.Score.Value, difference)
}

func getTopScorers(games []responses.FantasyGame) (string, error) {
	topScorers, err := utils.NewScorers(true, topNScorers)
	if err != nil {
		return "Error getting top scorers", fmt.Errorf("error getting top scorers: %v", err)
	}

	for i := 0; i < len(games); i++ {
		homeScorer := utils.GetScorer(true, games[i])
		awayScorer := utils.GetScorer(false, games[i])

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

func getBottomScorers(games []responses.FantasyGame) (string, error) {
	bottomScorers, err := utils.NewScorers(false, topNScorers)
	if err != nil {
		return "Error getting top scorers", fmt.Errorf("error getting top scorers: %v", err)
	}

	for i := 0; i < len(games); i++ {
		homeScorer := utils.GetScorer(true, games[i])
		awayScorer := utils.GetScorer(false, games[i])

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
