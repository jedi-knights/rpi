package calculators

import (
	"fmt"
	"github.com/jedi-knights/rpi/pkg/accumulators"
	"github.com/jedi-knights/rpi/pkg/errors"
	"github.com/jedi-knights/rpi/pkg/match"
)

// OWPCalculator is a calculator that calculates the OWP for a given team.
type OWPCalculator struct{}

// NewOWPCalculator returns a new OWP calculator.
func NewOWPCalculator() *OWPCalculator {
	return &OWPCalculator{}
}

// Calculate returns the specified team's opponents' average winning percentage.
// To determine Team A's opponents' average winning percentage, the NCAA first computes, for each of Team A's
// opponents, the opponent's wins and ties as compared to the opponent's total games played, in the same way it
// does the calculation for Team A's Element 1.  The only difference is that the NCAA excludes the opponent's
// games against Team A itself.  Thus this first part of the computation determines each opponent's Element 1
// based on games played against teams other than Team A.
func (w *OWPCalculator) Calculate(targetTeamName string, matches *[]match.Match) (float64, error) {
	var err error
	var opponents []string

	if targetTeamName == "" {
		return 0.0, fmt.Errorf(errors.ErrTeamNameRequired)
	}

	// retrieve the specified team's opponents as a slice of strings
	if opponents, err = match.GetOpponents(matches, targetTeamName); err != nil {
		return 0.0, err
	}

	// wpCalculator := NewWPCalculator(targetTeamName)
	winsAccumulator := accumulators.NewWins(targetTeamName)
	lossesAccumulator := accumulators.NewLosses(targetTeamName)
	tiesAccumulator := accumulators.NewTies(targetTeamName)

	// Step 1: for each of Team A's opponents, compute the opponents wins
	// Step 2: for each of the target team's opponents, compute ties minus matches played against the target team..
	accumulator := float64(0.0)
	for _, opponentName := range opponents {
		var wp float64
		var OW, OL, OT int

		if OW, err = winsAccumulator.Calculate(opponentName, matches); err != nil {
			return 0.0, err
		}

		if OL, err = lossesAccumulator.Calculate(opponentName, matches); err != nil {
			return 0.0, err
		}

		if OT, err := tiesAccumulator.Calculate(opponentName, matches); err != nil {
			return 0.0, err
		}

		// I think we need a subset of matches the target team played against a given opponent.

		//if wp, err = wpCalculator.Calculate(opponentName, matches); err != nil {
		//	return 0.0, err
		//}

		accumulator += wp
	}

	owp := accumulator / float64(len(opponents))

	return owp, nil
}
