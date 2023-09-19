package accumulators

import (
	"github.com/jedi-knights/rpi/pkg/match"
)

// Ties is an accumulator that calculates the number of draws for a given team.
type Ties struct {
	SkipTeamName string
}

// NewTies returns a new Ties accumulator.
func NewTies(skipTeamName string) *Ties {
	return &Ties{
		SkipTeamName: skipTeamName,
	}
}

// Calculate returns the number of ties for a given team.
func (w *Ties) Calculate(teamName string, matches *[]match.Match) (int, error) {
	total := 0

	for _, match := range *matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(w.SkipTeamName) > 0 && match.Contains(w.SkipTeamName) {
			continue
		}

		if match.IsDraw() {
			total++
		}
	}

	return total, nil
}
