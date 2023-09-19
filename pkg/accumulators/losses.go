package accumulators

import "github.com/jedi-knights/rpi/pkg"

// Losses is an accumulator that calculates the number of losses for a given team.
type Losses struct {
	SkipTeamName string
}

// NewLosses returns a new Losses accumulator.
func NewLosses(skipTeamName string) *Losses {
	return &Losses{
		SkipTeamName: skipTeamName,
	}
}

// Calculate returns the number of losses for a given team.
func (w *Losses) Calculate(teamName string, matches *[]pkg.Match) (int, error) {
	total := 0

	for _, match := range *matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(w.SkipTeamName) > 0 && match.Contains(w.SkipTeamName) {
			continue
		}

		if match.IsLoser(teamName) {
			total++
		}
	}

	return total, nil
}
