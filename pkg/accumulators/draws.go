package accumulators

import "github.com/jedi-knights/rpi/pkg"

// Draws is an accumulator that calculates the number of draws for a given team.
type Draws struct {
	SkipTeamName string
}

// NewDraws returns a new Draws accumulator.
func NewDraws(skipTeamName string) *Draws {
	return &Draws{
		SkipTeamName: skipTeamName,
	}
}

// Calculate returns the number of draws for a given team.
func (w *Draws) Calculate(teamName string, matches *[]pkg.Match) (int, error) {
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
