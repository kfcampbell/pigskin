package main

import (
	"fmt"
	"log"

	"github.com/kfcampbell/pigskin/clients/fleaflicker"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalf("error: %v", err)
	}

}

func realMain() error {
	res, err := fleaflicker.GetLeagueScoreboard()
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
