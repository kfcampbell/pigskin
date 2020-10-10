package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
	"github.com/kfcampbell/pigskin/utils"
)

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
		score := utils.FormatScore(scores.Games[i])
		fmt.Println(score)
	}

	biggestWin := utils.GetBiggestWin(scores.Games)
	fmt.Println(biggestWin)

	topScorers, err := utils.GetTopScorers(scores.Games)
	if err != nil {
		return err
	}
	fmt.Println(topScorers)

	bottomScorers, err := utils.GetBottomScorers(scores.Games)
	if err != nil {
		return err
	}
	fmt.Println(bottomScorers)

	body := fmt.Sprintf("%v\n%v\n%v\n", biggestWin, topScorers, bottomScorers)
	fmt.Printf("body message: \n%v", body)

	//err = groupme.PostMessage(body, groupmeChatID, groupmeAPIKey)
	if err != nil {
		return err
	}

	return nil
}
