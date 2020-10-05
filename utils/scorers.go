package utils

import (
	"fmt"

	"github.com/kfcampbell/pigskin/responses"
)

// Scorer is a convenience type for keeping track of a week's high- and low-scorers
type Scorer struct {
	Team  responses.Team
	Score float32
}

// GetScorer returns a home or away scorer from a given game
func GetScorer(home bool, game responses.FantasyGame) *Scorer {
	if home {
		return &Scorer{
			Team:  game.Home,
			Score: game.HomeScore.Score.Value,
		}
	}

	return &Scorer{
		Team:  game.Away,
		Score: game.AwayScore.Score.Value,
	}
}

// Scorers provides a way to track and order the top or bottom N scorers.
// No order guarantees are made about the Scorers slice
type Scorers struct {
	isTop   bool
	length  int
	scorers []Scorer
}

// NewScorers creates and returns a new Scorers struct
func NewScorers(isTop bool, length int) (*Scorers, error) {
	if length < 1 {
		return &Scorers{}, fmt.Errorf("cannot initialize Scorers with a length of %v", length)
	}
	scorers := make([]Scorer, length)

	if !isTop {
		for i := 0; i < len(scorers); i++ {
			scorers[i].Score = 1000
		}
	}
	return &Scorers{
		isTop:   isTop,
		length:  length,
		scorers: scorers,
	}, nil
}

// ShouldAddScorer returns true if the given scorer is in the top (or bottom) N scorers
func (s *Scorers) ShouldAddScorer(scorer Scorer) bool {
	if len(s.scorers) < s.length {
		return true
	}

	if s.isTop {
		for i := 0; i < len(s.scorers); i++ {
			if s.scorers[i].Score < scorer.Score {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(s.scorers); i++ {
		if s.scorers[i].Score > scorer.Score {
			return true
		}
	}
	return false
}

// AddScorer adds a given scorer to the slice of N scorers, replacing the previous high- or lowest-rated scorer
func (s *Scorers) AddScorer(scorer Scorer) {
	if s.isTop {
		lowestScoreIndex := 0
		lowestScore := float32(1000)
		for i := 0; i < len(s.scorers); i++ {
			if s.scorers[i].Score < lowestScore {
				lowestScore = s.scorers[i].Score
				lowestScoreIndex = i
			}
		}
		s.scorers[lowestScoreIndex] = scorer
		return
	}
	highestScoreIndex := 0
	highestScore := float32(-1000)
	for i := 0; i < len(s.scorers); i++ {
		if s.scorers[i].Score > highestScore {
			highestScore = s.scorers[i].Score
			highestScoreIndex = i
		}
	}
	s.scorers[highestScoreIndex] = scorer
}

// GetLength returns the length of a Scorers object
func (s *Scorers) GetLength() int {
	return s.length
}

// GetScorers returns the slice of Scorers
func (s *Scorers) GetScorers() []Scorer {
	return s.scorers
}
