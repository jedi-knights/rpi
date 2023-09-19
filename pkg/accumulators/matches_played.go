package accumulators

import (
	"github.com/jedi-knights/rpi/pkg/match"
)

// MatchesPlayed is an accumulator that calculates the number of Matches played by a team.
type MatchesPlayed struct {
	SkipTeamName string
}

// NewMatchesPlayed returns a new instance of MatchesPlayed.
func NewMatchesPlayed(skipTeamName string) *MatchesPlayed {
	return &MatchesPlayed{
		SkipTeamName: skipTeamName,
	}
}

// Calculate returns the number of Matches played by the given team.
func (w *MatchesPlayed) Calculate(teamName string, matches *[]match.Match) (int, error) {
	total := 0

	for _, match := range *matches {
		if len(w.SkipTeamName) > 0 && match.Contains(w.SkipTeamName) {
			continue
		}

		if match.Contains(teamName) {
			total++
		}
	}

	return total, nil
}
