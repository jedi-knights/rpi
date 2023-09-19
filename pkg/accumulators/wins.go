package accumulators

import "github.com/jedi-knights/rpi/pkg"

// Wins is an accumulator that calculates the number of wins for a given team.
type Wins struct {
	SkipTeamName string
}

// NewWins returns a new Wins accumulator.
func NewWins(skipTeamName string) *Wins {
	return &Wins{
		SkipTeamName: skipTeamName,
	}
}

// Calculate returns the number of wins for a given team.
func (w *Wins) Calculate(teamName string, matches *[]pkg.Match) (int, error) {
	total := 0

	for _, match := range *matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(w.SkipTeamName) > 0 && match.Contains(w.SkipTeamName) {
			continue
		}

		if match.IsWinner(teamName) {
			total++
		}
	}

	return total, nil
}
