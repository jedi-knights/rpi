package accumulators

import "github.com/jedi-knights/rpi/pkg"

// Draws is an accumulator that calculates the number of draws for a given team.
type Draws struct {
	skipTeamName string
	matches      *[]pkg.Match
}

// NewDraws returns a new Draws accumulator.
func NewDraws(skipTeamName string, matches *[]pkg.Match) *Draws {
	return &Draws{
		skipTeamName: skipTeamName,
		matches:      matches,
	}
}

// Calculate returns the number of draws for a given team.
func (w *Draws) Calculate(teamName string) (int, error) {
	total := 0

	for _, match := range *w.matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(w.skipTeamName) > 0 && match.Contains(w.skipTeamName) {
			continue
		}

		if match.IsDraw() {
			total++
		}
	}

	return total, nil
}
