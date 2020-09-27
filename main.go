package main

import (
	"log"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalf("error: %v", err)
	}

}

func realMain() error {
	err := fleaflicker.GetLeagueScoreboard()
	if err != nil {
		return err
	}
	return nil
}
