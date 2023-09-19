package calculators

import (
	"fmt"
	"github.com/jedi-knights/rpi/pkg"
	"github.com/jedi-knights/rpi/pkg/accumulators"
)

// WPCalculator is a calculator that calculates the winning percentage for a given team.
type WPCalculator struct {
	skipTeamName      string
	winsAccumulator   *accumulators.Wins
	lossesAccumulator *accumulators.Losses
	drawsAccumulator  *accumulators.Draws
	matches           []*pkg.Match
}

// NewWPCalculator returns a new WPCalculator.
func NewWPCalculator(skipTeamName string) *WPCalculator {
	return &WPCalculator{
		skipTeamName:      skipTeamName,
		winsAccumulator:   accumulators.NewWins(skipTeamName),
		lossesAccumulator: accumulators.NewLosses(skipTeamName),
		drawsAccumulator:  accumulators.NewDraws(skipTeamName),
	}
}

// Calculate returns the winning percentage for a given team.
func (w *WPCalculator) Calculate(teamName string, matches *[]pkg.Match) (float64, error) {
	if teamName == "" {
		return 0.0, fmt.Errorf("teamName cannot be empty")
	}

	wins, err := w.winsAccumulator.Calculate(teamName, matches)
	if err != nil {
		return 0.0, err
	}

	losses, err := w.lossesAccumulator.Calculate(teamName, matches)
	if err != nil {
		return 0.0, err
	}

	draws, err := w.drawsAccumulator.Calculate(teamName, matches)
	if err != nil {
		return 0.0, err
	}

	matchesPlayed := wins + losses + draws

	if matchesPlayed == 0 {
		return 0.0, nil
	}

	wp := (float64(wins) + 0.5*float64(draws)) / float64(matchesPlayed)

	return wp, nil
}
